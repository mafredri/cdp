// Code generated by cdpgen. DO NOT EDIT.

package storage

import (
	"github.com/mafredri/cdp/protocol/browser"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
)

// GetStorageKeyForFrameArgs represents the arguments for GetStorageKeyForFrame in the Storage domain.
type GetStorageKeyForFrameArgs struct {
	FrameID page.FrameID `json:"frameId"` // No description.
}

// NewGetStorageKeyForFrameArgs initializes GetStorageKeyForFrameArgs with the required arguments.
func NewGetStorageKeyForFrameArgs(frameID page.FrameID) *GetStorageKeyForFrameArgs {
	args := new(GetStorageKeyForFrameArgs)
	args.FrameID = frameID
	return args
}

// GetStorageKeyForFrameReply represents the return values for GetStorageKeyForFrame in the Storage domain.
type GetStorageKeyForFrameReply struct {
	StorageKey SerializedStorageKey `json:"storageKey"` // No description.
}

// ClearDataForOriginArgs represents the arguments for ClearDataForOrigin in the Storage domain.
type ClearDataForOriginArgs struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated list of StorageType to clear.
}

// NewClearDataForOriginArgs initializes ClearDataForOriginArgs with the required arguments.
func NewClearDataForOriginArgs(origin string, storageTypes string) *ClearDataForOriginArgs {
	args := new(ClearDataForOriginArgs)
	args.Origin = origin
	args.StorageTypes = storageTypes
	return args
}

// GetCookiesArgs represents the arguments for GetCookies in the Storage domain.
type GetCookiesArgs struct {
	BrowserContextID *browser.ContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// NewGetCookiesArgs initializes GetCookiesArgs with the required arguments.
func NewGetCookiesArgs() *GetCookiesArgs {
	args := new(GetCookiesArgs)

	return args
}

// SetBrowserContextID sets the BrowserContextID optional argument.
// Browser context to use when called on the browser endpoint.
func (a *GetCookiesArgs) SetBrowserContextID(browserContextID browser.ContextID) *GetCookiesArgs {
	a.BrowserContextID = &browserContextID
	return a
}

// GetCookiesReply represents the return values for GetCookies in the Storage domain.
type GetCookiesReply struct {
	Cookies []network.Cookie `json:"cookies"` // Array of cookie objects.
}

// SetCookiesArgs represents the arguments for SetCookies in the Storage domain.
type SetCookiesArgs struct {
	Cookies          []network.CookieParam `json:"cookies"`                    // Cookies to be set.
	BrowserContextID *browser.ContextID    `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// NewSetCookiesArgs initializes SetCookiesArgs with the required arguments.
func NewSetCookiesArgs(cookies []network.CookieParam) *SetCookiesArgs {
	args := new(SetCookiesArgs)
	args.Cookies = cookies
	return args
}

// SetBrowserContextID sets the BrowserContextID optional argument.
// Browser context to use when called on the browser endpoint.
func (a *SetCookiesArgs) SetBrowserContextID(browserContextID browser.ContextID) *SetCookiesArgs {
	a.BrowserContextID = &browserContextID
	return a
}

// ClearCookiesArgs represents the arguments for ClearCookies in the Storage domain.
type ClearCookiesArgs struct {
	BrowserContextID *browser.ContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// NewClearCookiesArgs initializes ClearCookiesArgs with the required arguments.
func NewClearCookiesArgs() *ClearCookiesArgs {
	args := new(ClearCookiesArgs)

	return args
}

// SetBrowserContextID sets the BrowserContextID optional argument.
// Browser context to use when called on the browser endpoint.
func (a *ClearCookiesArgs) SetBrowserContextID(browserContextID browser.ContextID) *ClearCookiesArgs {
	a.BrowserContextID = &browserContextID
	return a
}

// GetUsageAndQuotaArgs represents the arguments for GetUsageAndQuota in the Storage domain.
type GetUsageAndQuotaArgs struct {
	Origin string `json:"origin"` // Security origin.
}

// NewGetUsageAndQuotaArgs initializes GetUsageAndQuotaArgs with the required arguments.
func NewGetUsageAndQuotaArgs(origin string) *GetUsageAndQuotaArgs {
	args := new(GetUsageAndQuotaArgs)
	args.Origin = origin
	return args
}

// GetUsageAndQuotaReply represents the return values for GetUsageAndQuota in the Storage domain.
type GetUsageAndQuotaReply struct {
	Usage          float64        `json:"usage"`          // Storage usage (bytes).
	Quota          float64        `json:"quota"`          // Storage quota (bytes).
	OverrideActive bool           `json:"overrideActive"` // Whether or not the origin has an active storage quota override
	UsageBreakdown []UsageForType `json:"usageBreakdown"` // Storage usage per type (bytes).
}

// OverrideQuotaForOriginArgs represents the arguments for OverrideQuotaForOrigin in the Storage domain.
type OverrideQuotaForOriginArgs struct {
	Origin    string   `json:"origin"`              // Security origin.
	QuotaSize *float64 `json:"quotaSize,omitempty"` // The quota size (in bytes) to override the original quota with. If this is called multiple times, the overridden quota will be equal to the quotaSize provided in the final call. If this is called without specifying a quotaSize, the quota will be reset to the default value for the specified origin. If this is called multiple times with different origins, the override will be maintained for each origin until it is disabled (called without a quotaSize).
}

// NewOverrideQuotaForOriginArgs initializes OverrideQuotaForOriginArgs with the required arguments.
func NewOverrideQuotaForOriginArgs(origin string) *OverrideQuotaForOriginArgs {
	args := new(OverrideQuotaForOriginArgs)
	args.Origin = origin
	return args
}

// SetQuotaSize sets the QuotaSize optional argument. The quota size
// (in bytes) to override the original quota with. If this is called
// multiple times, the overridden quota will be equal to the quotaSize
// provided in the final call. If this is called without specifying a
// quotaSize, the quota will be reset to the default value for the
// specified origin. If this is called multiple times with different
// origins, the override will be maintained for each origin until it is
// disabled (called without a quotaSize).
func (a *OverrideQuotaForOriginArgs) SetQuotaSize(quotaSize float64) *OverrideQuotaForOriginArgs {
	a.QuotaSize = &quotaSize
	return a
}

// TrackCacheStorageForOriginArgs represents the arguments for TrackCacheStorageForOrigin in the Storage domain.
type TrackCacheStorageForOriginArgs struct {
	Origin string `json:"origin"` // Security origin.
}

// NewTrackCacheStorageForOriginArgs initializes TrackCacheStorageForOriginArgs with the required arguments.
func NewTrackCacheStorageForOriginArgs(origin string) *TrackCacheStorageForOriginArgs {
	args := new(TrackCacheStorageForOriginArgs)
	args.Origin = origin
	return args
}

// TrackIndexedDBForOriginArgs represents the arguments for TrackIndexedDBForOrigin in the Storage domain.
type TrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"` // Security origin.
}

// NewTrackIndexedDBForOriginArgs initializes TrackIndexedDBForOriginArgs with the required arguments.
func NewTrackIndexedDBForOriginArgs(origin string) *TrackIndexedDBForOriginArgs {
	args := new(TrackIndexedDBForOriginArgs)
	args.Origin = origin
	return args
}

// UntrackCacheStorageForOriginArgs represents the arguments for UntrackCacheStorageForOrigin in the Storage domain.
type UntrackCacheStorageForOriginArgs struct {
	Origin string `json:"origin"` // Security origin.
}

// NewUntrackCacheStorageForOriginArgs initializes UntrackCacheStorageForOriginArgs with the required arguments.
func NewUntrackCacheStorageForOriginArgs(origin string) *UntrackCacheStorageForOriginArgs {
	args := new(UntrackCacheStorageForOriginArgs)
	args.Origin = origin
	return args
}

// UntrackIndexedDBForOriginArgs represents the arguments for UntrackIndexedDBForOrigin in the Storage domain.
type UntrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"` // Security origin.
}

