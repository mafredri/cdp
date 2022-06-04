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

func (e OpError) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

// Cause implements error causer.
func (e *OpError) Cause() error {
	return e.Err
}

// Unwrap implements Wrapper.
func (e *OpError) Unwrap() error {
	return e.Err
}

type (
	causer  interface{ Cause() error }
	wrapper interface{ Unwrap() error }
)

var (
	_ error   = (*OpError)(nil)
	_ causer  = (*OpError)(nil)
	_ wrapper = (*OpError)(nil)
)
