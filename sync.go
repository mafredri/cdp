package cdp

import (
	"fmt"

	"github.com/mafredri/cdp/rpcc"
)

type eventClient interface {
	rpcc.Stream
}

type getStreamer interface {
	GetStream() rpcc.Stream
}

// Sync takes two or more event clients and sets them into synchronous operation,
// relative to each other. This operation cannot be undone. If an error is
// returned this function is no-op and the event clients will continue in
// asynchronous operation.
//
// All event clients must belong to the same connection and they must not be
// closed. Passing multiple clients of the same event type to Sync is not
// supported and will return an error.
//
// An event client that is closed is removed and has no further affect on the
// clients that were synchronized.
//
// When two event clients, A and B, are in sync they will receive events in the
// order of arrival. If an event for both A and B is triggered, in that order,
// it will not be possible to receive the event from B before the event from A
// has been received.
func Sync(c ...eventClient) error {
	var s []rpcc.Stream
	for _, cc := range c {
		cs, ok := cc.(getStreamer)
		if !ok {
			return fmt.Errorf("cdp: Sync: bad eventClient type: %T", cc)
		}
		s = append(s, cs.GetStream())
	}
	return rpcc.Sync(s...)
}
