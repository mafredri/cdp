package cdpdom

import (
	"errors"
	"strings"
	"testing"
)

func TestOpError_ErrorContainsErrorCauser(t *testing.T) {
	causer := errors.New("error causer")
	err := &OpError{
		Domain: "Test",
		Op:     "Method",
		Err:    causer,
	}

	got := err.Error()
	if !strings.Contains(got, causer.Error()) {
		t.Errorf("Error() should contain error causer, got: %s", got)
	}
}
