package cdp

type causer interface {
	Cause() error
}

// ErrorCause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
func ErrorCause(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(causer); ok {
		return err.Cause()
	}
	return err
}
