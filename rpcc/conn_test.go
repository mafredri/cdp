package rpcc

import (
	"compress/flate"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type testServer struct {
	srv    *httptest.Server
	wsConn *websocket.Conn // Set after Dial.
	conn   *Conn
}

func (ts *testServer) Close() error {
	defer ts.srv.Close()
	return ts.conn.Close()
}

func newTestServer(t testing.TB, respond func(*websocket.Conn, *Request) error) *testServer {
	// Timeouts to prevent tests from running forever.
	timeout := 5 * time.Second

	var err error
	ts := &testServer{}
	upgrader := &websocket.Upgrader{
		HandshakeTimeout:  timeout,
		EnableCompression: true,
	}

	setupDone := make(chan struct{})
	ts.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatal(err)
		}
		ts.wsConn = conn

		conn.SetReadDeadline(time.Now().Add(timeout))
		conn.SetWriteDeadline(time.Now().Add(timeout))

		close(setupDone)
		defer conn.Close()

		for {
			var req Request
			if err := conn.ReadJSON(&req); err != nil {
				break
			}
			if respond != nil {
				if err := respond(conn, &req); err != nil {
					break
				}
			}
		}
	}))

	ts.conn, err = Dial("ws"+strings.TrimPrefix(ts.srv.URL, "http"), WithCompression())
	if err != nil {
		t.Fatal(err)
	}
	ts.conn.SetCompressionLevel(flate.BestSpeed)

	<-setupDone
	return ts
}

func TestConn_Invoke(t *testing.T) {
	responses := []string{
		"hello",
		"world",
	}
	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		resp := Response{
			ID:     req.ID,
			Result: []byte(fmt.Sprintf("%q", responses[int(req.ID)-1])),
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

	err = Invoke(nil, "test.World", nil, &reply, srv.conn)
	if err != nil {
		t.Error(err)
	}

	if reply != "world" {
		t.Errorf("test.World: got reply %q, want %q", reply, "world")
	}
}

func TestConn_InvokeError(t *testing.T) {
	want := "bad request"

	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		resp := Response{
			ID:    req.ID,
			Error: &ResponseError{Message: want},
		}
		return conn.WriteJSON(&resp)
	})
	defer srv.Close()

	switch err := Invoke(nil, "test.Hello", nil, nil, srv.conn).(type) {
	case *ResponseError:
		if err.Message != want {
			t.Errorf("Invoke err.Message: got %q, want %q", err.Message, want)
		}
	default:
		t.Errorf("Invoke: want *rpcError, got %#v", err)
	}
}

func TestConn_InvokeRemoteDisconnected(t *testing.T) {
	srv := newTestServer(t, nil)
	defer srv.Close()

	srv.wsConn.Close()
	err := Invoke(nil, "test.Hello", nil, nil, srv.conn)
	if err == nil {
		t.Error("Invoke error: got nil, want error")
	}
}

func TestConn_InvokeConnectionClosed(t *testing.T) {
	srv := newTestServer(t, nil)
	defer srv.Close()

	srv.conn.Close()
	err := Invoke(nil, "test.Hello", nil, nil, srv.conn)
	if err != ErrConnClosing {
		t.Errorf("Invoke error: got %v, want ErrConnClosing", err)
	}
}

func TestConn_InvokeDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	srv := newTestServer(t, nil)
	defer srv.Close()

	err := Invoke(ctx, "test.Hello", nil, nil, srv.conn)
	if err != context.DeadlineExceeded {
		t.Errorf("Invoke error: got %v, want DeadlineExceeded", err)
	}
}

func TestConn_InvokeContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		cancel()
		return nil
	})
	defer srv.Close()

	err := Invoke(ctx, "test.Hello", nil, nil, srv.conn)
	if err != context.Canceled {
		t.Errorf("Invoke error: got %v, want %v", err, context.Canceled)
	}
}

