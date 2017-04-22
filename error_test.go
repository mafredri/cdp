package cdp

import (
	"errors"
	"testing"
)

type opError struct {
	err error
}

func (o *opError) Cause() error {
	return o.err
}

func (o opError) Error() string {
	return o.err.Error()
}

func TestErrorCause(t *testing.T) {
	err1 := errors.New("trigger")

	tests := []struct {
		name string
		err  error
		want error
	}{
		{"Returns underlying error", &opError{err: err1}, err1},
		{"Returns original error", err1, err1},
		{"Returns nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrorCause(tt.err)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
