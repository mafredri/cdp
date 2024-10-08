// Code generated by cdpgen. DO NOT EDIT.

package headlessexperimental

// ScreenshotParams Encoding options for a screenshot.
type ScreenshotParams struct {
	// Format Image compression format (defaults to png).
	//
	// Values: "jpeg", "png", "webp".
	Format           *string `json:"format,omitempty"`
	Quality          *int    `json:"quality,omitempty"`          // Compression quality from range [0..100] (jpeg and webp only).
	OptimizeForSpeed *bool   `json:"optimizeForSpeed,omitempty"` // Optimize image encoding for speed, not for resulting size (defaults to false)
}
