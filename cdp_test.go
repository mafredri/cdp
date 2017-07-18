package cdp_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"
)

var TestSockSrv string

type testSocketServer struct {
	ws websocket.Upgrader
}

func (tc *testSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := tc.ws.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	switch r.RequestURI {
	case "/example_logging":
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, []byte(`{"id":1,"result":{}}`))
			if err != nil {
				return
			}
		}
	}
}

func TestBrowser_RemoteDebuggingProtocol(t *testing.T) {
	if !*testBrowser {
		t.SkipNow()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	devt := devtool.New(fmt.Sprintf("http://localhost:%d", *remoteDebuggingPort))
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := cdp.NewClient(conn)

	if err = c.Page.Enable(ctx); err != nil {
		panic(err)
	}
	if err = c.Runtime.Enable(ctx); err != nil {
		panic(err)
	}

	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		panic(err)
	}
	defer domContentEventFired.Close()

	// TODO(mafredri): Create a testdata HTML instead of relying on google.com.
	_, err = c.Page.Navigate(ctx, page.NewNavigateArgs("https://www.google.com"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = domContentEventFired.Recv()
	if err != nil {
		t.Fatal(err)
	}

	eval, err := c.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs("document.title"))
	if err != nil {
		t.Fatal(err)
	}

	want := "Google"
	var got string
	if err = json.Unmarshal(eval.Result.Value, &got); err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Evaluate(document.title): got %q, want %q", got, want)
	}
}

var (
	testBrowser         = flag.Bool("browser", false, "Run browser tests")
	remoteDebuggingPort = flag.Int("rdp", 9222, "Remote debugging port (for browser tests)")
)

func TestMain(m *testing.M) {
	flag.Parse()

	srv := httptest.NewServer(&testSocketServer{})
	TestSockSrv = strings.TrimPrefix(srv.URL, "http://")
	code := m.Run()
	srv.Close()
	os.Exit(code)
}
