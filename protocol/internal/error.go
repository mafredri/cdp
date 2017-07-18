package internal

import (
	"fmt"
)

// OpError represents an operational error.
type OpError struct {
	Domain string
	Op     string
	Err    error
}

// Cause implements error causer.
func (e *OpError) Cause() error {
	return e.Err
}

func (e OpError) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

type causer interface {
	Cause() error
}

var (
	_ error  = (*OpError)(nil)
	_ causer = (*OpError)(nil)
)
