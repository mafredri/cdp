package cdp

import (
	"github.com/mafredri/cdp/cdpdom"
)

// ErrorCauser returns the error that caused this error.
// Returns err if it is not a cdpdom OpError.
func ErrorCauser(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*cdpdom.OpError); ok {
		return err.Err
	}
	return err
}
