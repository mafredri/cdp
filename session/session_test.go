package session_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/internal/testutil"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/session"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Test Page One</title></head><body><h1>Page 1</h1></body></html>`)
	})
	mux.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Test Page Two</title></head><body><h1>Page 2</h1></body></html>`)
	})
	return httptest.NewServer(mux)
}

func checkBrowser(t *testing.T) {
	t.Helper()
	if !*browserFlag {
		t.Skip("Test requires browser, skipping...")
	}
}

func TestManager(t *testing.T) {
	type sessionMode struct {
		name string
		opts []session.ManagerOption
	}
	modes := []sessionMode{
		{name: "Flatten", opts: nil},
		{name: "Legacy", opts: []session.ManagerOption{session.WithNoFlatten()}},
	}
	for _, mode := range modes {
		t.Run(mode.name, func(t *testing.T) {
			t.Run("Dial", func(t *testing.T) {
				testManagerDial(t, mode.opts)
			})
			t.Run("Close", func(t *testing.T) {
				testManagerClose(t, mode.opts)
			})
			t.Run("CloseUnderlyingConn", func(t *testing.T) {
				testManagerCloseUnderlyingConn(t, mode.opts)
			})
			t.Run("MultipleSessions", func(t *testing.T) {
				testManagerMultipleSessions(t, mode.opts)
			})
		})
	}

	t.Run("ClosesErrorChan", testManagerClosesErrorChan)
	t.Run("NewOnClosedConn", testManagerNewOnClosedConn)
}

func testManagerDial(t *testing.T, opts []session.ManagerOption) {
	checkBrowser(t)

	srv := newTestServer()
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// Give time for goroutines to settle at the end (increases coverage).
	defer func() {
		time.Sleep(10 * time.Millisecond)
	}()

	c := testutil.NewClient(ctx, t)

	m, err := session.NewManager(c.Client, opts...)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	newPage := c.NewPage(ctx)

	// Test session usage.
	pageConn, err := m.Dial(ctx, newPage.ID())
	if err != nil {
		t.Fatal(err)
	}
	defer pageConn.Close()

	pageC := cdp.NewClient(pageConn)

	err = pageC.Page.Enable(ctx, nil)
	if err != nil {
		t.Error(err)
	}
	fired, err := pageC.Page.DOMContentEventFired(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fired.Close()

	_, err = pageC.Page.Navigate(ctx,
		page.NewNavigateArgs(srv.URL+"/page1"))
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

	if title != "Test Page One" {
		t.Errorf("bad title: got %q, want %q", title, "Test Page One")
	}

	// Close the page, this should also close pageConn.
	newPage.Close()
	select {
	case <-pageConn.Context().Done():
	case <-ctx.Done():
		t.Error("timed out waiting for session to close")
	}
}

func testManagerClose(t *testing.T, opts []session.ManagerOption) {
	checkBrowser(t)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)

	// Test connection closure, should close session client.
	m, err := session.NewManager(c.Client, opts...)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	newPage := c.NewPage(ctx)
	defer newPage.Close()

	_, err = m.Dial(ctx, newPage.ID())
	if err != nil {
		t.Error(err)
	}

	m.Close()
	_, err = m.Dial(ctx, newPage.ID())
	if err == nil {
		t.Error("Dial: expected error after Close, got nil")
	}
}

func testManagerClosesErrorChan(t *testing.T) {
	checkBrowser(t)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)
	m, err := session.NewManager(c.Client)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		m.Close()
	}()

	for range m.Err() {
		t.Fatal("channel should have been closed")
	}
}

func testManagerCloseUnderlyingConn(t *testing.T, opts []session.ManagerOption) {
	checkBrowser(t)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)

	// Get a target ID for use in Dial (doesn't matter that it's closed).
	p := c.NewPage(ctx)
	p.Close()
	targetID := p.ID()

	// Test connection closure, should close session Manager.
	m, err := session.NewManager(c.Client, opts...)
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

func testManagerMultipleSessions(t *testing.T, opts []session.ManagerOption) {
	checkBrowser(t)

	srv := newTestServer()
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)

	m, err := session.NewManager(c.Client, opts...)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	page1 := c.NewPage(ctx)
	defer page1.Close()
	page2 := c.NewPage(ctx)
	defer page2.Close()

	conn1, err := m.Dial(ctx, page1.ID())
	if err != nil {
		t.Fatal(err)
	}
	defer conn1.Close()

	conn2, err := m.Dial(ctx, page2.ID())
	if err != nil {
		t.Fatal(err)
	}
	defer conn2.Close()

	client1 := cdp.NewClient(conn1)
	client2 := cdp.NewClient(conn2)

	if err := client1.Page.Enable(ctx, nil); err != nil {
		t.Fatal(err)
	}
	if err := client2.Page.Enable(ctx, nil); err != nil {
		t.Fatal(err)
	}

	fired1, err := client1.Page.DOMContentEventFired(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fired1.Close()

	fired2, err := client2.Page.DOMContentEventFired(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fired2.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, err := client1.Page.Navigate(ctx,
			page.NewNavigateArgs(srv.URL+"/page1"))
		if err != nil {
			t.Error("page1 navigate:", err)
		}
	}()

	go func() {
		defer wg.Done()
		_, err := client2.Page.Navigate(ctx,
			page.NewNavigateArgs(srv.URL+"/page2"))
		if err != nil {
			t.Error("page2 navigate:", err)
		}
	}()

	wg.Wait()

	wg.Add(2)

	var title1, title2 string

	go func() {
		defer wg.Done()
		if _, err := fired1.Recv(); err != nil {
			t.Error("page1 recv:", err)
			return
		}
		eval, err := client1.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs(`document.title`))
		if err != nil {
			t.Error("page1 eval:", err)
			return
		}
		if err := json.Unmarshal(eval.Result.Value, &title1); err != nil {
			t.Error("page1 unmarshal:", err)
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := fired2.Recv(); err != nil {
			t.Error("page2 recv:", err)
			return
		}
		eval, err := client2.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs(`document.title`))
		if err != nil {
			t.Error("page2 eval:", err)
			return
		}
		if err := json.Unmarshal(eval.Result.Value, &title2); err != nil {
			t.Error("page2 unmarshal:", err)
		}
	}()

	wg.Wait()

	if title1 != "Test Page One" {
		t.Errorf("page1: got title %q, want %q", title1, "Test Page One")
	}
	if title2 != "Test Page Two" {
		t.Errorf("page2: got title %q, want %q", title2, "Test Page Two")
	}
}

func testManagerNewOnClosedConn(t *testing.T) {
	checkBrowser(t)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := testutil.NewClient(ctx, t)
	c.Conn.Close()

	_, err := session.NewManager(c.Client)
	if err == nil {
		t.Error("NewManager: rpcc.Conn is closed, expected error, got nil ")
	}
}

var browserFlag = flag.Bool("browser", false, "Test with browser")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
