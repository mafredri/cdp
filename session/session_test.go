package session_test

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp/session"
)

type testClient struct {
	t    *testing.T
	conn *rpcc.Conn
	cl   *cdp.Client
}

func (c *testClient) NewPage(ctx context.Context) *testTarget {
	return newTestTarget(c.t, ctx, c.cl)
}

func getClient(t *testing.T, ctx context.Context) *testClient {
	// TODO(mafredri): Uncomment helper when no longer testing Go 1.8.
	// t.Helper()
	if !*browserFlag {
		t.Skip("This test only runs in the browser, skipping")
	}

	devt := devtool.New("http://localhost:9222")
	v, err := devt.Version(ctx)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := rpcc.DialContext(ctx, v.WebSocketDebuggerURL)
	if err != nil {
		t.Fatal(err)
	}

	return &testClient{
		t:    t,
		conn: conn,
		cl:   cdp.NewClient(conn),
	}
}

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// Give time for goroutines to settle at the end (increases coverage).
	defer func() {
		time.Sleep(10 * time.Millisecond)
	}()

	c := getClient(t, ctx)

	sc, err := session.NewClient(c.cl)
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Close()

	newPage := c.NewPage(ctx)
	// Close later.
	// defer newPage.Close()

	// Test session usage.
	pageConn, err := sc.Dial(ctx, newPage.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer pageConn.Close()

	pageC := cdp.NewClient(pageConn)

	err = pageC.Page.Enable(ctx)
	if err != nil {
		t.Error(err)
	}
	fired, err := pageC.Page.DOMContentEventFired(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fired.Close()

	// TODO(maf): Use testdata / sample HTML for test.
	_, err = pageC.Page.Navigate(ctx,
		page.NewNavigateArgs("https://www.google.com"))
	if err != nil {
		t.Error(err)
	}

	_, err = fired.Recv()
	if err != nil {
		t.Error(err)
	}

	eval, err := pageC.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs(`document.title`))
	if err != nil {
		t.Error(err)
	}

	var title string
	err = json.Unmarshal(eval.Result.Value, &title)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(title, "Google") {
		t.Error("bad title:", title)
	}

	// Close the page, this should also close pageConn.
	newPage.Close()
	select {
	case <-pageConn.Context().Done():
	case <-ctx.Done():
		t.Error("timed out waiting for session to close")
	}
}

func TestClient_CloseClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := getClient(t, ctx)

	// Test connection closure, should close session client.
	sc, err := session.NewClient(c.cl)
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Close()

	newPage := newTestTarget(t, ctx, c.cl)
	defer newPage.Close()

	_, err = sc.Dial(ctx, newPage.ID) // Closed by sc.Close().
	if err != nil {
		t.Error(err)
	}

	sc.Close()
	_, err = sc.Dial(ctx, newPage.ID)
	if err == nil {
		t.Error("Dial: expected error after Close, got nil")
	}
}
func TestClient_CloseUnderlyingConn(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := getClient(t, ctx)

	// Get a target ID for use in Dial (doesn't matter that it's closed).
	p := c.NewPage(ctx)
	p.Close()
	targetID := p.ID

	// Test connection closure, should close session client.
	sc, err := session.NewClient(c.cl)
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Close()

	c.conn.Close()
	time.Sleep(10 * time.Millisecond) // Give time for context propagation.

	_, err = sc.Dial(ctx, targetID)
	if err == nil {
		t.Error("Dial succeeded on a closed connection")
	}
}

func TestClient_NewClientOnClosedConn(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := getClient(t, ctx)
	c.conn.Close()

	_, err := session.NewClient(c.cl)
	if err == nil {
		t.Error("NewClient: rpcc.Conn is closed, expected error, got nil ")
	}
}

type testTarget struct {
	t   *testing.T
	ctx context.Context
	c   *cdp.Client
	ID  target.ID
}

func (t *testTarget) Close() {
	reply, err := t.c.Target.CloseTarget(t.ctx,
		target.NewCloseTargetArgs(t.ID))
	if err != nil {
		t.t.Error(err)
	}
	if !reply.Success {
		t.t.Error("close target failed")
	}
}

func newTestTarget(t *testing.T, ctx context.Context, c *cdp.Client) *testTarget {
	reply, err := c.Target.CreateTarget(ctx,
		target.NewCreateTargetArgs("about:blank"))
	if err != nil {
		t.Fatal(err)
	}

	return &testTarget{t: t, ctx: ctx, c: c, ID: reply.TargetID}
}

var (
	browserFlag = flag.Bool("browser", false, "Test with browser")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
