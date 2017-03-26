package cdpdom

import (
	"context"
	"errors"
	"io"
	"net"
	"testing"

	"github.com/mafredri/cdp/rpcc"
)

type testCodec struct {
	conn      chan *rpcc.Request
	writeErr  error
	readErr   error
	respErr   error
	event     string
	eventArgs []byte
}

func (c *testCodec) WriteRequest(r *rpcc.Request) error {
	if c.writeErr != nil {
		return c.writeErr
	}
	c.conn <- r
	return nil
}
func (c *testCodec) ReadResponse(r *rpcc.Response) error {
	req, ok := <-c.conn
	if !ok {
		return errors.New("closed")
	}
	if c.event != "" {
		r.Method = c.event
		if c.eventArgs != nil {
			r.Args = c.eventArgs
		} else {
			r.Args = []byte(`{}`)
		}
		return c.readErr
	}
	r.ID = req.ID
	if c.respErr != nil {
		r.Error = &rpcc.ResponseError{Message: c.respErr.Error()}
	} else {
		r.Result = []byte(`{}`)
	}
	return c.readErr
}

func newTestConn(t *testing.T) (*rpcc.Conn, *testCodec, func()) {
	tc := &testCodec{conn: make(chan *rpcc.Request, 1)}
	dialer := func(_ context.Context, _ string) (net.Conn, error) { return nil, nil }
	codec := func(_ io.ReadWriter) rpcc.Codec {
		return tc
	}
	conn, err := rpcc.Dial("", rpcc.WithDialer(dialer), rpcc.WithCodec(codec))
	if err != nil {
		t.Fatal(err)
	}
	return conn, tc, func() {
		conn.Close()
		close(tc.conn)
	}
}
