package cdp

type causer interface {
	Cause() error
}

// ErrorCauser returns the error that caused this error.
// Returns err if it is not a cdpdom OpError.
func ErrorCauser(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(causer); ok {
		return err.Cause()
	}
	return err
}
