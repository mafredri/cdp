package rpcc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

type testServer struct {
	srv    *httptest.Server
	wsConn *websocket.Conn
	conn   *Conn
}

func (ts *testServer) Close() error {
	defer ts.srv.Close()
	return ts.conn.Close()
}

func newTestServer(t *testing.T, respond func(*websocket.Conn, *rpcRequest) error) *testServer {
	upgrader := &websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, r.Header)
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		for {
			var req rpcRequest
			if err := conn.ReadJSON(&req); err != nil {
				break
			}
			if err := respond(conn, &req); err != nil {
				break
			}
		}
	}))

	conn, err := Dial("ws" + strings.TrimPrefix(srv.URL, "http"))
	if err != nil {
		t.Fatal(err)
	}

	return &testServer{
		srv:  srv,
		conn: conn,
	}
}

func TestConn_Invoke(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *rpcRequest) error {
		resp := rpcResponse{
			ID:     req.ID,
			Result: []byte(`"hello"`),
		}
		return conn.WriteJSON(&resp)
	})
	defer srv.Close()

	var reply string
	err := Invoke(nil, "test.Hello", nil, &reply, srv.conn)
	if err != nil {
		t.Error(err)
	}

	if reply != "hello" {
		t.Errorf("test.Hello: got reply %q, want %q", reply, "hello")
	}
}

func TestConn_InvokeError(t *testing.T) {
	want := "bad request"

	srv := newTestServer(t, func(conn *websocket.Conn, req *rpcRequest) error {
		resp := rpcResponse{
			ID:    req.ID,
			Error: &rpcError{Message: want},
		}
		return conn.WriteJSON(&resp)
	})
	defer srv.Close()

	switch err := Invoke(nil, "test.Hello", nil, nil, srv.conn).(type) {
	case *rpcError:
		if err.Message != want {
			t.Errorf("Invoke err.Message: got %q, want %q", err.Message, want)
		}
	default:
		t.Errorf("Invoke: want *rpcError, got %#v", err)
	}
}

func TestConn_Notify(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *rpcRequest) error {
		resp := rpcResponse{ID: req.ID}
		if err := conn.WriteJSON(&resp); err != nil {
			t.Fatal(err)
		}

		resp.Method = "test.Notify"
		resp.Params = []byte(`"hello"`)
		if err := conn.WriteJSON(&resp); err != nil {
			t.Fatal(err)
		}
		return nil
	})
	defer srv.Close()

	s, err := NewStream(nil, "test.Notify", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	// Fire off server response.
	go Invoke(nil, "test.Hello", nil, nil, srv.conn)

	var reply string
	if err = s.RecvMsg(&reply); err != nil {
		t.Error(err)
	}

	want := "hello"
	if reply != want {
		t.Errorf("test.Notify reply: got %q, want %q", reply, want)
	}
}

func TestConn_RemoteDisconnected(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *rpcRequest) error {
		conn.Close()
		return nil
	})
	defer srv.Close()

	if err := Invoke(nil, "test.Hello", nil, nil, srv.conn); err != ErrConnClosing {
		t.Errorf("Invoke error: got %v, want ErrConnClosing", err)
	}
}
