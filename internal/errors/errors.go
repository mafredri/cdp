package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Interfaces for common error unwrapping.
type causer interface{ Cause() error }
type wrapper interface{ Unwrap() error }

// Cause returns the underlying cause for this error, if possible.
// If err does not implement causer.Cause(), then err is returned.
//
// Deprecated: Use errors.Unwrap, errors.Is or errors.As instead.
func Cause(err error) error {
	for err != nil {
		if c, ok := err.(wrapper); ok {
			err = c.Unwrap()
		} else if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return err
		}
	}
	return err
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
var _ wrapper = (*wrapped)(nil)

func (e *wrapped) Error() string { return fmt.Sprintf("%s: %s", e.msg, e.err.Error()) }
func (e *wrapped) Cause() error  { return e.err }
func (e *wrapped) Unwrap() error { return e.err }

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
	return &merged{s: errs}
}

type merged struct {
	s []error
}

var _ error = (*merged)(nil)
var _ causer = (*merged)(nil)
var _ wrapper = (*merged)(nil)

func (e *merged) Error() string {
	var m []string
	for _, err := range e.s {
		m = append(m, err.Error())
	}
	return strings.Join(m, ": ")
}

// Unwrap returns only the first error, there is
// no way to create a queue of errors.
func (e *merged) Unwrap() error { return e.s[0] }

// Cause returns only the first error, there is
// no way to create a queue of errors.
func (e *merged) Cause() error { return e.s[0] }

// Is runs errors.Is on all merged errors.
func (e *merged) Is(target error) bool {
	if target == nil {
		return nil == e.s
	}
	for _, err := range e.s {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// As runs errors.As on all merged errors.
func (e *merged) As(target interface{}) bool {
	for _, err := range e.s {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}
