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

const (
	defaultWriteBufferSize = 4096
)

// DialOption represents a dial option passed to Dial.
type DialOption func(*dialOptions)

// WithCodec returns a DialOption that sets the codec responsible for
// encoding and decoding requests and responses onto the connection.
// This option overrides the default json codec.
func WithCodec(f func(conn io.ReadWriter) Codec) DialOption {
	return func(o *dialOptions) {
		o.codec = f
	}
}

// WithDialer returns a DialOption that sets the dialer for the underlying
// net.Conn. This option overrides the default WebSocket dialer.
func WithDialer(f func(ctx context.Context, addr string) (net.Conn, error)) DialOption {
	return func(o *dialOptions) {
		o.dialer = f
	}
}

// WithWriteBufferSize returns a DialOption that sets the size of the write
// buffer for the underlying websocket connection. Messages larger than this
// size are fragmented according to the websocket specification.
//
// The maximum buffer size for recent versions of Chrome is 104857586 (~100MB),
// for older versions a maximum of 1048562 (~1MB) can be used. This is because
// Chrome does not support websocket fragmentation.
func WithWriteBufferSize(n int) DialOption {
	return func(o *dialOptions) {
		o.wsDialer.WriteBufferSize = n
	}
}

// WithCompression returns a DialOption that enables compression for the
// underlying websocket connection. Use SetCompressionLevel on Conn to
// change the default compression level for subsequent writes.
func WithCompression() DialOption {
	return func(o *dialOptions) {
		o.wsDialer.EnableCompression = true
	}
}

type dialOptions struct {
	codec    func(io.ReadWriter) Codec
	dialer   func(context.Context, string) (net.Conn, error)
	wsDialer websocket.Dialer
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
			ws := &c.dialOpts.wsDialer

			if ws.WriteBufferSize == 0 {
				// Set the default size for use in writeLimiter.
				ws.WriteBufferSize = defaultWriteBufferSize
			}

			// Set NetDial to dial with context, this action will
			// override the HandshakeTimeout setting.
			ws.NetDial = func(network, addr string) (net.Conn, error) {
				conn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
				// Use writeLimiter to avoid writing fragmented
				// websocket messages. We're not accounting for
				// the header length here because it varies, as
				// a result we might block some valid writes
				// that are a few bytes too large.
				return &writeLimiter{
					limit: ws.WriteBufferSize,
					Conn:  conn,
				}, err
			}

			wsConn, _, err := ws.Dial(addr, nil)
			if err != nil {
				return nil, err
			}

			if ws.EnableCompression {
				c.compressionLevel = wsConn.SetCompressionLevel
			}

			return &wsNetConn{conn: wsConn}, nil
		}
	}

	c.conn, err = netDial(ctx, target)
	if err != nil {
		return nil, err
	}
	newCodec := c.dialOpts.codec
	if newCodec == nil {
		newCodec = func(conn io.ReadWriter) Codec {
			return &jsonCodec{
				enc: json.NewEncoder(conn),
				dec: json.NewDecoder(conn),
			}
		}
	}
	c.codec = newCodec(c.conn)

	recvDone := func(err error) {
		// When we receive Inspector.detached the remote will close
		// the connection afterwards and recvDone will return. Maybe
		// we could give the user time to react to the event before
		// closing?
		// TODO(mafredri): Do we want to close here, like this?
		c.close(err)
	}
	go c.recv(c.notify, recvDone)

	return c, nil
}

// writeLimiter wraps a net.Conn and prevents writes of greater length
// than limit. Works around Chrome's lack of support for large or
// fragmented websocket messages and prevents sudden termination of the
// websocket connection. Gives the user an actionable error message when
// writes exceed limit.
type writeLimiter struct {
	limit int
	net.Conn
}

// BUG(mafredri): Chrome does not support websocket fragmentation
// (continuation messages) or messages that exceed 1MB in size.
// This limit was bumped in more recent versions of Chrome which can
// receive messages up to 100MB in size.
// See https://github.com/mafredri/cdp/issues/4 and
// https://github.com/ChromeDevTools/devtools-protocol/issues/24.
func (c *writeLimiter) Write(b []byte) (n int, err error) {
	if len(b) > c.limit {
		return 0, errors.New("rpcc: message too large (increase write buffer size or enable compression)")
	}
	return c.Conn.Write(b)
}

// Codec is used by recv and dispatcher to
// send and receive RPC communication.
type Codec interface {
	// WriteRequest encodes and writes the request onto the
	// underlying connection. Request is re-used between writes and
	// references to it should not be kept.
	WriteRequest(*Request) error
	// ReadResponse decodes a response from the underlying
	// connection. Response is re-used between reads and references
	// to it should not be kept.
	ReadResponse(*Response) error
}

// jsonCodec implements codec.
type jsonCodec struct {
	enc *json.Encoder
	dec *json.Decoder
}

func (c *jsonCodec) WriteRequest(r *Request) error  { return c.enc.Encode(r) }
func (c *jsonCodec) ReadResponse(r *Response) error { return c.dec.Decode(r) }

