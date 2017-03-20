package cdptype

import (
	"encoding/json"
	"fmt"
)

// Map returns the headers decoded into a map.
func (n NetworkHeaders) Map() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(n, &m)
	return m, err
}

// MustMap panics if the headers cannot be decoded into a map.
func (n NetworkHeaders) MustMap() map[string]string {
	m := make(map[string]string)
	if err := json.Unmarshal(n, &m); err != nil {
		panic(err)
	}
	return m
}

// Error implements error for RuntimeExceptionDetails.
func (r RuntimeExceptionDetails) Error() string {
	var desc string
	if r.Exception.Description != nil {
		desc = ": " + *r.Exception.Description
	}
	return fmt.Sprintf("cdptype.RuntimeExceptionDetails: %s exception at %d:%d%s", r.Text, r.LineNumber, r.ColumnNumber, desc)
}

var (
	_ error = (*RuntimeExceptionDetails)(nil)
)
