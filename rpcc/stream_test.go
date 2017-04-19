package rpcc

import (
	"context"
	"strconv"
	"testing"
)

func newTestStreamConn() (*Conn, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &Conn{ctx: ctx, streams: make(map[string]*streamClients)}
	return conn, cancel
}

func TestNewStream_AfterClose(t *testing.T) {
	srv := newTestServer(t, nil)
	defer srv.Close()

	srv.conn.Close()
	_, err := NewStream(nil, "test", srv.conn)
	if err != ErrConnClosing {
		t.Errorf("NewStream() after closed conn; got %v, want %v", err, ErrConnClosing)
	}
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

func TestMessageBuffer(t *testing.T) {
	n := 1000
	b := &messageBuffer{
		ch: make(chan []byte, 1),
	}

	go func() {
		for i := 0; i < n; i++ {
			b.store([]byte(strconv.Itoa(i)))
		}
	}()

	i := 0
	for bi := range b.get() {
		b.load()
		if strconv.Itoa(i) != string(bi) {
			t.Errorf("Got n = %s, want %d", bi, i)
		}
		i++
		if i >= n {
			break
		}
	}
}
