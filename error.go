package cdp

import "fmt"

// Error represents.
type Error struct {
	Domain string
	Op     string
	Err    error
}

func (e Error) Error() string {
	return fmt.Sprintf("cdp.%s: %s: %s", e.Domain, e.Op, e.Err.Error())
}

var (
	_ error = (*Error)(nil)
)
