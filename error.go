package cdp

import (
	"github.com/mafredri/cdp/internal/errors"
)

// ErrorCause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
//
// Deprecated: Use errors.Unwrap, errors.Is or errors.As instead.
func ErrorCause(err error) error { return errors.Cause(err) }
