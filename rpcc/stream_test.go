package rpcc

import (
	"context"
	"testing"
)

func newTestStreamConn() (*Conn, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &Conn{ctx: ctx, streams: make(map[string]*streamClients)}
	return conn, cancel
}

func TestStream_UserCancel(t *testing.T) {
	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test", conn)
	if err != nil {
		t.Fatal(err)
	}

	conn.notify("test", []byte(`"message"`))

	connCancel()
	cancel() // User cancellation has priority.

	err = s.RecvMsg(nil)
	if err != ctx.Err() {
		t.Errorf("err != ctx.Err(); got %v, want %v", err, ctx.Err())
	}
}

func TestStream_RecvAfterConnClose(t *testing.T) {
	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test", conn)
	if err != nil {
		t.Fatal(err)
	}

	conn.notify("test", []byte(`"message1"`))
	conn.notify("test", []byte(`"message2"`))
	conn.notify("test", []byte(`"message3"`))

	connCancel()

	for i := 0; i < 3; i++ {
		var reply string
		err = s.RecvMsg(&reply)
		if err != nil {
			t.Error(err)
		}
	}

	err = s.RecvMsg(nil)
	if err != ErrConnClosing {
		t.Errorf("err got %v, want ErrConnClosing", err)
	}
}
