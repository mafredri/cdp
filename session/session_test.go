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

// TODO(maf): Improve test, it is just a high-level use case atm.
func TestClient(t *testing.T) {
	if !*browserFlag {
		t.Skip("This test only runs in the browser, skipping")
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// Give time for goroutines to settle at the end (increases coverage).
	defer func() {
		time.Sleep(10 * time.Millisecond)
	}()

	devt := devtool.New("http://localhost:9222")
	p, err := devt.Create(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer devt.Close(ctx, p)

	conn, err := rpcc.DialContext(ctx, p.WebSocketDebuggerURL)
	if err != nil {
		t.Fatal(err)
	}

	c := cdp.NewClient(conn)

	sc, err := session.NewClient(c)
	if err != nil {
		t.Error(err)
	}
	defer sc.Close()

	createReply, err := c.Target.CreateTarget(ctx,
		target.NewCreateTargetArgs("about:blank"))
	if err != nil {
		t.Error(err)
	}

	// Increase test coverage, slightly...
	unusedConn, err := sc.Dial(ctx, createReply.TargetID)
	if err != nil {
		t.Error(err)
	}
	_ = unusedConn // Will be closed via sc.Close().

	pageConn, err := sc.Dial(ctx, createReply.TargetID)
	if err != nil {
		t.Error(err)
	}
	defer pageConn.Close()

	pageC := cdp.NewClient(pageConn)

	err = pageC.Page.Enable(ctx)
	if err != nil {
		t.Error(err)
	}
	fired, err := pageC.Page.DOMContentEventFired(ctx)
	if err != nil {
		t.Error(err)
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

	closeReply, err := c.Target.CloseTarget(ctx,
		target.NewCloseTargetArgs(createReply.TargetID))
	_ = closeReply
	if err != nil {
		t.Error(err)
	}
	if !closeReply.Success {
		t.Error("close target failed")
	}

	// CloseTarget should cause pageConn to close.
	select {
	case <-pageConn.Context().Done():
	case <-ctx.Done():
		t.Error("timed out waiting for session to close")
	}
}

var (
	browserFlag = flag.Bool("browser", false, "Test with browser")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
