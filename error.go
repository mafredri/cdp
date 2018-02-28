package cdp

import (
	"github.com/mafredri/cdp/internal/errors"
)

// ErrorCause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
func ErrorCause(err error) error { return errors.Cause(err) }
