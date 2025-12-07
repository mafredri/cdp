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
	"sync"
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
	if ts.srv != nil {
		defer ts.srv.Close()
	}
	if ts.wsConn == nil {
		return nil
	}
	return ts.wsConn.Close()
}

func newTestServer(t testing.TB, respond func(*websocket.Conn, *Request) error) *testServer {
	t.Helper()

	// Timeouts to prevent tests from running forever.
	timeout := 5 * time.Second

	var err error
	ts := &testServer{}
	t.Cleanup(func() {
		ts.Close()
	})
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

		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			t.Fatal(err)
		}
		err = conn.SetWriteDeadline(time.Now().Add(timeout))
		if err != nil {
			t.Fatal(err)
		}

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
	err = ts.conn.SetCompressionLevel(flate.BestSpeed)
	if err != nil {
		t.Error(err)
	}

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var reply string
	err := Invoke(ctx, "test.Hello", nil, &reply, srv.conn)
	if err != nil {
		t.Error(err)
	}

	if reply != "hello" {
		t.Errorf("test.Hello: got reply %q, want %q", reply, "hello")
	}

	err = Invoke(ctx, "test.World", nil, &reply, srv.conn)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch err := Invoke(ctx, "test.Hello", nil, nil, srv.conn).(type) {
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv.wsConn.Close()
	err := Invoke(ctx, "test.Hello", nil, nil, srv.conn)
	if err == nil {
		t.Error("Invoke error: got nil, want error")
	}
}

func TestConn_InvokeConnectionClosed(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv.conn.Close()
	err := Invoke(ctx, "test.Hello", nil, nil, srv.conn)
	if err != ErrConnClosing {
		t.Errorf("Invoke error: got %v, want ErrConnClosing", err)
	}
}

func TestConn_InvokeDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	srv := newTestServer(t, nil)

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var reply string
	err := Invoke(ctx, "test.DecodeError", nil, &reply, srv.conn)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := Invoke(ctx, "test.Hello", nil, nil, conn)
	if err != enc.err {
		t.Errorf("Encode: got %v, want %v", err, enc.err)
	}
}