// Conn represents an active RPC connection.
type Conn struct {
	ctx    context.Context
	cancel context.CancelFunc

	dialOpts dialOptions
	conn     net.Conn

	compressionLevel func(level int) error

	mu      sync.Mutex // Protects following.
	closed  bool
	reqSeq  uint64
	pending map[uint64]*rpcCall

	reqMu sync.Mutex // Protects following.
	req   Request
	// Encodes and decodes JSON onto conn. Encoding is
	// guarded by mutex and decoding is done by recv.
	codec Codec

	streamMu sync.Mutex // Protects following.
	streams  map[string]*streamClients
}

// Response represents an RPC response or notification sent by the server.
type Response struct {
	// RPC response to a Request.
	ID     uint64          `json:"id"`     // Echoes that of the Request.
	Result json.RawMessage `json:"result"` // Result from invokation, if any.
	Error  *ResponseError  `json:"error"`  // Error, if any.

	// RPC notification from remote.
	Method string          `json:"method"` // Method invokation requested by remote.
	Args   json.RawMessage `json:"params"` // Method parameters, if any.
}

func (r *Response) reset() {
	r.ID = 0
	r.Result = nil
	r.Error = nil
	r.Method = ""
	r.Args = nil
}

func (r *Response) String() string {
	if r.Method != "" {
		return fmt.Sprintf("Method = %s, Params = %s", r.Method, r.Args)
	}
	if r.Error != nil {
		return fmt.Sprintf("ID = %d, Error = %s", r.ID, r.Error.Error())
	}
	return fmt.Sprintf("ID = %d, Result = %s", r.ID, r.Result)
}

// ResponseError represents the RPC response error sent by the server.
type ResponseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *ResponseError) Error() string {
	var data string
	if e.Data != "" {
		data = ", data = " + e.Data
	}
	return fmt.Sprintf("rpc error: %s (code = %d%s)", e.Message, e.Code, data)
}

var (
	_ error = (*ResponseError)(nil)
)

// recv decodes and handles RPC responses. Responses to RPC requests
// are forwarded to the pending call, if any. RPC Notifications are
// forwarded by calling notify, synchronously.
func (c *Conn) recv(notify func(string, []byte), done func(error)) {
	var resp Response
	var err error
	for {
		resp.reset()
		if err = c.codec.ReadResponse(&resp); err != nil {
			done(err)
			return
		}

		// Check if this is an RPC notification from the server.
		if resp.Method != "" {
			// Method represents the event that was triggered over the
			// Chrome Debugging Protocol. We do not expect to receive
			// RPC requests, if this was one, the ID field would be set.
			notify(resp.Method, resp.Args)
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

// Request represents an RPC request to be sent to the server.
type Request struct {
	ID     uint64      `json:"id"`               // ID chosen by client.
	Method string      `json:"method"`           // Method invoked on remote.
	Args   interface{} `json:"params,omitempty"` // Method parameters, if any.
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
	if c.closed {
		c.mu.Unlock()
		return ErrConnClosing
	}
	c.reqSeq++
	reqID := c.reqSeq
	c.pending[reqID] = call
	c.mu.Unlock()

	done := make(chan error, 1)
	go func() {
		c.reqMu.Lock()
		c.req.ID = reqID
		c.req.Method = call.Method
		c.req.Args = call.Args

		err := c.codec.WriteRequest(&c.req)

		c.req.Args = nil
		c.reqMu.Unlock()
		done <- err
	}()

	// Abort on user or connection cancellation.
	select {
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
		stream.write(method, data)
	}
	c.streamMu.Unlock()
}

// listen registers a new stream listener (chan) for the RPC notification
// method. Returns a function for removing the listener. Error if the
// connection is closed.
func (c *Conn) listen(method string, w streamWriter) (func(), error) {
	c.streamMu.Lock()
	defer c.streamMu.Unlock()

	if c.streams == nil {
		return nil, ErrConnClosing
	}

	stream, ok := c.streams[method]
	if !ok {
		stream = newStreamClients()
		c.streams[method] = stream
	}
	seq := stream.add(w)

	return func() { stream.remove(seq) }, nil
}

// Close closes the connection.
func (c *Conn) close(err error) error {
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

	// Stop sending on all streams.
	c.streamMu.Lock()
	c.streams = nil
	c.streamMu.Unlock()

	// Conn can be nil if DialContext did not complete.
	if c.conn != nil {
		err = c.conn.Close()
	}

	return err
}

// SetCompressionLevel sets the flate compressions level for writes. Valid level
// range is [-2, 9]. Returns error if compression is not enabled for Conn. See
// package compress/flate for a description of compression levels.
func (c *Conn) SetCompressionLevel(level int) error {
	if c.compressionLevel == nil {
		return errors.New("rpcc: compression is not enabled for Conn")
	}
	return c.compressionLevel(level)
}

// Close closes the connection.
func (c *Conn) Close() error {
	return c.close(nil)
}

// Debugging, enabled in tests.
var enableDebug = false
