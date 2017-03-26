package cdp_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
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

func TestMain(m *testing.M) {
	srv := httptest.NewServer(&testSocketServer{})
	TestSockSrv = strings.TrimPrefix(srv.URL, "http://")
	code := m.Run()
	srv.Close()
	os.Exit(code)
}
