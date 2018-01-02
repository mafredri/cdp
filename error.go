package cdp

type causer interface {
	Cause() error
}

// ErrorCause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
func ErrorCause(err error) error {
	for err != nil {
		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return err
		}
	}
	return err
}
