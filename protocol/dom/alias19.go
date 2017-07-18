// +build go1.9

package dom

import (
	"github.com/mafredri/cdp/protocol/internal"
)

// FrameID is an alias for page.FrameID to avoid a circular dependency.
//
// Deprecated: Use page.FrameID instead.
type FrameID = internal.PageFrameID
