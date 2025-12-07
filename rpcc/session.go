package rpcc

import (
	"context"
	"errors"
	"log"
)

// SessionOption represents an option for session creation.
type SessionOption func(*sessionOptions)

type sessionOptions struct {
	close func(context.Context) error
}

// WithSessionClose returns a SessionOption that sets a function to be
// invoked when the session connection is closed. The function receives
// a context with a timeout for cleanup operations.
func WithSessionClose(fn func(ctx context.Context) error) SessionOption {
	return func(o *sessionOptions) {
		o.close = fn
	}
}

// NewSession creates a session-scoped connection that shares the
// underlying websocket I/O with the parent. All requests sent via the
// returned connection include the sessionID.
//
// The returned connection should be closed when no longer needed.
// Use WithSessionClose to register a cleanup function.
func NewSession(conn *Conn, sessionID string, opts ...SessionOption) (*Conn, error) {
	if conn.parent != nil {
		return nil, errors.New("rpcc: cannot create session from session connection")
	}

	var opt sessionOptions
	for _, o := range opts {
		o(&opt)
	}

	ctx, cancel := context.WithCancel(conn.ctx)

	sess := &Conn{
		ctx:          ctx,
		cancel:       cancel,
		sessionID:    sessionID,
		parent:       conn,
		pending:      make(map[uint64]*rpcCall),
		streams:      make(map[string]*streamClients),
		sessionClose: opt.close,
	}

	conn.sessionsMu.Lock()
	if conn.sessions == nil {
		conn.sessionsMu.Unlock()
		cancel()
		return nil, ErrConnClosing
	}
	conn.sessions[sessionID] = sess
	conn.sessionsMu.Unlock()

	return sess, nil
}

// SessionID returns the session identifier for this connection.
// Returns an empty string for parent connections.
func (c *Conn) SessionID() string {
	return c.sessionID
}

// recvSessionResponse forwards a response
// to the appropriate session connection.
func (c *Conn) recvSessionResponse(resp *Response) {
	c.sessionsMu.RLock()
	sess := c.sessions[resp.SessionID]
	c.sessionsMu.RUnlock()

	if sess != nil {
		sess.recvResponse(resp)
	} else if enableDebug {
		log.Println("rpcc: no session for response: " + resp.SessionID)
	}
}

// notifySession forwards a notification to
// the appropriate session connection.
func (c *Conn) notifySession(method string, data []byte, sessionID string) {
	c.sessionsMu.RLock()
	sess := c.sessions[sessionID]
	c.sessionsMu.RUnlock()

	if sess != nil {
		sess.notify(method, data)
	}
}
