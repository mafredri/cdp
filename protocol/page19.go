// +build go1.9

package protocol

import "github.com/mafredri/cdp/protocol/page"

// PageResourceType is provided for backwards compatibility with Go 1.8.
//
// Deprecated: Use page.ResourceType instead.
type PageResourceType = page.ResourceType

// PageFrameID is provided for backwards compatibility with Go 1.8.
//
// Deprecated: Use page.FrameID instead.
type PageFrameID = page.FrameID
