package cdp

import "fmt"

type cdpError struct {
	Domain string
	Op     string
	Err    error
}

func (e cdpError) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

var (
	_ error = (*cdpError)(nil)
)

// ErrorCauser returns the error that caused this
// error. Returns itself if it is not a cdpError.
func ErrorCauser(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*cdpError); ok {
		return err.Err
	}
	return err
}
