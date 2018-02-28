package errors

import (
	"errors"
	"fmt"
	"strings"
)

type causer interface {
	Cause() error
}

// Cause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
func Cause(err error) error {
	for err != nil {
		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return err
		}
	}
	return err
}

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}

// Errorf wraps New and fmt.Sprintf.
func Errorf(format string, a ...interface{}) error {
	return New(fmt.Sprintf(format, a...))
}

// Wrapf wraps an error with a message. Wrapf returns nil if error is nil.
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrapped{
		err: err,
		msg: fmt.Sprintf(format, a...),
	}
}

type wrapped struct {
	err error
	msg string
}

var _ error = (*wrapped)(nil)
var _ causer = (*wrapped)(nil)

func (e *wrapped) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err)
}

func (e *wrapped) Cause() error {
	return e.err
}

// Merge merges multiple errors into one.
// Merge returns nil if all errors are nil.
func Merge(err ...error) error {
	var errs []error
	for _, e := range err {
		if e != nil {
			errs = append(errs, e)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return &merged{s: err}
}

type merged struct {
	s []error
}

var _ error = (*merged)(nil)

func (e *merged) Error() string {
	var m []string
	for _, err := range e.s {
		m = append(m, err.Error())
	}
	return strings.Join(m, ": ")
}
