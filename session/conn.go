package session

import (
	"io"

	"github.com/mafredri/cdp/internal/errors"
)

// closeConn is a fake connection with a close function.
type closeConn struct{ close func() error }

var _ io.ReadWriteCloser = (*closeConn)(nil)

func (c *closeConn) Close() error                      { return c.close() }
func (c *closeConn) Read(b []byte) (n int, err error)  { return 0, errors.New("not allowed") }
func (c *closeConn) Write(b []byte) (n int, err error) { return 0, errors.New("not allowed") }
