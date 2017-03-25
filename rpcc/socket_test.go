package rpcc

import (
	"errors"
	"io"
	"net"
	"testing"
	"time"
)

type fakeSocketConn struct {
	nextReaderCount int
	nextWriterCount int
	nextWriterErr   error
	addr            *testAddr
	closed          bool
	reader          *fakeReader
	writer          *fakeWriteCloser
	writerErr       error
	deadlineErr     error
}

// Implements net.Addr.
type testAddr struct {
	net string
	str string
}

func (a *testAddr) Network() string { return a.net }
func (a *testAddr) String() string  { return a.str }

// Implement websocketConn.
func (c *fakeSocketConn) LocalAddr() net.Addr                { return c.addr }
func (c *fakeSocketConn) RemoteAddr() net.Addr               { return c.addr }
func (c *fakeSocketConn) SetReadDeadline(t time.Time) error  { return c.deadlineErr }
func (c *fakeSocketConn) SetWriteDeadline(t time.Time) error { return c.deadlineErr }

func (c *fakeSocketConn) NextReader() (int, io.Reader, error) {
	c.nextReaderCount++
	c.reader = new(fakeReader)
	return 0, c.reader, nil
}
func (c *fakeSocketConn) NextWriter(int) (io.WriteCloser, error) {
	c.nextWriterCount++
	c.writer = new(fakeWriteCloser)
	c.writer.err = c.writerErr
	return c.writer, c.nextWriterErr
}
func (c *fakeSocketConn) Close() error {
	c.closed = true
	return nil
}

type fakeReader struct {
	count int
	n     int
	err   error
}

func (c *fakeReader) Read(p []byte) (n int, err error) {
	c.count++
	return c.n, c.err
}

type fakeWriteCloser struct {
	count  int
	n      int
	err    error
	closed bool
}

func (c *fakeWriteCloser) Write(p []byte) (n int, err error) {
	c.count++
	return c.n, c.err
}
func (c *fakeWriteCloser) Close() error {
	c.closed = true
	return c.err
}

func TestSocket_netConn(t *testing.T) {
	fakeConn := &fakeSocketConn{addr: &testAddr{net: "tcp", str: "127.0.0.1:80"}}
	conn := &wsNetConn{conn: fakeConn}

	addr := conn.LocalAddr()
	if addr != fakeConn.addr {
		t.Errorf("LocalAddr() wrong addr: got %v, want %v", addr, fakeConn.addr)
	}
	addr = conn.RemoteAddr()
	if addr != fakeConn.addr {
		t.Errorf("RemoteAddr() wrong addr: got %v, want %v", addr, fakeConn.addr)
	}

	err := conn.SetDeadline(time.Now())
	if err != nil {
		t.Errorf("SetDeadline() returned error, got %v, want nil", err)
	}

	fakeConn.deadlineErr = errors.New("error")
	err = conn.SetDeadline(time.Now())
	if err != fakeConn.deadlineErr {
		t.Errorf("SetDeadline() want error, got %v", err)
	}
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

	fakeConn.reader.n = 10
	fakeConn.reader.err = errors.New("read error")
	n, err := conn.Read(nil)
	if n != 10 {
		t.Errorf("want n = 10; got %d", n)
	}
	if err != fakeConn.reader.err {
		t.Errorf("got error: %v; want error: %v", err, fakeConn.reader.err)
	}
}

func TestSocket_Write(t *testing.T) {
	fakeConn := &fakeSocketConn{}
	conn := &wsNetConn{conn: fakeConn}

	_, err := conn.Write(nil)
	if err != nil {
		t.Errorf("Write() got error: %v; want nil", err)
	}
	if fakeConn.nextWriterCount != 1 {
		t.Errorf("expected NextWriter to be called once, got %d", fakeConn.nextWriterCount)
	}
	if fakeConn.writer.count != 1 {
		t.Errorf("expected Write to be called once, got %d", fakeConn.writer.count)
	}
	if !fakeConn.writer.closed {
		t.Errorf("expected Close to be called")
	}

	wantErr := errors.New("disconnect")
	fakeConn.nextWriterErr = wantErr
	_, err = conn.Write(nil)
	if err != wantErr {
		t.Errorf("Write() got %v, want %v", err, wantErr)
	}

	fakeConn.nextWriterErr = nil
	wantErr = errors.New("could not write")
	fakeConn.writerErr = wantErr
	_, err = conn.Write(nil)
	if err != wantErr {
		t.Errorf("Write() got %v, want %v", err, wantErr)
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
