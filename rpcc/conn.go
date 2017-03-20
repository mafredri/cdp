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
		streams: make(map[string]*streamClients),
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

	recvDone := make(chan error, 1)
	go c.recv(c.notify, recvDone)
	go func() {
		select {
		case <-c.ctx.Done():
		case err := <-recvDone:
			// When we receive Inspector.detached the remote will close
			// the connection afterwards and recvDone will return. Maybe
			// we could give the user time to react to the event before
			// closing?
			// TODO: Do we want to close here, like this?
			c.close(err)
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
	closed   bool

	// Codec encodes and decodes JSON onto conn. There is only one
	// active decoder (recv) and encoder (guaranteed via reqMu).
	codec codec

	mu      sync.Mutex // Protects following.
	reqSeq  uint64
	pending map[uint64]*rpcCall

	reqMu sync.Mutex // Protects following.
	req   rpcRequest

	streamMu sync.Mutex // Protects following.
	streams  map[string]*streamClients
}

type rpcResponse struct {
	// Response to request.
	ID     uint64          `json:"id"`     // Echoes that of the rpcRequest.
	Result json.RawMessage `json:"result"` // Result from invokation, if any.
	Error  *rpcError       `json:"error"`  // Error, if any.

	// RPC notification from remote.
	Method string          `json:"method"` // Method invokation requested by remote.
	Params json.RawMessage `json:"params"` // Method parameters, if any.
}

func (r *rpcResponse) reset() {
	r.ID = 0
	r.Result = nil
	r.Error = nil
	r.Method = ""
	r.Params = nil
}

func (r *rpcResponse) String() string {
	if r.Method != "" {
		return fmt.Sprintf("Method = %s, Params = %s", r.Method, r.Params)
	}
	if r.Error != nil {
		return fmt.Sprintf("ID = %d, Error = %s", r.ID, r.Error.Error())
	}
	return fmt.Sprintf("ID = %d, Result = %s", r.ID, r.Result)
}

type rpcError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *rpcError) Error() string {
	var data string
	if e.Data != "" {
		data = ", data = " + e.Data
	}
	return fmt.Sprintf("rpc error: %s (code = %d%s)", e.Message, e.Code, data)
}

var (
	_ error = (*rpcError)(nil)
)

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

		// Check if this is an RPC notification from the server.
		if resp.Method != "" {
			// Method represents the event that was triggered over the
			// Chrome Debugging Protocol. We do not expect to receive
			// RPC requests, if this was one, the ID field would be set.
			notify(resp.Method, resp.Params)
			continue
		}

		c.mu.Lock()
		call := c.pending[resp.ID]
		delete(c.pending, resp.ID)
		c.mu.Unlock()

		switch {
		case call == nil:
			// No pending call, this could mean there was an error during
			// send or the server sent an unexpected response.
			if enableDebug {
				log.Println("rpcc: no pending call: " + resp.String())
			}
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

type rpcRequest struct {
	ID     uint64      `json:"id"`               // ID chosen by client.
	Method string      `json:"method"`           // Method invoked on remote.
	Params interface{} `json:"params,omitempty"` // Method parameters, if any.
}

// send returns after the call has successfully been dispatched over
// the RPC connection.
func (c *Conn) send(ctx context.Context, call *rpcCall) (err error) {
	defer func() {
		// Give precedence for user cancellation.
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}()

	c.mu.Lock()
	c.reqSeq++
	reqID := c.reqSeq
	c.pending[reqID] = call
	c.mu.Unlock()

	done := make(chan error, 1)
	go func() {
		c.reqMu.Lock()
		c.req.ID = reqID
		c.req.Method = call.Method
		c.req.Params = call.Args

		err := c.codec.Encode(&c.req)

		c.req.Params = nil
		c.reqMu.Unlock()
		done <- err
	}()

	// Abort on user or connection cancellation.
	select {
	case <-c.ctx.Done():
		err = ErrConnClosing
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-done:
	}

	if err != nil {
		// Remove reference on error, avoid
		// unnecessary work in recv.
		c.mu.Lock()
		delete(c.pending, reqID)
		c.mu.Unlock()
		return err
	}

	return nil
}

// notify handles RPC notifications and sends them
// to the appropriate stream listeners.
func (c *Conn) notify(method string, data []byte) {
	c.streamMu.Lock()
	stream := c.streams[method]
	if stream != nil {
		stream.send(data)
	}
	c.streamMu.Unlock()
}

// listen registers a new stream listener (chan) for the RPC notification
// method. Returns a function for removing the listener. Error if the
// connection is closed.
func (c *Conn) listen(method string, client *streamClient) (func(), error) {
	c.streamMu.Lock()
	defer c.streamMu.Unlock()

	if c.streams == nil {
		return nil, ErrConnClosing
	}

	stream, ok := c.streams[method]
	if !ok {
		stream = newStreamService()
		c.streams[method] = stream
	}
	seq := stream.add(client)

	return func() { stream.remove(seq) }, nil
}

// Close closes the connection.
func (c *Conn) close(err error) error {
	// Stop sending on all streams.
	c.streamMu.Lock()
	c.streams = nil
	c.streamMu.Unlock()

	c.cancel()

	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return ErrConnClosing
	}
	c.closed = true
	if err == nil {
		err = ErrConnClosing
	}
	for id, call := range c.pending {
		delete(c.pending, id)
		call.done(err)
	}
	c.mu.Unlock()

	// Conn can be nil if DialContext did not complete.
	if c.conn != nil {
		err = c.conn.Close()
	}

	return err
}

// Close closes the connection.
func (c *Conn) Close() error {
	return c.close(nil)
}

// Debugging, enabled in tests.
var enableDebug = false
