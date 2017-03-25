package cdpdom

import (
	"fmt"
)

// OpError represents an operational error.
type OpError struct {
	Domain string
	Op     string
	Err    error
}

func (e OpError) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

var (
	_ error = (*OpError)(nil)
)
