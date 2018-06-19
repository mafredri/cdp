// Code generated by cdpgen. DO NOT EDIT.

package emulation

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
)

// CanEmulateReply represents the return values for CanEmulate in the Emulation domain.
type CanEmulateReply struct {
	Result bool `json:"result"` // True if emulation is supported.
}

// SetCPUThrottlingRateArgs represents the arguments for SetCPUThrottlingRate in the Emulation domain.
type SetCPUThrottlingRateArgs struct {
	Rate float64 `json:"rate"` // Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
}

// NewSetCPUThrottlingRateArgs initializes SetCPUThrottlingRateArgs with the required arguments.
func NewSetCPUThrottlingRateArgs(rate float64) *SetCPUThrottlingRateArgs {
	args := new(SetCPUThrottlingRateArgs)
	args.Rate = rate
	return args
}

// SetDefaultBackgroundColorOverrideArgs represents the arguments for SetDefaultBackgroundColorOverride in the Emulation domain.
type SetDefaultBackgroundColorOverrideArgs struct {
	Color *dom.RGBA `json:"color,omitempty"` // RGBA of the default background color. If not specified, any existing override will be cleared.
}

// NewSetDefaultBackgroundColorOverrideArgs initializes SetDefaultBackgroundColorOverrideArgs with the required arguments.
func NewSetDefaultBackgroundColorOverrideArgs() *SetDefaultBackgroundColorOverrideArgs {
	args := new(SetDefaultBackgroundColorOverrideArgs)

	return args
}

// SetColor sets the Color optional argument. RGBA of the default
// background color. If not specified, any existing override will be
// cleared.
func (a *SetDefaultBackgroundColorOverrideArgs) SetColor(color dom.RGBA) *SetDefaultBackgroundColorOverrideArgs {
	a.Color = &color
	return a
}