func TestConn_DecodeError(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		msg := fmt.Sprintf(`{"id": %d, "result": {}}`, req.ID)
		w, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			t.Fatal(err)
		}
		_, err = w.Write([]byte(msg))
		if err != nil {
			t.Fatal(err)
		}
		w.Close()
		return nil
	})
	defer srv.Close()

	var reply string
	err := Invoke(nil, "test.DecodeError", nil, &reply, srv.conn)
	if err == nil || !strings.HasPrefix(err.Error(), "rpcc: decoding") {
		t.Errorf("test.DecodeError: got %v, want error with %v", err, "rpcc: decoding")
	}
}

type badEncoder struct {
	ch  chan struct{}
	err error
}

func (enc *badEncoder) WriteRequest(r *Request) error  { return enc.err }
func (enc *badEncoder) ReadResponse(r *Response) error { return nil }

func TestConn_EncodeFailed(t *testing.T) {
	enc := &badEncoder{err: errors.New("fail"), ch: make(chan struct{})}
	conn := &Conn{
		ctx:     context.Background(),
		pending: make(map[uint64]*rpcCall),
		codec:   enc,
	}

	err := Invoke(nil, "test.Hello", nil, nil, conn)
	if err != enc.err {
		t.Errorf("Encode: got %v, want %v", err, enc.err)
	}
}

func TestConn_Notify(t *testing.T) {
	srv := newTestServer(t, nil)
	defer srv.Close()

	s, err := NewStream(nil, "test.Notify", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	go func() {
		resp := Response{
			Method: "test.Notify",
			Args:   []byte(`"hello"`),
		}
		if err := srv.wsConn.WriteJSON(&resp); err != nil {
			t.Fatal(err)
		}
	}()

	var reply string
	if err = s.RecvMsg(&reply); err != nil {
		t.Error(err)
	}

	want := "hello"
	if reply != want {
		t.Errorf("test.Notify reply: got %q, want %q", reply, want)
	}

	s.Close()
	if err = s.RecvMsg(nil); err == nil {
		t.Error("test.Notify read after closed: want error, got nil")
	}
}

func TestConn_StreamRecv(t *testing.T) {
	srv := newTestServer(t, nil)
	defer srv.Close()

	s, err := NewStream(nil, "test.Stream", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	messages := []Response{
		{Method: "test.Stream", Args: []byte(`"first"`)},
		{Method: "test.Stream", Args: []byte(`"second"`)},
		{Method: "test.Stream", Args: []byte(`"third"`)},
	}

	for _, m := range messages {
		if err = srv.wsConn.WriteJSON(&m); err != nil {
			t.Fatal(err)
		}
	}
	// Allow messages to propagate and trigger buffering
	// (multiple messages before Recv).
	// TODO: Remove the reliance on sleep here.
	time.Sleep(10 * time.Millisecond)

	for _, m := range messages {
		var want string
		if err = json.Unmarshal(m.Args, &want); err != nil {
			t.Error(err)
		}
		var reply string
		if err = s.RecvMsg(&reply); err != nil {
			t.Error(err)
		}
		if reply != want {
			t.Errorf("RecvMsg: got %v, want %v", reply, want)
		}
	}
}

func TestDialContext_PanicOnNilContext(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DialContext() should panic")
		}
	}()

	DialContext(nil, "")

	t.Errorf("DialContext() unreachable code after panic")
}

func TestDialContext_DeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	_, err := DialContext(ctx, "")

	// Should return deadline even when dial address is bad.
	if err != context.DeadlineExceeded {
		t.Errorf("DialContext: got %v, want %v", err, context.DeadlineExceeded)
	}
}

func TestResponse_String(t *testing.T) {
	tests := []struct {
		name string
		resp Response
		want string
	}{
		{"Notification", Response{Method: "notification", Args: []byte(`{}`)}, "Method = notification, Params = {}"},
		{"Response", Response{ID: 1, Result: []byte(`{}`)}, "ID = 1, Result = {}"},
		{"Error", Response{ID: 2, Error: &ResponseError{Code: -1000, Message: "bad request"}}, "ID = 2, Error = rpc error: bad request (code = -1000)"},
		{"Error with data", Response{ID: 3, Error: &ResponseError{Code: -1000, Message: "bad request", Data: "extra"}}, "ID = 3, Error = rpc error: bad request (code = -1000, data = extra)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.resp.String()
			if got != tt.want {
				t.Errorf("String() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	enableDebug = true
	os.Exit(m.Run())
}
