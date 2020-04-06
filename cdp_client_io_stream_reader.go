package cdp

import (
	"context"

	"github.com/mafredri/cdp/protocol/io"
)

// NewIOStreamReader returns a reader for io.Stream that implements io.Reader
// from the standard library.
func (c *Client) NewIOStreamReader(ctx context.Context, handle io.StreamHandle) *io.StreamReader {
	return io.NewStreamReader(ctx, c.IO, handle)
}
