package rpcc

import (
	"io"
	"net"
	"testing"
	"time"
)

type fakeSocketConn struct {
	nextReaderCount int
	nextWriterCount int
	closed          bool
	reader          *fakeReader
	writer          *fakeWriteCloser
}

// Implement websocketConn.
func (c *fakeSocketConn) LocalAddr() net.Addr                { return nil }
func (c *fakeSocketConn) RemoteAddr() net.Addr               { return nil }
func (c *fakeSocketConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *fakeSocketConn) SetWriteDeadline(t time.Time) error { return nil }

func (c *fakeSocketConn) NextReader() (int, io.Reader, error) {
	c.nextReaderCount++
	c.reader = new(fakeReader)
	return 0, c.reader, nil
}
func (c *fakeSocketConn) NextWriter(int) (io.WriteCloser, error) {
	c.nextWriterCount++
	c.writer = new(fakeWriteCloser)
	return c.writer, nil
}
func (c *fakeSocketConn) Close() error {
	c.closed = true
	return nil
}

type fakeReader struct {
	count int
	err   error
}

func (c *fakeReader) Read(p []byte) (n int, err error) {
	c.count++
	return 0, c.err
}

type fakeWriteCloser struct {
	count  int
	err    error
	closed bool
}

func (c *fakeWriteCloser) Write(p []byte) (n int, err error) {
	c.count++
	return 0, nil
}
func (c *fakeWriteCloser) Close() error {
	c.closed = true
	return nil
}

func TestSocket_Read(t *testing.T) {
	fakeConn := &fakeSocketConn{}
	conn := &wsNetConn{conn: fakeConn}
	conn.Read(nil)
	if fakeConn.nextReaderCount != 1 {
		t.Errorf("expected NextReader to be called once, got %d", fakeConn.nextReaderCount)
	}
	if fakeConn.reader.count != 1 {
		t.Errorf("expected Read to be called once, got %d", fakeConn.nextReaderCount)
	}

	prevReader := fakeConn.reader
	// Should fetch next reader on EOF.
	fakeConn.reader.err = io.EOF
	conn.Read(nil)
	if prevReader.count != 2 {
		t.Errorf("expected Read to be called 2 times on previous reader, got %d", prevReader.count)
	}
	if fakeConn.nextReaderCount != 2 {
		t.Errorf("expected NextReader to be called 2 times, got %d", fakeConn.nextReaderCount)
	}
	if fakeConn.reader.count != 1 {
		t.Errorf("expected Read to be called once on new reader, got %d", fakeConn.reader.count)
	}
}

func TestSocket_Write(t *testing.T) {
	fakeConn := &fakeSocketConn{}
	conn := &wsNetConn{conn: fakeConn}

	conn.Write(nil)
	if fakeConn.nextWriterCount != 1 {
		t.Errorf("expected NextWriter to be called once, got %d", fakeConn.nextWriterCount)
	}
	if fakeConn.writer.count != 1 {
		t.Errorf("expected Write to be called once, got %d", fakeConn.writer.count)
	}
	if !fakeConn.writer.closed {
		t.Errorf("expected Close to be called")
	}
}

func TestSocket_Close(t *testing.T) {
	fakeConn := &fakeSocketConn{}
	conn := &wsNetConn{conn: fakeConn}

	conn.Close()
	if !fakeConn.closed {
		t.Errorf("expected Close to be called, connection is not closed")
	}
}
