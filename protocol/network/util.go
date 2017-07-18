package network

import (
	"encoding/json"
)

// Map returns the headers decoded into a map.
func (n Headers) Map() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(n, &m)
	return m, err
}
