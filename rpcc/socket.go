package rpcc

import (
	"io"

	"github.com/gorilla/websocket"
)

type wsConn interface {
	NextReader() (messageType int, r io.Reader, err error)
	NextWriter(messageType int) (io.WriteCloser, error)
	Close() error
}

// wsReadWriteCloser wraps a gorilla/websocket connection
// and implements io.Reader and io.Writer.
type wsReadWriteCloser struct {
	wsConn
	r io.Reader
}

var (
	_ io.ReadWriteCloser = (*wsReadWriteCloser)(nil)
)

// Read calls Read on the WebSocket Reader and requests the NextReader
// when io.EOF is encountered. Imlpements io.Reader.
func (cw *wsReadWriteCloser) Read(p []byte) (n int, err error) {
	if cw.r != nil {
		// Check if previous reader still has data in the buffer.
		// Otherwise pass on to next reader.
		n, err = cw.r.Read(p)
		if err != io.EOF {
			return n, err
		}
	}
	_, r, err := cw.wsConn.NextReader()
	if err != nil {
		return 0, err
	}
	cw.r = r // Store reader for next call to Read.

	n, err = r.Read(p)
	return n, err
}

// Write requests the NextWriter for the WebSocket and writes the
// message. Implements io.Writer.
func (cw *wsReadWriteCloser) Write(p []byte) (n int, err error) {
	w, err := cw.wsConn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	if n, err = w.Write(p); err != nil {
		return n, err
	}
	err = w.Close()
	return n, err
}
