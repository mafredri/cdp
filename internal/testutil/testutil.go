package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

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
	conn, err := rpcc.DialContext(ctx, v.WebSocketDebuggerURL)
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
