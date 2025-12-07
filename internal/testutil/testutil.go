package testutil

import (
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// logCodec wraps an rpcc.Codec and logs all requests and responses.
type logCodec struct {
	t   *testing.T
	enc *json.Encoder
	dec *json.Decoder
}

func (c *logCodec) WriteRequest(r *rpcc.Request) error {
	c.t.Helper()
	data, _ := json.Marshal(r)
	c.t.Logf("SEND: %s", data)
	return c.enc.Encode(r)
}

func (c *logCodec) ReadResponse(r *rpcc.Response) error {
	c.t.Helper()
	if err := c.dec.Decode(r); err != nil {
		return err
	}
	data, _ := json.Marshal(r)
	c.t.Logf("RECV: %s", data)
	return nil
}

// NewLogCodec returns a new rpcc codec that logs all requests and responses.
func NewLogCodec(t *testing.T) rpcc.DialOption {
	return rpcc.WithCodec(func(conn io.ReadWriter) rpcc.Codec {
		return &logCodec{
			t:   t,
			enc: json.NewEncoder(conn),
			dec: json.NewDecoder(conn),
		}
	})
}

// Client represents a test client.
type Client struct {
	t      *testing.T
	Conn   *rpcc.Conn
	Client *cdp.Client
}

// NewPage creates a new page target.
func (c *Client) NewPage(ctx context.Context) *Target {
	return NewTarget(ctx, c.t, c.Client)
}

// NewClient returns a new test client.
func NewClient(ctx context.Context, t *testing.T) *Client {
	t.Helper()

	devt := devtool.New("http://localhost:9222")
	v, err := devt.Version(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var opts []rpcc.DialOption
	if testing.Verbose() {
		opts = append(opts, NewLogCodec(t))
	}

	conn, err := rpcc.DialContext(ctx, v.WebSocketDebuggerURL, opts...)
	if err != nil {
		t.Fatal(err)
	}

	return &Client{
		t:      t,
		Conn:   conn,
		Client: cdp.NewClient(conn),
	}
}

// Target represents a (page) target.
type Target struct {
	t  *testing.T
	c  *cdp.Client
	id target.ID
}

// ID returns the target ID.
func (t *Target) ID() target.ID {
	return t.id
}

// Close closes the target.
func (t *Target) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reply, err := t.c.Target.CloseTarget(ctx,
		target.NewCloseTargetArgs(t.ID()))
	if err != nil {
		t.t.Error(err)
	}
	if !reply.Success {
		t.t.Error("close target failed")
	}
}

// NewTarget creates a new target.
func NewTarget(ctx context.Context, t *testing.T, c *cdp.Client) *Target {
	reply, err := c.Target.CreateTarget(ctx,
		target.NewCreateTargetArgs("about:blank"))
	if err != nil {
		t.Fatal(err)
	}

	return &Target{t: t, c: c, id: reply.TargetID}
}