// NewUntrackIndexedDBForOriginArgs initializes UntrackIndexedDBForOriginArgs with the required arguments.
func NewUntrackIndexedDBForOriginArgs(origin string) *UntrackIndexedDBForOriginArgs {
	args := new(UntrackIndexedDBForOriginArgs)
	args.Origin = origin
	return args
}

// GetTrustTokensReply represents the return values for GetTrustTokens in the Storage domain.
type GetTrustTokensReply struct {
	Tokens []TrustTokens `json:"tokens"` // No description.
}

// ClearTrustTokensArgs represents the arguments for ClearTrustTokens in the Storage domain.
type ClearTrustTokensArgs struct {
	IssuerOrigin string `json:"issuerOrigin"` // No description.
}

// NewClearTrustTokensArgs initializes ClearTrustTokensArgs with the required arguments.
func NewClearTrustTokensArgs(issuerOrigin string) *ClearTrustTokensArgs {
	args := new(ClearTrustTokensArgs)
	args.IssuerOrigin = issuerOrigin
	return args
}

// ClearTrustTokensReply represents the return values for ClearTrustTokens in the Storage domain.
type ClearTrustTokensReply struct {
	DidDeleteTokens bool `json:"didDeleteTokens"` // True if any tokens were deleted, false otherwise.
}

// GetInterestGroupDetailsArgs represents the arguments for GetInterestGroupDetails in the Storage domain.
type GetInterestGroupDetailsArgs struct {
	OwnerOrigin string `json:"ownerOrigin"` // No description.
	Name        string `json:"name"`        // No description.
}

// NewGetInterestGroupDetailsArgs initializes GetInterestGroupDetailsArgs with the required arguments.
func NewGetInterestGroupDetailsArgs(ownerOrigin string, name string) *GetInterestGroupDetailsArgs {
	args := new(GetInterestGroupDetailsArgs)
	args.OwnerOrigin = ownerOrigin
	args.Name = name
	return args
}

// GetInterestGroupDetailsReply represents the return values for GetInterestGroupDetails in the Storage domain.
type GetInterestGroupDetailsReply struct {
	Details InterestGroupDetails `json:"details"` // No description.
}

// SetInterestGroupTrackingArgs represents the arguments for SetInterestGroupTracking in the Storage domain.
type SetInterestGroupTrackingArgs struct {
	Enable bool `json:"enable"` // No description.
}

// NewSetInterestGroupTrackingArgs initializes SetInterestGroupTrackingArgs with the required arguments.
func NewSetInterestGroupTrackingArgs(enable bool) *SetInterestGroupTrackingArgs {
	args := new(SetInterestGroupTrackingArgs)
	args.Enable = enable
	return args
}
