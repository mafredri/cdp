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
	"github.com/mafredri/cdp/internal/testutil"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/session"
)

func TestManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// Give time for goroutines to settle at the end (increases coverage).
	defer func() {
		time.Sleep(10 * time.Millisecond)
	}()

	c := testutil.NewClient(ctx, t)

	m, err := session.NewManager(c.Client)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	newPage := c.NewPage(ctx)
	// Close later.
	// defer newPage.Close()

	// Test session usage.
	pageConn, err := m.Dial(ctx, newPage.ID())
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

func TestManager_Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)

	// Test connection closure, should close session client.
	m, err := session.NewManager(c.Client)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	newPage := c.NewPage(ctx)
	defer newPage.Close()

	_, err = m.Dial(ctx, newPage.ID()) // Closed by sc.Close().
	if err != nil {
		t.Error(err)
	}

	m.Close()
	_, err = m.Dial(ctx, newPage.ID())
	if err == nil {
		t.Error("Dial: expected error after Close, got nil")
	}
}
func TestManager_CloseUnderlyingConn(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)

	// Get a target ID for use in Dial (doesn't matter that it's closed).
	p := c.NewPage(ctx)
	p.Close()
	targetID := p.ID()

	// Test connection closure, should close session Manager.
	m, err := session.NewManager(c.Client)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	c.Conn.Close()
	time.Sleep(10 * time.Millisecond) // Give time for context propagation.

	_, err = m.Dial(ctx, targetID)
	if err == nil {
		t.Error("Dial succeeded on a closed connection")
	}
}

func TestManager_NewOnClosedConn(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)
	c.Conn.Close()

	_, err := session.NewManager(c.Client)
	if err == nil {
		t.Error("NewManager: rpcc.Conn is closed, expected error, got nil ")
	}
}

var (
	browserFlag = flag.Bool("browser", false, "Test with browser")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