// SetDeviceMetricsOverrideArgs represents the arguments for SetDeviceMetricsOverride in the Emulation domain.
type SetDeviceMetricsOverrideArgs struct {
	Width             int     `json:"width"`             // Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height            int     `json:"height"`            // Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"` // Overriding device scale factor value. 0 disables the override.
	Mobile            bool    `json:"mobile"`            // Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	// Scale Scale to apply to resulting view image.
	//
	// Note: This property is experimental.
	Scale *float64 `json:"scale,omitempty"`
	// ScreenWidth Overriding screen width value in pixels (minimum 0,
	// maximum 10000000).
	//
	// Note: This property is experimental.
	ScreenWidth *int `json:"screenWidth,omitempty"`
	// ScreenHeight Overriding screen height value in pixels (minimum 0,
	// maximum 10000000).
	//
	// Note: This property is experimental.
	ScreenHeight *int `json:"screenHeight,omitempty"`
	// PositionX Overriding view X position on screen in pixels (minimum
	// 0, maximum 10000000).
	//
	// Note: This property is experimental.
	PositionX *int `json:"positionX,omitempty"`
	// PositionY Overriding view Y position on screen in pixels (minimum
	// 0, maximum 10000000).
	//
	// Note: This property is experimental.
	PositionY *int `json:"positionY,omitempty"`
	// DontSetVisibleSize Do not set visible view size, rely upon explicit
	// setVisibleSize call.
	//
	// Note: This property is experimental.
	DontSetVisibleSize *bool              `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation  *ScreenOrientation `json:"screenOrientation,omitempty"` // Screen orientation override.
	// Viewport If set, the visible area of the page will be overridden to
	// this viewport. This viewport change is not observed by the page,
	// e.g. viewport-relative elements do not change positions.
	//
	// Note: This property is experimental.
	Viewport *page.Viewport `json:"viewport,omitempty"`
}

// NewSetDeviceMetricsOverrideArgs initializes SetDeviceMetricsOverrideArgs with the required arguments.
func NewSetDeviceMetricsOverrideArgs(width int, height int, deviceScaleFactor float64, mobile bool) *SetDeviceMetricsOverrideArgs {
	args := new(SetDeviceMetricsOverrideArgs)
	args.Width = width
	args.Height = height
	args.DeviceScaleFactor = deviceScaleFactor
	args.Mobile = mobile
	return args
}

// SetScale sets the Scale optional argument. Scale to apply to
// resulting view image.
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetScale(scale float64) *SetDeviceMetricsOverrideArgs {
	a.Scale = &scale
	return a
}

// SetScreenWidth sets the ScreenWidth optional argument. Overriding
// screen width value in pixels (minimum 0, maximum 10000000).
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetScreenWidth(screenWidth int) *SetDeviceMetricsOverrideArgs {
	a.ScreenWidth = &screenWidth
	return a
}

// SetScreenHeight sets the ScreenHeight optional argument. Overriding
// screen height value in pixels (minimum 0, maximum 10000000).
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetScreenHeight(screenHeight int) *SetDeviceMetricsOverrideArgs {
	a.ScreenHeight = &screenHeight
	return a
}

// SetPositionX sets the PositionX optional argument. Overriding view
// X position on screen in pixels (minimum 0, maximum 10000000).
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetPositionX(positionX int) *SetDeviceMetricsOverrideArgs {
	a.PositionX = &positionX
	return a
}

// SetPositionY sets the PositionY optional argument. Overriding view
// Y position on screen in pixels (minimum 0, maximum 10000000).
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetPositionY(positionY int) *SetDeviceMetricsOverrideArgs {
	a.PositionY = &positionY
	return a
}

// SetDontSetVisibleSize sets the DontSetVisibleSize optional argument.
// Do not set visible view size, rely upon explicit setVisibleSize
// call.
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetDontSetVisibleSize(dontSetVisibleSize bool) *SetDeviceMetricsOverrideArgs {
	a.DontSetVisibleSize = &dontSetVisibleSize
	return a
}

// SetScreenOrientation sets the ScreenOrientation optional argument.
// Screen orientation override.
func (a *SetDeviceMetricsOverrideArgs) SetScreenOrientation(screenOrientation ScreenOrientation) *SetDeviceMetricsOverrideArgs {
	a.ScreenOrientation = &screenOrientation
	return a
}

// SetViewport sets the Viewport optional argument. If set, the
// visible area of the page will be overridden to this viewport. This
// viewport change is not observed by the page, e.g. viewport-relative
// elements do not change positions.
//
// Note: This property is experimental.
func (a *SetDeviceMetricsOverrideArgs) SetViewport(viewport page.Viewport) *SetDeviceMetricsOverrideArgs {
	a.Viewport = &viewport
	return a
}

// SetScrollbarsHiddenArgs represents the arguments for SetScrollbarsHidden in the Emulation domain.
type SetScrollbarsHiddenArgs struct {
	Hidden bool `json:"hidden"` // Whether scrollbars should be always hidden.
}

// NewSetScrollbarsHiddenArgs initializes SetScrollbarsHiddenArgs with the required arguments.
func NewSetScrollbarsHiddenArgs(hidden bool) *SetScrollbarsHiddenArgs {
	args := new(SetScrollbarsHiddenArgs)
	args.Hidden = hidden
	return args
}

// SetDocumentCookieDisabledArgs represents the arguments for SetDocumentCookieDisabled in the Emulation domain.
type SetDocumentCookieDisabledArgs struct {
	Disabled bool `json:"disabled"` // Whether document.coookie API should be disabled.
}

// NewSetDocumentCookieDisabledArgs initializes SetDocumentCookieDisabledArgs with the required arguments.
func NewSetDocumentCookieDisabledArgs(disabled bool) *SetDocumentCookieDisabledArgs {
	args := new(SetDocumentCookieDisabledArgs)
	args.Disabled = disabled
	return args
}

// SetEmitTouchEventsForMouseArgs represents the arguments for SetEmitTouchEventsForMouse in the Emulation domain.
type SetEmitTouchEventsForMouseArgs struct {
	Enabled bool `json:"enabled"` // Whether touch emulation based on mouse input should be enabled.
	// Configuration Touch/gesture events configuration. Default: current
	// platform.
	//
	// Values: "mobile", "desktop".
	Configuration *string `json:"configuration,omitempty"`
}

// NewSetEmitTouchEventsForMouseArgs initializes SetEmitTouchEventsForMouseArgs with the required arguments.
func NewSetEmitTouchEventsForMouseArgs(enabled bool) *SetEmitTouchEventsForMouseArgs {
	args := new(SetEmitTouchEventsForMouseArgs)
	args.Enabled = enabled
	return args
}

// SetConfiguration sets the Configuration optional argument.
// Touch/gesture events configuration. Default: current platform.
//
// Values: "mobile", "desktop".
func (a *SetEmitTouchEventsForMouseArgs) SetConfiguration(configuration string) *SetEmitTouchEventsForMouseArgs {
	a.Configuration = &configuration
	return a
}

// SetEmulatedMediaArgs represents the arguments for SetEmulatedMedia in the Emulation domain.
type SetEmulatedMediaArgs struct {
	Media string `json:"media"` // Media type to emulate. Empty string disables the override.
}

// NewSetEmulatedMediaArgs initializes SetEmulatedMediaArgs with the required arguments.
func NewSetEmulatedMediaArgs(media string) *SetEmulatedMediaArgs {
	args := new(SetEmulatedMediaArgs)
	args.Media = media
	return args
}

// SetGeolocationOverrideArgs represents the arguments for SetGeolocationOverride in the Emulation domain.
type SetGeolocationOverrideArgs struct {
	Latitude  *float64 `json:"latitude,omitempty"`  // Mock latitude
	Longitude *float64 `json:"longitude,omitempty"` // Mock longitude
	Accuracy  *float64 `json:"accuracy,omitempty"`  // Mock accuracy
}

// NewSetGeolocationOverrideArgs initializes SetGeolocationOverrideArgs with the required arguments.
func NewSetGeolocationOverrideArgs() *SetGeolocationOverrideArgs {
	args := new(SetGeolocationOverrideArgs)

	return args
}

// SetLatitude sets the Latitude optional argument. Mock latitude
func (a *SetGeolocationOverrideArgs) SetLatitude(latitude float64) *SetGeolocationOverrideArgs {
	a.Latitude = &latitude
	return a
}

// SetLongitude sets the Longitude optional argument. Mock longitude
func (a *SetGeolocationOverrideArgs) SetLongitude(longitude float64) *SetGeolocationOverrideArgs {
	a.Longitude = &longitude
	return a
}

// SetAccuracy sets the Accuracy optional argument. Mock accuracy
func (a *SetGeolocationOverrideArgs) SetAccuracy(accuracy float64) *SetGeolocationOverrideArgs {
	a.Accuracy = &accuracy
	return a
}

// SetNavigatorOverridesArgs represents the arguments for SetNavigatorOverrides in the Emulation domain.
type SetNavigatorOverridesArgs struct {
	Platform string `json:"platform"` // The platform navigator.platform should return.
}

// NewSetNavigatorOverridesArgs initializes SetNavigatorOverridesArgs with the required arguments.
func NewSetNavigatorOverridesArgs(platform string) *SetNavigatorOverridesArgs {
	args := new(SetNavigatorOverridesArgs)
	args.Platform = platform
	return args
}

// SetPageScaleFactorArgs represents the arguments for SetPageScaleFactor in the Emulation domain.
type SetPageScaleFactorArgs struct {
	PageScaleFactor float64 `json:"pageScaleFactor"` // Page scale factor.
}

// NewSetPageScaleFactorArgs initializes SetPageScaleFactorArgs with the required arguments.
func NewSetPageScaleFactorArgs(pageScaleFactor float64) *SetPageScaleFactorArgs {
	args := new(SetPageScaleFactorArgs)
	args.PageScaleFactor = pageScaleFactor
	return args
}

// SetScriptExecutionDisabledArgs represents the arguments for SetScriptExecutionDisabled in the Emulation domain.
type SetScriptExecutionDisabledArgs struct {
	Value bool `json:"value"` // Whether script execution should be disabled in the page.
}

// NewSetScriptExecutionDisabledArgs initializes SetScriptExecutionDisabledArgs with the required arguments.
func NewSetScriptExecutionDisabledArgs(value bool) *SetScriptExecutionDisabledArgs {
	args := new(SetScriptExecutionDisabledArgs)
	args.Value = value
	return args
}

// SetTouchEmulationEnabledArgs represents the arguments for SetTouchEmulationEnabled in the Emulation domain.
type SetTouchEmulationEnabledArgs struct {
	Enabled        bool `json:"enabled"`                  // Whether the touch event emulation should be enabled.
	MaxTouchPoints *int `json:"maxTouchPoints,omitempty"` // Maximum touch points supported. Defaults to one.
}

// NewSetTouchEmulationEnabledArgs initializes SetTouchEmulationEnabledArgs with the required arguments.
func NewSetTouchEmulationEnabledArgs(enabled bool) *SetTouchEmulationEnabledArgs {
	args := new(SetTouchEmulationEnabledArgs)
	args.Enabled = enabled
	return args
}

// SetMaxTouchPoints sets the MaxTouchPoints optional argument.
// Maximum touch points supported. Defaults to one.
func (a *SetTouchEmulationEnabledArgs) SetMaxTouchPoints(maxTouchPoints int) *SetTouchEmulationEnabledArgs {
	a.MaxTouchPoints = &maxTouchPoints
	return a
}

// SetVirtualTimePolicyArgs represents the arguments for SetVirtualTimePolicy in the Emulation domain.
type SetVirtualTimePolicyArgs struct {
	Policy                            VirtualTimePolicy      `json:"policy"`                                      // No description.
	Budget                            *float64               `json:"budget,omitempty"`                            // If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	MaxVirtualTimeTaskStarvationCount *int                   `json:"maxVirtualTimeTaskStarvationCount,omitempty"` // If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
	WaitForNavigation                 *bool                  `json:"waitForNavigation,omitempty"`                 // If set the virtual time policy change should be deferred until any frame starts navigating. Note any previous deferred policy change is superseded.
	InitialVirtualTime                network.TimeSinceEpoch `json:"initialVirtualTime,omitempty"`                // If set, base::Time::Now will be overridden to initially return this value.
}

// NewSetVirtualTimePolicyArgs initializes SetVirtualTimePolicyArgs with the required arguments.
func NewSetVirtualTimePolicyArgs(policy VirtualTimePolicy) *SetVirtualTimePolicyArgs {
	args := new(SetVirtualTimePolicyArgs)
	args.Policy = policy
	return args
}

// SetBudget sets the Budget optional argument. If set, after this
// many virtual milliseconds have elapsed virtual time will be paused
// and a virtualTimeBudgetExpired event is sent.
func (a *SetVirtualTimePolicyArgs) SetBudget(budget float64) *SetVirtualTimePolicyArgs {
	a.Budget = &budget
	return a
}

// SetMaxVirtualTimeTaskStarvationCount sets the MaxVirtualTimeTaskStarvationCount optional argument.
// If set this specifies the maximum number of tasks that can be run
// before virtual is forced forwards to prevent deadlock.
func (a *SetVirtualTimePolicyArgs) SetMaxVirtualTimeTaskStarvationCount(maxVirtualTimeTaskStarvationCount int) *SetVirtualTimePolicyArgs {
	a.MaxVirtualTimeTaskStarvationCount = &maxVirtualTimeTaskStarvationCount
	return a
}

// SetWaitForNavigation sets the WaitForNavigation optional argument.
// If set the virtual time policy change should be deferred until any
// frame starts navigating. Note any previous deferred policy change is
// superseded.
func (a *SetVirtualTimePolicyArgs) SetWaitForNavigation(waitForNavigation bool) *SetVirtualTimePolicyArgs {
	a.WaitForNavigation = &waitForNavigation
	return a
}

// SetInitialVirtualTime sets the InitialVirtualTime optional argument.
// If set, base::Time::Now will be overridden to initially return this
// value.
func (a *SetVirtualTimePolicyArgs) SetInitialVirtualTime(initialVirtualTime network.TimeSinceEpoch) *SetVirtualTimePolicyArgs {
	a.InitialVirtualTime = initialVirtualTime
	return a
}

// SetVirtualTimePolicyReply represents the return values for SetVirtualTimePolicy in the Emulation domain.
type SetVirtualTimePolicyReply struct {
	VirtualTimeTicksBase float64 `json:"virtualTimeTicksBase"` // Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
}

// SetVisibleSizeArgs represents the arguments for SetVisibleSize in the Emulation domain.
type SetVisibleSizeArgs struct {
	Width  int `json:"width"`  // Frame width (DIP).
	Height int `json:"height"` // Frame height (DIP).
}

// NewSetVisibleSizeArgs initializes SetVisibleSizeArgs with the required arguments.
func NewSetVisibleSizeArgs(width int, height int) *SetVisibleSizeArgs {
	args := new(SetVisibleSizeArgs)
	args.Width = width
	args.Height = height
	return args
}

// SetUserAgentOverrideArgs represents the arguments for SetUserAgentOverride in the Emulation domain.
type SetUserAgentOverrideArgs struct {
	UserAgent      string  `json:"userAgent"`                // User agent to use.
	AcceptLanguage *string `json:"acceptLanguage,omitempty"` // Browser langugage to emulate.
	Platform       *string `json:"platform,omitempty"`       // The platform navigator.platform should return.
}

// NewSetUserAgentOverrideArgs initializes SetUserAgentOverrideArgs with the required arguments.
func NewSetUserAgentOverrideArgs(userAgent string) *SetUserAgentOverrideArgs {
	args := new(SetUserAgentOverrideArgs)
	args.UserAgent = userAgent
	return args
}

// SetAcceptLanguage sets the AcceptLanguage optional argument.
// Browser langugage to emulate.
func (a *SetUserAgentOverrideArgs) SetAcceptLanguage(acceptLanguage string) *SetUserAgentOverrideArgs {
	a.AcceptLanguage = &acceptLanguage
	return a
}

// SetPlatform sets the Platform optional argument. The platform
// navigator.platform should return.
func (a *SetUserAgentOverrideArgs) SetPlatform(platform string) *SetUserAgentOverrideArgs {
	a.Platform = &platform
	return a
}
