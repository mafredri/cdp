package io

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newReply(data string, base64, eof bool) *ReadReply {
	return &ReadReply{
		Data:          data,
		Base64Encoded: &base64,
		EOF:           eof,
	}
}

func TestStreamReader_base64Decode(t *testing.T) {
	want := "Hello world"
	r := &StreamReader{
		next: func(pos, size int) (*ReadReply, error) {
			return newReply(base64.StdEncoding.EncodeToString([]byte(want)), true, true), nil
		},
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	got := string(b)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStreamReader_string(t *testing.T) {
	want := "Hello world"
	r := &StreamReader{
		next: func(pos, size int) (*ReadReply, error) {
			return newReply(want, false, true), nil
		},
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	got := string(b)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStreamReader_replyTooBig(t *testing.T) {
	want := "helloworld"
	data := [][]byte{
		[]byte("hello"),
		[]byte("world"),
	}
	r := &StreamReader{
		next: func(pos, size int) (*ReadReply, error) {
			if len(data) == 0 {
				return newReply("", false, true), nil
			}

			eof := false
			d := data[0]
			data = data[1:]
			if len(data) == 0 {
				eof = true
			}
			return newReply(base64.StdEncoding.EncodeToString(d), true, eof), nil
		},
	}

	b := make([]byte, 1)
	var got string
	for {
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Error(err)
		}
		got += string(b[:n])
	}
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

type streamReaderTestClient struct {
	reply []ReadReply
	close error
}

func (c *streamReaderTestClient) Read(context.Context, *ReadArgs) (*ReadReply, error) {
	if len(c.reply) == 0 {
		return nil, io.EOF
	}
	reply := c.reply[0]
	c.reply = c.reply[1:]
	return &reply, nil
}
func (c *streamReaderTestClient) Close(context.Context, *CloseArgs) error {
	return c.close
}

func TestNewStreamReader(t *testing.T) {
	c := &streamReaderTestClient{
		reply: []ReadReply{
			ReadReply{Data: "hello ", EOF: false},
			ReadReply{Data: "world!", EOF: true},
		},
		close: nil,
	}
	r := NewStreamReader(context.Background(), c, "")
	if r == nil {
		t.Error("want reader, got nil")
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	want := "hello world!"
	got := string(b)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("StreamReader.Read() failed (-want +got)\n%s", diff)
	}

	err = r.Close()
	if err != nil {
		t.Error(err)
	}
}