func TestConn_Notify(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test.Notify", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	errc := make(chan error, 1)
	go func() {
		resp := Response{
			Method: "test.Notify",
			Args:   []byte(`"hello"`),
		}
		errc <- srv.wsConn.WriteJSON(&resp)
	}()

	var reply string
	if err = s.RecvMsg(&reply); err != nil {
		t.Error(err)
	}
	if err = <-errc; err != nil {
		t.Fatal(err)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test.Stream", srv.conn)
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

func TestConn_PropagateError(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s1, err := NewStream(ctx, "test.Stream1", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s1.Close()
	s2, err := NewStream(ctx, "test.Stream2", srv.conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()

	errC := make(chan error, 2)
	go func() {
		errC <- Invoke(ctx, "test.Invoke", nil, nil, srv.conn)
	}()
	go func() {
		var reply string
		errC <- s1.RecvMsg(&reply)
	}()

	// Give a little time for both Invoke & Recv.
	time.Sleep(5 * time.Millisecond)

	srv.wsConn.Close()

	// Give a little time for connection to close.
	time.Sleep(5 * time.Millisecond)

	lastErr := Invoke(ctx, "test.Invoke", nil, nil, srv.conn)
	if lastErr == nil {
		t.Error("RecvMsg on closed connection: got nil, want an error")
	}

	var reply string
	err = s2.RecvMsg(&reply)
	if err != lastErr {
		t.Errorf("Error was not repeated, got %v, want %v", err, lastErr)
	}

	for i := 0; i < 2; i++ {
		err := <-errC
		if err != lastErr {
			t.Errorf("Error was not repeated, got %v, want %v", err, lastErr)
		}
	}
}

func TestConn_Context(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx := srv.conn.Context()
	if ctx == nil {
		t.Fatal("Context is nil")
	}

	srv.conn.Close()
	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		t.Error("Timeout waiting for context to be done")
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

func TestConn_NewSession(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		resp := Response{
			ID:        req.ID,
			Result:    []byte(fmt.Sprintf("%q", req.Method)),
			SessionID: req.SessionID,
		}
		return conn.WriteJSON(&resp)
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess, err := NewSession(srv.conn, "session-123")
	if err != nil {
		t.Fatal(err)
	}
	defer sess.Close()

	if sess.SessionID() != "session-123" {
		t.Errorf("SessionID: got %q, want %q", sess.SessionID(), "session-123")
	}

	if srv.conn.SessionID() != "" {
		t.Errorf("Parent SessionID: got %q, want empty", srv.conn.SessionID())
	}

	var reply string
	err = Invoke(ctx, "test.SessionMethod", nil, &reply, sess)
	if err != nil {
		t.Fatal(err)
	}

	if reply != "test.SessionMethod" {
		t.Errorf("Session reply: got %q, want %q", reply, "test.SessionMethod")
	}
}

func TestConn_NewSession_Response(t *testing.T) {
	srv := newTestServer(t, func(conn *websocket.Conn, req *Request) error {
		resp := Response{
			ID:        req.ID,
			Result:    []byte(fmt.Sprintf("%q", req.SessionID)),
			SessionID: req.SessionID,
		}
		return conn.WriteJSON(&resp)
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess1, err := NewSession(srv.conn, "session-1")
	if err != nil {
		t.Fatal(err)
	}
	defer sess1.Close()

	sess2, err := NewSession(srv.conn, "session-2")
	if err != nil {
		t.Fatal(err)
	}
	defer sess2.Close()

	var wg sync.WaitGroup
	results := make([]string, 2)
	for i, sess := range []*Conn{sess1, sess2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var reply string
			err := Invoke(ctx, "test.Method", nil, &reply, sess)
			if err != nil {
				t.Error(err)
				return
			}
			results[i] = reply
		}()
	}

	wg.Wait()

	if got := results[0]; got != "session-1" {
		t.Errorf("Session-1 got response for %q", got)
	}
	if got := results[1]; got != "session-2" {
		t.Errorf("Session-2 got response for %q", got)
	}
}

func TestConn_NewSession_Notification(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess1, err := NewSession(srv.conn, "session-1")
	if err != nil {
		t.Fatal(err)
	}
	defer sess1.Close()

	sess2, err := NewSession(srv.conn, "session-2")
	if err != nil {
		t.Fatal(err)
	}
	defer sess2.Close()

	stream1, err := NewStream(ctx, "test.Event", sess1)
	if err != nil {
		t.Fatal(err)
	}
	defer stream1.Close()

	stream2, err := NewStream(ctx, "test.Event", sess2)
	if err != nil {
		t.Fatal(err)
	}
	defer stream2.Close()

	err = srv.wsConn.WriteJSON(&Response{
		Method:    "test.Event",
		Args:      []byte(`"event-for-session-1"`),
		SessionID: "session-1",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = srv.wsConn.WriteJSON(&Response{
		Method:    "test.Event",
		Args:      []byte(`"event-for-session-2"`),
		SessionID: "session-2",
	})
	if err != nil {
		t.Fatal(err)
	}

	var reply1 string
	if err := stream1.RecvMsg(&reply1); err != nil {
		t.Fatal(err)
	}
	if reply1 != "event-for-session-1" {
		t.Errorf("Session-1 event: got %q, want %q", reply1, "event-for-session-1")
	}

	var reply2 string
	if err := stream2.RecvMsg(&reply2); err != nil {
		t.Fatal(err)
	}
	if reply2 != "event-for-session-2" {
		t.Errorf("Session-2 event: got %q, want %q", reply2, "event-for-session-2")
	}
}

func TestConn_NewSession_CloseFunc(t *testing.T) {
	srv := newTestServer(t, nil)

	called := make(chan struct{})
	sess, err := NewSession(srv.conn, "session-123", WithSessionClose(func(ctx context.Context) error {
		close(called)
		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}

	err = sess.Close()
	if err != nil {
		t.Errorf("Session.Close error: %v", err)
	}

	select {
	case <-called:
		// OK.
	case <-time.After(time.Second):
		t.Error("Close callback was not called")
	}
}

func TestConn_NewSession_ParentCloseClosesSession(t *testing.T) {
	srv := newTestServer(t, nil)

	called := make(chan struct{})
	sess, err := NewSession(srv.conn, "session-123", WithSessionClose(func(ctx context.Context) error {
		close(called)
		return nil
	}))
	if err != nil {
		srv.Close()
		t.Fatal(err)
	}

	err = srv.conn.Close()
	if err != nil {
		t.Errorf("Parent.Close error: %v", err)
	}
	srv.srv.Close()

	select {
	case <-called:
		// OK.
	case <-time.After(time.Second):
		t.Error("Session close callback was not called when parent closed")
	}

	select {
	case <-sess.Context().Done():
	default:
		t.Error("Session context should be done after parent close")
	}
}

func TestConn_NewSession_CannotCreateFromSession(t *testing.T) {
	srv := newTestServer(t, nil)

	sess, err := NewSession(srv.conn, "session-123")
	if err != nil {
		t.Fatal(err)
	}
	defer sess.Close()

	_, err = NewSession(sess, "nested-session")
	if err == nil {
		t.Error("NewSession on session should return error")
	}
}

func TestConn_NewSession_ClosedParent(t *testing.T) {
	srv := newTestServer(t, nil)

	srv.conn.Close()
	srv.srv.Close()

	_, err := NewSession(srv.conn, "session-123")
	if err == nil {
		t.Error("NewSession on closed parent should return error")
	}
}

func TestConn_NewSession_InvokeAfterClose(t *testing.T) {
	srv := newTestServer(t, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess, err := NewSession(srv.conn, "session-123")
	if err != nil {
		t.Fatal(err)
	}

	sess.Close()

	var reply string
	err = Invoke(ctx, "test.Method", nil, &reply, sess)
	if err != ErrConnClosing {
		t.Errorf("Invoke after close: got %v, want ErrConnClosing", err)
	}
}

func TestConn_NewSession_ConcurrentCreateAndClose(t *testing.T) {
	srv := newTestServer(t, nil)

	const numGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines + 1)

	res := make(chan error, numGoroutines)
	for i := range numGoroutines {
		go func() {
			defer wg.Done()
			sess, err := NewSession(srv.conn, fmt.Sprintf("session-%d", i))
			if err != nil {
				res <- err
				return
			}
			sess.Close()
			res <- nil
		}()
	}

	go func() {
		defer wg.Done()
		srv.conn.Close()
		srv.srv.Close()
	}()

	wg.Wait()
	close(res)

	for err := range res {
		if err != nil && err != ErrConnClosing {
			t.Errorf("NewSession returned unexpected error: %v", err)
		}
	}
}

func TestMain(m *testing.M) {
	enableDebug = true
	os.Exit(m.Run())
}
