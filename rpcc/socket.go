package rpcc

import (
	"io"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type websocketConn interface {
	NextReader() (messageType int, r io.Reader, err error)
	NextWriter(messageType int) (io.WriteCloser, error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

// wsNetConn wraps a gorilla/websocket connection
// and implements the net.Conn interface.
type wsNetConn struct {
	conn websocketConn
	r    io.Reader
}

var (
	_ net.Conn = (*wsNetConn)(nil)
)

// Implement net.Conn.
func (cw *wsNetConn) LocalAddr() net.Addr                { return cw.conn.LocalAddr() }
func (cw *wsNetConn) RemoteAddr() net.Addr               { return cw.conn.RemoteAddr() }
func (cw *wsNetConn) SetReadDeadline(t time.Time) error  { return cw.conn.SetReadDeadline(t) }
func (cw *wsNetConn) SetWriteDeadline(t time.Time) error { return cw.conn.SetWriteDeadline(t) }
func (cw *wsNetConn) SetDeadline(t time.Time) error {
	err := cw.SetReadDeadline(t)
	if err != nil {
		return err
	}
	return cw.SetWriteDeadline(t)
}

// Read calls Read on the WebSocket Reader and requests the NextReader
// when io.EOF is encountered. Imlpements io.Reader as part of net.Conn.
func (cw *wsNetConn) Read(p []byte) (n int, err error) {
	if cw.r != nil {
		// Check if previous reader still has data in the buffer.
		// Otherwise pass on to next reader.
		n, err = cw.r.Read(p)
		if err != io.EOF {
			return n, err
		}
	}
	_, r, err := cw.conn.NextReader()
	if err != nil {
		return 0, err
	}
	cw.r = r // Store reader for next call to Read.

	n, err = r.Read(p)
	return n, err
}

// Write requests the NextWriter for the WebSocket and writes the
// message. Implements io.Writer as part of net.Conn.
func (cw *wsNetConn) Write(p []byte) (n int, err error) {
	w, err := cw.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	if n, err = w.Write(p); err != nil {
		return n, err
	}
	err = w.Close()
	return n, err
}

// Close calls Close on the underlying connection.
func (cw *wsNetConn) Close() error {
	return cw.conn.Close()
}
