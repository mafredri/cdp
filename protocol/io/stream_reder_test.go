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

func TestNewStreamReader_Read(t *testing.T) {
	trueBool := true

	type args struct {
		ctx    context.Context
		reply  []ReadReply
		handle StreamHandle
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hello world",
			args: args{
				ctx: context.Background(),
				reply: []ReadReply{
					{Data: "Hello ", EOF: false},
					{Data: "world!", EOF: true},
				},
				handle: "",
			},
			want: "Hello world!",
		},
		{
			name: "Hello world empty EOF",
			args: args{
				ctx: context.Background(),
				reply: []ReadReply{
					{Data: "Hello ", EOF: false},
					{Data: "world!", EOF: false},
					{Data: "", EOF: true},
				},
			},
			want: "Hello world!",
		},
		{
			name: "Hello world base64 encoded",
			args: args{
				ctx: context.Background(),
				reply: []ReadReply{
					{
						Data:          base64.StdEncoding.EncodeToString([]byte("Hello world!")),
						Base64Encoded: &trueBool,
						EOF:           false,
					},
					{Data: "", EOF: true},
				},
			},
			want: "Hello world!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &streamReaderTestClient{
				reply: tt.args.reply,
				close: nil,
			}

			r := NewStreamReader(tt.args.ctx, c, tt.args.handle)

			b, err := ioutil.ReadAll(r)
			if err != nil {
				t.Error(err)
			}

			got := string(b)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("StreamReader.Read() failed (-want +got)\n%s", diff)
			}

			err = r.Close()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
