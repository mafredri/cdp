package devtool

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

type multiHandler struct {
	t *testing.T
	h []http.Handler
}

func (mh *multiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(mh.h) == 0 {
		mh.t.Error("no handlers for multiHandler")
	}
	h := mh.h[0]
	mh.h = mh.h[1:]
	h.ServeHTTP(w, r)
}

type wsHandler struct {
	t        *testing.T
	upgrader websocket.Upgrader
	message  interface{}
}

func (ws *wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.t.Fatal(err)
	}
	defer c.Close()

	var req rpcc.Request
	if err := c.ReadJSON(&req); err != nil {
		ws.t.Fatal(err)
	}

	args, err := json.Marshal(ws.message)
	if err != nil {
		ws.t.Fatal(err)
	}

	if err := c.WriteJSON(&rpcc.Response{
		ID:     req.ID,
		Result: json.RawMessage(args),
	}); err != nil {
		ws.t.Fatal(err)
	}
}

func TestDevTools_HeadlessCreateURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mh := &multiHandler{
		t: t,
		h: []http.Handler{
			&testHandler{status: 500},
			&wsHandler{t: t, message: &target.CreateTargetReply{TargetID: "abcd-abcd-abcd-abcd"}},
			&testHandler{status: 200, body: []byte(`[{"id": "abcd-abcd-abcd-abcd"}]`)},
		},
	}

	srv := httptest.NewServer(mh)
	defer srv.Close()

	devt := New(srv.URL)
	targ, err := devt.Create(ctx)
	if err != nil {
		t.Fatal(err)
	}

	want := "abcd-abcd-abcd-abcd"
	if targ.ID != want {
		t.Errorf("Target.ID: got %q, want %q", targ.ID, want)
	}
}
