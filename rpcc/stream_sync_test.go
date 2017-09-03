package rpcc

import (
	"context"
	"strconv"
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

var (
	_ Stream = (*fakeStream)(nil)
)

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
