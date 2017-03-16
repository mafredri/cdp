package rpcc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	// ErrConnClosing indicates that the operation is illegal because
	// the connection is closing.
	ErrConnClosing = errors.New("rpcc: the connection is closing")
)

// DialOption represents a dial option passed to Dial.
type DialOption func(*dialOptions)

// WithDialer returns a DialOption that sets the dialer for the underlying
// net.Conn. This option overrides the default WebSocket dialer.
func WithDialer(f func(ctx context.Context, addr string) (net.Conn, error)) DialOption {
	return func(o *dialOptions) {
		o.dialer = f
	}
}

type dialOptions struct {
	dialer      func(context.Context, string) (net.Conn, error)
	interceptor func(conn io.ReadWriteCloser) io.ReadWriteCloser
}

// Dial connects to target and returns an active connection.
// The target should be a WebSocket URL, format:
// "ws://localhost:9222/target".
func Dial(target string, opts ...DialOption) (*Conn, error) {
	return DialContext(context.Background(), target, opts...)
}

// DialContext is like Dial, with a caller provided Context.
// A nil Context will panic.
func DialContext(ctx context.Context, target string, opts ...DialOption) (conn *Conn, err error) {
	if ctx == nil {
		panic("nil Context")
	}

	c := &Conn{
		pending: make(map[uint64]*rpcCall),
		streams: make(map[string]*streamService),
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())

	defer func() {
		select {
		case <-ctx.Done():
			conn, err = nil, ctx.Err()
		default:
		}
		if err != nil {
			c.Close()
		}
	}()

	for _, o := range opts {
		o(&c.dialOpts)
	}

	netDial := c.dialOpts.dialer
	if netDial == nil {
		netDial = func(ctx context.Context, addr string) (net.Conn, error) {
			d := websocket.Dialer{
				// Set NetDial to dial with context, this action will
				// override the HandshakeTimeout setting.
				NetDial: func(network, addr string) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, addr)
				},
			}
			wsConn, _, err := d.Dial(addr, nil)
			if err != nil {
				return nil, err
			}
			return &wsNetConn{conn: wsConn}, nil
		}
	}

	c.conn, err = netDial(ctx, target)
	if err != nil {
		return nil, err
	}
	c.codec = &jsonCodec{
		Encoder: json.NewEncoder(c.conn),
		Decoder: json.NewDecoder(c.conn),
	}

	recvDone := make(chan error)
	go c.recv(c.notify, recvDone)
	go func() {
		defer close(recvDone)
		select {
		case <-ctx.Done():
		case <-recvDone:
			// When we receive Inspector.detached the remote will close
			// the connection afterwards and recvDone will return. Maybe
			// we could give the user time to react to the event before
			// closing?
			// TODO: Do we want to close here, like this?
			c.Close()
		}
	}()

	return c, nil
}

// codec is used by recv and dispatcher to
// send and receive RPC communication.
type codec interface {
	Encode(v interface{}) error
	Decode(v interface{}) error
}

// jsonCodec implements codec.
type jsonCodec struct {
	*json.Encoder
	*json.Decoder
}

// Conn represents an active RPC connection.
type Conn struct {
	ctx    context.Context
	cancel context.CancelFunc

	dialOpts dialOptions
	conn     net.Conn
	codec    codec

	mu     sync.Mutex
	closed bool

	reqMu sync.Mutex // Protects following.
	req   rpcRequest

	pendingMu sync.Mutex // Protects following.
	reqSeq    uint64
	pending   map[uint64]*rpcCall

	streamMu sync.Mutex // Protects following.
	streams  map[string]*streamService
}

// recv decodes and handles RPC responses. Respones to RPC requests
// are forwarded to the pending call, if any. RPC Notifications are
// forwarded by calling notify, synchronously.
func (c *Conn) recv(notify func(string, []byte), done chan<- error) {
	var resp rpcResponse
	var err error
	for {
		resp.reset()
		if err = c.codec.Decode(&resp); err != nil {
			done <- err
			return
		}

		if resp.Method != "" {
			// Response contained a method, this means we received a request
			// from the server. This is a simplistic approach since we only
			// need to support notifications. If this was a RPC request from
			// the server, the ID field would be included and the server would
			// expect a response.
			notify(resp.Method, resp.Params)
			continue
		}

		c.pendingMu.Lock()
		call := c.pending[resp.ID]
		delete(c.pending, resp.ID)
		c.pendingMu.Unlock()

		switch {
		case call == nil:
			log.Println("rpcc: no pending call: " + resp.String())
		case resp.Error != nil:
			call.done(resp.Error)
		default:
			var err error
			if call.Reply != nil {
				if err = json.Unmarshal(resp.Result, call.Reply); err != nil {
					err = fmt.Errorf("rpcc: decoding %s: %s", call.Method, err.Error())
				}
			}
			call.done(err)
		}
	}
}

func (c *Conn) send(call *rpcCall) {
	c.pendingMu.Lock()
	c.reqSeq++
	reqID := c.reqSeq
	c.pending[reqID] = call
	c.pendingMu.Unlock()

	c.reqMu.Lock()
	c.req.ID = reqID
	c.req.Method = call.Method
	c.req.Params = call.Args
	if err := c.codec.Encode(&c.req); err != nil {
		c.req.Params = nil
		c.reqMu.Unlock()

		// Clean up after error.
		c.pendingMu.Lock()
		call := c.pending[reqID]
		delete(c.pending, reqID)
		c.pendingMu.Unlock()

		if call != nil {
			// Call was not handled by recv.
			call.done(err)
		}
	} else {
		c.req.Params = nil
		c.reqMu.Unlock()
	}
}

// notify handles RPC notifications and sends them
// to the appropriate stream listeners.
func (c *Conn) notify(method string, data []byte) {
	c.streamMu.Lock()
	s := c.streams[method]
	c.streamMu.Unlock()
	if s != nil {
		s.send(data)
	}
}

// listen registers a new stream listener (chan) for the RPC notificaiton
// method. Returns a function for removing the listener. Error if the
// connection is closed.
func (c *Conn) listen(method string, ch chan<- []byte) (func(), error) {
	c.streamMu.Lock()
	defer c.streamMu.Unlock()

	if c.streams == nil {
		return nil, ErrConnClosing
	}

	service, ok := c.streams[method]
	if !ok {
		service = newStreamService()
		c.streams[method] = service
	}
	seq := service.add(ch)

	return func() { service.remove(seq) }, nil
}

// Close closes the connection.
func (c *Conn) Close() (err error) {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	if closed {
		return ErrConnClosing
	}
	c.cancel()

	c.streamMu.Lock()
	streams := c.streams
	c.streams = nil
	c.streamMu.Unlock()
	_ = streams

	if c.conn != nil {
		err = c.conn.Close()
	}
	return err
}
