package session

import (
	"fmt"
	"strings"
)

var (
	_ error = (*serror)(nil)
	_ error = (multiError)(nil)
)

func wrapf(err error, format string, a ...interface{}) error {
	return &serror{
		err: err,
		msg: fmt.Sprintf(format, a...),
	}
}

type serror struct {
	err error
	msg string
}

func (e *serror) Error() string {
	if e.err == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %s", e.msg, e.err)
}

func (e *serror) Causer() error {
	return e.err
}

type multiError []error

func (e multiError) Error() string {
	var m []string
	for _, err := range e {
		m = append(m, err.Error())
	}
	return strings.Join(m, ": ")
}
