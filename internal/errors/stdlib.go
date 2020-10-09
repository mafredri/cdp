// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	stderrors "errors"
	"fmt"
)

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error { return stderrors.New(text) }

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error { return stderrors.Unwrap(err) }

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == os.ErrExist }
//
// then Is(MyError{}, os.ErrExist) returns true. See syscall.Errno.Is for
// an example in the standard library.
func Is(err, target error) bool { return stderrors.Is(err, target) }

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
func As(err error, target interface{}) bool { return stderrors.As(err, target) }

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand. It is
// invalid to include more than one %w verb or to supply it with an operand
// that does not implement the error interface. The %w verb is otherwise
// a synonym for %v.
func Errorf(format string, a ...interface{}) error { return fmt.Errorf(format, a...) }
