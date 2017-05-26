package cdpdom

import (
	"fmt"
)

// opError represents an operational error.
type opError struct {
	Domain string
	Op     string
	Err    error
}

// Cause implements error causer.
func (e *opError) Cause() error {
	return e.Err
}

func (e opError) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

type causer interface {
	Cause() error
}

var (
	_ error  = (*opError)(nil)
	_ causer = (*opError)(nil)
)
