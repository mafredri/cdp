package rpcc

import (
	"context"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestStream_RecvMsg(t *testing.T) {
	params := struct {
		arg1  string
		want1 string
		arg2  []byte
		want2 []byte
		arg3  string
		want3 string
	}{
		want1: "hello",
		want2: []byte(`"raw"`),
		want3: "",
	}
	type fields struct {
		payload string
	}
	type args struct {
		m interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Receive string as string", fields{`"hello"`}, args{&params.arg1}, &params.want1, false},
		{"Receive []byte as raw []byte", fields{`"raw"`}, args{&params.arg2}, &params.want2, false},
		{"Receive int as string error", fields{`42`}, args{&params.arg3}, &params.want3, true},
	}

	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test", conn)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn.notify("test", []byte(tt.fields.payload))

			if err := s.RecvMsg(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Stream.RecvMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(tt.args.m, tt.want) {
				t.Errorf("Stream.RecvMsg(): got %#v, want %#v", tt.args.m, tt.want)
			}
		})
	}
}

func TestMessageBuffer(t *testing.T) {
	n := 1000
	b := newMessageBuffer()

	go func() {
		for i := 0; i < n; i++ {
			b.store(&streamMsg{data: []byte(strconv.Itoa(i))})
		}
	}()

	i := 0
	for bi := range b.get() {
		b.load()
		if strconv.Itoa(i) != string(bi.data) {
			t.Errorf("Got n = %s, want %d", bi, i)
		}
		i++
		if i >= n {
			break
		}
	}
}
