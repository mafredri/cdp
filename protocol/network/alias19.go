// +build go1.9

package network

import (
	"github.com/mafredri/cdp/protocol/internal"
)

// ResourceType is an alias for page.ResourceType to avoid a circular dependency.
//
// Deprecated: Use page.ResourceType instead.
type ResourceType = internal.PageResourceType

// FrameID is an alias for page.FrameID to avoid a circular dependency.
//
// Deprecated: Use page.FrameID instead.
type FrameID = internal.PageFrameID
