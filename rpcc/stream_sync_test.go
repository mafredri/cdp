package rpcc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSync(t *testing.T) {
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

	// These notifications should disappear after Sync.
	conn.notify("test1", []byte(strconv.Itoa(1)))
	conn.notify("test1", []byte(strconv.Itoa(2)))

	err = Sync(s1, s2)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for i := 0; i < 3; i++ {
			conn.notify("test1", []byte(strconv.Itoa(100+i)))
		}
		for i := 0; i < 3; i++ {
			conn.notify("test2", []byte(strconv.Itoa(200+i)))
		}
		for i := 0; i < 4; i++ {
			conn.notify("test1", []byte(strconv.Itoa(100+i)))
			conn.notify("test2", []byte(strconv.Itoa(200+i)))
		}
	}()

	var got []int
	for i := 0; i < 14; i++ {
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

	want := []int{100, 101, 102, 200, 201, 202, 100, 200, 101, 201, 102, 202, 103, 203}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Output differs (-got +want)\n%s", diff)
	}
}

type fakeStream struct{}

func (s *fakeStream) Ready() <-chan struct{}      { return nil }
func (s *fakeStream) RecvMsg(m interface{}) error { return nil }
func (s *fakeStream) Close() error                { return nil }

var _ Stream = (*fakeStream)(nil)

func TestSyncError(t *testing.T) {
	conn1, connCancel1 := newTestStreamConn()
	defer connCancel1()

	conn2, connCancel2 := newTestStreamConn()
	defer connCancel2()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s1, err := NewStream(ctx, "test", conn1)
	if err != nil {
		t.Fatal(err)
	}
	defer s1.Close()

	s2, err := NewStream(ctx, "duplicate", conn1)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()

	s3, err := NewStream(ctx, "duplicate", conn1)
	if err != nil {
		t.Fatal(err)
	}
	defer s3.Close()

	s4, err := NewStream(ctx, "other-conn", conn2)
	if err != nil {
		t.Fatal(err)
	}
	defer s4.Close()

	err = Sync(s1, s2)
	if err != nil {
		t.Errorf("Sync failed: %v", err)
	}

	tests := []struct {
		name string
		run  func() error
	}{
		{"Single stream", func() error { return Sync(s1) }},
		{"Invalid stream", func() error { return Sync(&fakeStream{}, s1) }},
		{"Duplicate stream", func() error { return Sync(s1, s2, s3) }},
		{"Mixed Conn", func() error { return Sync(s1, s2, s4) }},
		{"Closed Conn", func() error {
			conn2.Close()
			return Sync(s4, s1)
		}},
		{"Closed Stream", func() error {
			s3.Close()
			return Sync(s3, s1)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.run()
			if err == nil {
				t.Error("Expected error, got nil")
			}
		})
	}
}

func TestStreamSyncNotifyDeadlock(t *testing.T) {
	conn, cancel := newTestStreamConn()
	defer cancel()

	ctx := context.Background()

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

	syncErr := make(chan error)
	go func() {
		// This could cause a deadlock due to competition for same mutexes:
		// https://github.com/mafredri/cdp/issues/90
		// https://github.com/mafredri/cdp/pull/91
		syncErr <- Sync(s1, s2)
	}()
	for i := 0; i < 10; i++ {
		conn.notify("test1", []byte(fmt.Sprintf(`{"hello": "world", "count": %d}`, i)))
		conn.notify("test2", []byte(fmt.Sprintf(`{"hello": "world", "count": %d}`, i)))
	}
	err = <-syncErr
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamSyncSameStreamDeadlock(t *testing.T) {
	conn, cancel := newTestStreamConn()
	defer cancel()

	ctx := context.Background()

	s1, err := NewStream(ctx, "test1", conn)
	if err != nil {
		t.Fatal(err)
	}
	s2, err := NewStream(ctx, "test2", conn)
	if err != nil {
		t.Fatal(err)
	}

	err = Sync(s1, s2, s1)
	if err == nil {
		t.Error("Same stream passed multiple times: want error, got nil")
	}
}

// This test is partially an observability test, order can be tweaked to try to
// catch data races in Streams using syncMessageStore. It also tries to catch
// issues with missing messages.
func TestStreamSyncClosingStreams(t *testing.T) {
	conn, cancel := newTestStreamConn()
	defer cancel()

	ctx := context.Background()

	testStreams := []struct {
		name string
	}{
		{name: "test0"},
		{name: "test1"},
		{name: "test2"},
		{name: "test3"},
		{name: "test4"},
	}

	var streams []Stream
	for _, s := range testStreams {
		ss, err := NewStream(ctx, s.name, conn)
		if err != nil {
			t.Fatal(err)
		}
		streams = append(streams, ss)
	}

	readySteadyGo := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-readySteadyGo
		go streams[2].Close()
		conn.notify("test0", []byte("test0.0"))
		conn.notify("test1", []byte("test1.0"))
		conn.notify("test2", []byte("test2.0"))
		conn.notify("test2", []byte("test2.1"))
		conn.notify("test2", []byte("test2.2"))
		conn.notify("test3", []byte("test3.0"))
		conn.notify("test3", []byte("test3.1"))
		streams[3].Close()
		conn.notify("test4", []byte("test4.0"))
	}()

	for i, s := range streams {
		i := i
		s := s
		wg.Add(1)
		go func() {
			defer wg.Done()

			<-readySteadyGo
			<-s.Ready()
			var m []byte
			err := s.RecvMsg(&m)
			if i != 2 && i != 3 && err != nil {
				t.Error(err)
			}
			sm := string(m)
			if i != 2 && i != 3 {
				if sm == "" {
					t.Error("missing message", i)
				} else if !strings.HasSuffix(sm, ".0") {
					t.Error("wrong message", i, sm)
				}
			}
			t.Log(sm)
		}()
	}

	err := Sync(streams...)
	if err != nil {
		t.Fatal(err)
	}

	close(readySteadyGo)

	wg.Wait()
	t.Log("The end.")
}
