package io

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"strings"
	"time"
)

// StreamReader represents a stream reader.
type StreamReader struct {
	next  func(pos, size int) (*ReadReply, error)
	close func() error
	r     io.Reader
	pos   int
	eof   bool
}

var _ io.ReadCloser = (*StreamReader)(nil)

// NewStreamReader returns a reader for io.Streams that implements io.Reader
// from the standard library.
func NewStreamReader(ctx context.Context, ioClient interface {
	Read(context.Context, *ReadArgs) (*ReadReply, error)
	Close(context.Context, *CloseArgs) error
}, handle StreamHandle,
) *StreamReader {
	args := NewReadArgs(handle)
	return &StreamReader{
		next: func(pos, size int) (*ReadReply, error) {
			args.SetOffset(pos).SetSize(size)
			return ioClient.Read(ctx, args)
		},
		close: func() error {
			// TODO(mafredri): We should ideally allow the user to define a timeout here.
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			return ioClient.Close(ctx, NewCloseArgs(handle))
		},
	}
}

func (r *StreamReader) read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.pos += n
	if !r.eof && err == io.EOF {
		r.r = nil
		err = nil
	}
	return n, err
}

// Read a chunk of the stream.
func (r *StreamReader) Read(p []byte) (n int, err error) {
	if r.r != nil {
		// Continue reading from buffer.
		return r.read(p)
	}
	if r.eof {
		return 0, io.EOF
	}
	if len(p) == 0 {
		return 0, nil
	}

	// Chrome might have an off-by-one when deciding the maximum
	// size (at least for base64 encoded data), usually it will
	// overflow. We subtract one to make sure it fits into p.
	size := len(p) - 1
	if size < 1 {
		// Safety-check to avoid crashing Chrome (e.g. via SetSize(-1)).
		size = 1
	}

	reply, err := r.next(r.pos, size)
	if err != nil {
		return 0, err
	}

	r.eof = reply.EOF

	switch {
	case reply.Base64Encoded != nil && *reply.Base64Encoded:
		b := []byte(reply.Data)
		size := base64.StdEncoding.DecodedLen(len(b))

		// Safety-check for fast-path to avoid panics.
		if len(p) >= size {
			n, err = base64.StdEncoding.Decode(p, b)
			r.pos += n
			return n, err
		}

		r.r = base64.NewDecoder(base64.StdEncoding, bytes.NewReader(b))
	default:
		r.r = strings.NewReader(reply.Data)
	}

	return r.read(p)
}

// Close the stream, discard any temporary backing storage.
func (r *StreamReader) Close() error {
	return r.close()
}
