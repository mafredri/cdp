package cdp

import (
	"errors"
	"testing"

	"github.com/mafredri/cdp/cdpdom"
)

func TestErrorCauser(t *testing.T) {
	err1 := errors.New("trigger")

	tests := []struct {
		name string
		err  error
		want error
	}{
		{"Returns underlying error", &cdpdom.OpError{Err: err1}, err1},
		{"Returns original error", err1, err1},
		{"Returns nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrorCauser(tt.err)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
