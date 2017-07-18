// +build go1.9

package internal

import (
	"encoding/json"
	"testing"
)

func TestFrameID_UnmarshalJSON(t *testing.T) {
	var id PageFrameID

	// Unmarshals string.
	err := json.Unmarshal([]byte(`"1000.1"`), &id)
	if err != nil {
		t.Error(err)
	}
	if id != "1000.1" {
		t.Errorf("Unmarshal string got %q; want %q", id, "1000.1")
	}

	// Also unmarshals float.
	err = json.Unmarshal([]byte(`2000.1`), &id)
	if err != nil {
		t.Error(err)
	}
	if id != "2000.1" {
		t.Errorf("Unmarshal float got %q; want %q", id, "2000.1")
	}
}
