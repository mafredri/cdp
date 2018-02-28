package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestErrorf(t *testing.T) {
	got := Errorf("%s%s", "a", "b")
	if !strings.Contains(got.Error(), "ab") {
		t.Errorf("error does not contain %q, got %v", "ab", got)
	}
}

func TestWrap(t *testing.T) {
	err := errors.New("first")

	msg := "second"
	got := Wrapf(err, msg)
	if !strings.Contains(got.Error(), err.Error()) {
		t.Errorf("wrapped error did not contain the first error, got: %v", got)
	}
	if !strings.Contains(got.Error(), msg) {
		t.Errorf("wrapped error did not contain message, got: %v", got)
	}
	if cause := Cause(got); cause != err {
		t.Errorf("wrong cause: expected %v, got %v", err, cause)
	}
}

func TestWrapNilError(t *testing.T) {
	var err error
	got := Wrapf(err, "test")
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestMergeError(t *testing.T) {
	err1 := errors.New("first")
	err2 := errors.New("second")

	got := Merge(err1, err2)
	if !strings.Contains(got.Error(), err1.Error()) {
		t.Errorf("merged error did not contain first error, want: %v, got: %v", err1.Error(), got.Error())
	}
	if !strings.Contains(got.Error(), err2.Error()) {
		t.Errorf("merged error did not contain second error, want: %v, got: %v", err2.Error(), got.Error())
	}
}

func TestMergeNoError(t *testing.T) {
	got := Merge(nil, nil)
	if got != nil {
		t.Errorf("expected no error, got %v", got)
	}
}
