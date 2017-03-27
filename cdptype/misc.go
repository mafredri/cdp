package cdptype

import (
	"encoding/json"
	"strconv"
)

// UnmarshalJSON decodes the PageFrameID from either string or float.
func (p *PageFrameID) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		// Check for incorrect float value in response
		// (looking at you, edge-diagnostics-adapter).
		var f float64
		err2 := json.Unmarshal(data, &f)
		if err2 != nil {
			return err
		}
		s = strconv.FormatFloat(f, 'f', -1, 64)
	}
	*p = PageFrameID(s)
	return nil
}
