package cdp

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"testing"

	cdpio "github.com/mafredri/cdp/protocol/io"
)

func newReply(data string, base64, eof bool) *cdpio.ReadReply {
	return &cdpio.ReadReply{
		Data:          data,
		Base64Encoded: &base64,
		EOF:           eof,
	}
}

func TestIOStreamReader_base64Decode(t *testing.T) {
	want := "Hello world"
	r := &IOStreamReader{
		next: func(pos, size int) (*cdpio.ReadReply, error) {
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

func TestIOStreamReader_string(t *testing.T) {
	want := "Hello world"
	r := &IOStreamReader{
		next: func(pos, size int) (*cdpio.ReadReply, error) {
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

func TestIOStreamReader_replyTooBig(t *testing.T) {
	want := "helloworld"
	data := [][]byte{
		[]byte("hello"),
		[]byte("world"),
	}
	r := &IOStreamReader{
		next: func(pos, size int) (*cdpio.ReadReply, error) {
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

func TestNewIOStreamReader(t *testing.T) {
	r := NewIOStreamReader(context.Background(), nil, "")
	if r == nil {
		t.Error("want reader, got nil")
	}
}
