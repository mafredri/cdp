package cdp

import (
	"github.com/mafredri/cdp/rpcc"
)

type eventClient interface {
	rpcc.Stream
}

// Sync synchronizes two or more event clients.
// An error will be returned if the event clients do not have the same
// connection or event clients of the same type are synchronized.
func Sync(c ...eventClient) error {
	var s []rpcc.Stream
	for _, cc := range c {
		s = append(s, cc)
	}
	return rpcc.Sync(s...)
}
