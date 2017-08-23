package rpcc

import (
	"context"
	"sort"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newTestStreamConn() (*Conn, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &Conn{
		ctx:     ctx,
		cancel:  cancel,
		streams: make(map[string]*streamClients),
	}
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

func TestStream_Ready(t *testing.T) {
	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run := func(t *testing.T, closeEarly bool) {
		s, err := NewStream(ctx, "test", conn)
		if err != nil {
			t.Fatal(err)
		}

		go func() {
			for i := 0; i < 10; i++ {
				conn.notify("test", []byte(strconv.Itoa(i)))
			}
			if closeEarly {
				s.Close()
			}

		}()

		for i := 0; i < 10; i++ {
			<-s.Ready()
			var x int
			err := s.RecvMsg(&x)
			if err != nil {
				t.Error(err)
			}
			if x != i {
				t.Errorf("x != i; got %d == %d, want %d == %d", x, i, i, i)
			}
		}

		s.Close()
		if _, ok := <-s.Ready(); ok {
			t.Errorf("s.Read(), got channel open, want channel closed")
		}
	}

	t.Run("Simple", func(t *testing.T) { run(t, false) })
	t.Run("Close early", func(t *testing.T) { run(t, true) })
}

func TestStream_Ready_Multiple_Streams(t *testing.T) {
	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s1, err := NewStream(ctx, "test1", conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s1.Close()

	s2, err := NewStream(ctx, "test2", conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()

	go func() {
		for i := 0; i < 10; i++ {
			conn.notify("test1", []byte(strconv.Itoa(i)))
			conn.notify("test2", []byte(strconv.Itoa(i)))
		}
	}()

	var got []int
	for i := 0; i < 10*2; i++ {
		var x int
		select {
		case <-s1.Ready():
			err := s1.RecvMsg(&x)
			if err != nil {
				t.Error(err)
			}
		case <-s2.Ready():
			err := s2.RecvMsg(&x)
			if err != nil {
				t.Error(err)
			}
		}
		got = append(got, x)
	}

	sort.Ints(got)
	want := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Output differs (-got +want)\n%s", diff)
	}
}

func TestStream_Wait_Blocks_After_RecvMsg(t *testing.T) {
	conn, connCancel := newTestStreamConn()
	defer connCancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewStream(ctx, "test", conn)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	// Notify will trigger a send on the ready channel.
	conn.notify("test", []byte(`"hello"`))

	var got string
	// RecvMsg should empty the ready channel so Ready can block.
	err = s.RecvMsg(&got)
	if err != nil {
		t.Error(err)
	}

	select {
	case <-s.Ready():
		t.Errorf("Ready() did not block, expected a blocking call")
	default:
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
			b.store(&message{data: []byte(strconv.Itoa(i)), next: func() { b.load() }})
		}
	}()

	i := 0
	for m := range b.get() {
		m.next()
		if strconv.Itoa(i) != string(m.data) {
			t.Errorf("Got n = %s, want %d", m.data, i)
		}
		i++
		if i >= n {
			break
		}
	}
}
