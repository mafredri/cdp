package cdpcmd

import (
	"testing"

	"github.com/mafredri/cdp/cdptype"
)

func TestNewAccessibilityGetPartialAXTreeArgs(t *testing.T) {
	args := NewAccessibilityGetPartialAXTreeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewAccessibilityGetPartialAXTreeArgs returned nil args")
	}
}

func TestNewAnimationSetPlaybackRateArgs(t *testing.T) {
	args := NewAnimationSetPlaybackRateArgs(func() (playbackRate float64) { return }())
	if args == nil {
		t.Errorf("NewAnimationSetPlaybackRateArgs returned nil args")
	}
}

func TestNewAnimationGetCurrentTimeArgs(t *testing.T) {
	args := NewAnimationGetCurrentTimeArgs(func() (id string) { return }())
	if args == nil {
		t.Errorf("NewAnimationGetCurrentTimeArgs returned nil args")
	}
}

func TestNewAnimationSetPausedArgs(t *testing.T) {
	args := NewAnimationSetPausedArgs(func() (animations []string, paused bool) { return }())
	if args == nil {
		t.Errorf("NewAnimationSetPausedArgs returned nil args")
	}
}

func TestNewAnimationSetTimingArgs(t *testing.T) {
	args := NewAnimationSetTimingArgs(func() (animationID string, duration float64, delay float64) { return }())
	if args == nil {
		t.Errorf("NewAnimationSetTimingArgs returned nil args")
	}
}

func TestNewAnimationSeekAnimationsArgs(t *testing.T) {
	args := NewAnimationSeekAnimationsArgs(func() (animations []string, currentTime float64) { return }())
	if args == nil {
		t.Errorf("NewAnimationSeekAnimationsArgs returned nil args")
	}
}

func TestNewAnimationReleaseAnimationsArgs(t *testing.T) {
	args := NewAnimationReleaseAnimationsArgs(func() (animations []string) { return }())
	if args == nil {
		t.Errorf("NewAnimationReleaseAnimationsArgs returned nil args")
	}
}

func TestNewAnimationResolveAnimationArgs(t *testing.T) {
	args := NewAnimationResolveAnimationArgs(func() (animationID string) { return }())
	if args == nil {
		t.Errorf("NewAnimationResolveAnimationArgs returned nil args")
	}
}

func TestNewApplicationCacheGetManifestForFrameArgs(t *testing.T) {
	args := NewApplicationCacheGetManifestForFrameArgs(func() (frameID cdptype.PageFrameID) { return }())
	if args == nil {
		t.Errorf("NewApplicationCacheGetManifestForFrameArgs returned nil args")
	}
}

func TestNewApplicationCacheGetApplicationCacheForFrameArgs(t *testing.T) {
	args := NewApplicationCacheGetApplicationCacheForFrameArgs(func() (frameID cdptype.PageFrameID) { return }())
	if args == nil {
		t.Errorf("NewApplicationCacheGetApplicationCacheForFrameArgs returned nil args")
	}
}

func TestNewBrowserGetWindowForTargetArgs(t *testing.T) {
	args := NewBrowserGetWindowForTargetArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewBrowserGetWindowForTargetArgs returned nil args")
	}
}

func TestNewBrowserSetWindowBoundsArgs(t *testing.T) {
	args := NewBrowserSetWindowBoundsArgs(func() (windowID cdptype.BrowserWindowID, bounds cdptype.BrowserBounds) { return }())
	if args == nil {
		t.Errorf("NewBrowserSetWindowBoundsArgs returned nil args")
	}
}

func TestNewBrowserGetWindowBoundsArgs(t *testing.T) {
	args := NewBrowserGetWindowBoundsArgs(func() (windowID cdptype.BrowserWindowID) { return }())
	if args == nil {
		t.Errorf("NewBrowserGetWindowBoundsArgs returned nil args")
	}
}

func TestNewCSSGetMatchedStylesForNodeArgs(t *testing.T) {
	args := NewCSSGetMatchedStylesForNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetMatchedStylesForNodeArgs returned nil args")
	}
}

func TestNewCSSGetInlineStylesForNodeArgs(t *testing.T) {
	args := NewCSSGetInlineStylesForNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetInlineStylesForNodeArgs returned nil args")
	}
}

func TestNewCSSGetComputedStyleForNodeArgs(t *testing.T) {
	args := NewCSSGetComputedStyleForNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetComputedStyleForNodeArgs returned nil args")
	}
}

func TestNewCSSGetPlatformFontsForNodeArgs(t *testing.T) {
	args := NewCSSGetPlatformFontsForNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetPlatformFontsForNodeArgs returned nil args")
	}
}

func TestNewCSSGetStyleSheetTextArgs(t *testing.T) {
	args := NewCSSGetStyleSheetTextArgs(func() (styleSheetID cdptype.CSSStyleSheetID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetStyleSheetTextArgs returned nil args")
	}
}

func TestNewCSSCollectClassNamesArgs(t *testing.T) {
	args := NewCSSCollectClassNamesArgs(func() (styleSheetID cdptype.CSSStyleSheetID) { return }())
	if args == nil {
		t.Errorf("NewCSSCollectClassNamesArgs returned nil args")
	}
}

func TestNewCSSSetStyleSheetTextArgs(t *testing.T) {
	args := NewCSSSetStyleSheetTextArgs(func() (styleSheetID cdptype.CSSStyleSheetID, text string) { return }())
	if args == nil {
		t.Errorf("NewCSSSetStyleSheetTextArgs returned nil args")
	}
}

func TestNewCSSSetRuleSelectorArgs(t *testing.T) {
	args := NewCSSSetRuleSelectorArgs(func() (styleSheetID cdptype.CSSStyleSheetID, rang cdptype.CSSSourceRange, selector string) { return }())
	if args == nil {
		t.Errorf("NewCSSSetRuleSelectorArgs returned nil args")
	}
}

func TestNewCSSSetKeyframeKeyArgs(t *testing.T) {
	args := NewCSSSetKeyframeKeyArgs(func() (styleSheetID cdptype.CSSStyleSheetID, rang cdptype.CSSSourceRange, keyText string) { return }())
	if args == nil {
		t.Errorf("NewCSSSetKeyframeKeyArgs returned nil args")
	}
}

func TestNewCSSSetStyleTextsArgs(t *testing.T) {
	args := NewCSSSetStyleTextsArgs(func() (edits []cdptype.CSSStyleDeclarationEdit) { return }())
	if args == nil {
		t.Errorf("NewCSSSetStyleTextsArgs returned nil args")
	}
}

func TestNewCSSSetMediaTextArgs(t *testing.T) {
	args := NewCSSSetMediaTextArgs(func() (styleSheetID cdptype.CSSStyleSheetID, rang cdptype.CSSSourceRange, text string) { return }())
	if args == nil {
		t.Errorf("NewCSSSetMediaTextArgs returned nil args")
	}
}

func TestNewCSSCreateStyleSheetArgs(t *testing.T) {
	args := NewCSSCreateStyleSheetArgs(func() (frameID cdptype.PageFrameID) { return }())
	if args == nil {
		t.Errorf("NewCSSCreateStyleSheetArgs returned nil args")
	}
}

func TestNewCSSAddRuleArgs(t *testing.T) {
	args := NewCSSAddRuleArgs(func() (styleSheetID cdptype.CSSStyleSheetID, ruleText string, location cdptype.CSSSourceRange) {
		return
	}())
	if args == nil {
		t.Errorf("NewCSSAddRuleArgs returned nil args")
	}
}

func TestNewCSSForcePseudoStateArgs(t *testing.T) {
	args := NewCSSForcePseudoStateArgs(func() (nodeID cdptype.DOMNodeID, forcedPseudoClasses []string) { return }())
	if args == nil {
		t.Errorf("NewCSSForcePseudoStateArgs returned nil args")
	}
}

func TestNewCSSSetEffectivePropertyValueForNodeArgs(t *testing.T) {
	args := NewCSSSetEffectivePropertyValueForNodeArgs(func() (nodeID cdptype.DOMNodeID, propertyName string, value string) { return }())
	if args == nil {
		t.Errorf("NewCSSSetEffectivePropertyValueForNodeArgs returned nil args")
	}
}

func TestNewCSSGetBackgroundColorsArgs(t *testing.T) {
	args := NewCSSGetBackgroundColorsArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewCSSGetBackgroundColorsArgs returned nil args")
	}
}

func TestNewCSSGetLayoutTreeAndStylesArgs(t *testing.T) {
	args := NewCSSGetLayoutTreeAndStylesArgs(func() (computedStyleWhitelist []string) { return }())
	if args == nil {
		t.Errorf("NewCSSGetLayoutTreeAndStylesArgs returned nil args")
	}
}

func TestNewCacheStorageRequestCacheNamesArgs(t *testing.T) {
	args := NewCacheStorageRequestCacheNamesArgs(func() (securityOrigin string) { return }())
	if args == nil {
		t.Errorf("NewCacheStorageRequestCacheNamesArgs returned nil args")
	}
}

func TestNewCacheStorageRequestEntriesArgs(t *testing.T) {
	args := NewCacheStorageRequestEntriesArgs(func() (cacheID cdptype.CacheStorageCacheID, skipCount int, pageSize int) { return }())
	if args == nil {
		t.Errorf("NewCacheStorageRequestEntriesArgs returned nil args")
	}
}

func TestNewCacheStorageDeleteCacheArgs(t *testing.T) {
	args := NewCacheStorageDeleteCacheArgs(func() (cacheID cdptype.CacheStorageCacheID) { return }())
	if args == nil {
		t.Errorf("NewCacheStorageDeleteCacheArgs returned nil args")
	}
}

func TestNewCacheStorageDeleteEntryArgs(t *testing.T) {
	args := NewCacheStorageDeleteEntryArgs(func() (cacheID cdptype.CacheStorageCacheID, request string) { return }())
	if args == nil {
		t.Errorf("NewCacheStorageDeleteEntryArgs returned nil args")
	}
}

func TestNewDOMGetDocumentArgs(t *testing.T) {
	args := NewDOMGetDocumentArgs()
	if args == nil {
		t.Errorf("NewDOMGetDocumentArgs returned nil args")
	}
}

func TestNewDOMGetFlattenedDocumentArgs(t *testing.T) {
	args := NewDOMGetFlattenedDocumentArgs()
	if args == nil {
		t.Errorf("NewDOMGetFlattenedDocumentArgs returned nil args")
	}
}

func TestNewDOMCollectClassNamesFromSubtreeArgs(t *testing.T) {
	args := NewDOMCollectClassNamesFromSubtreeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMCollectClassNamesFromSubtreeArgs returned nil args")
	}
}

func TestNewDOMRequestChildNodesArgs(t *testing.T) {
	args := NewDOMRequestChildNodesArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMRequestChildNodesArgs returned nil args")
	}
}

func TestNewDOMQuerySelectorArgs(t *testing.T) {
	args := NewDOMQuerySelectorArgs(func() (nodeID cdptype.DOMNodeID, selector string) { return }())
	if args == nil {
		t.Errorf("NewDOMQuerySelectorArgs returned nil args")
	}
}

func TestNewDOMQuerySelectorAllArgs(t *testing.T) {
	args := NewDOMQuerySelectorAllArgs(func() (nodeID cdptype.DOMNodeID, selector string) { return }())
	if args == nil {
		t.Errorf("NewDOMQuerySelectorAllArgs returned nil args")
	}
}

func TestNewDOMSetNodeNameArgs(t *testing.T) {
	args := NewDOMSetNodeNameArgs(func() (nodeID cdptype.DOMNodeID, name string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetNodeNameArgs returned nil args")
	}
}

func TestNewDOMSetNodeValueArgs(t *testing.T) {
	args := NewDOMSetNodeValueArgs(func() (nodeID cdptype.DOMNodeID, value string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetNodeValueArgs returned nil args")
	}
}

func TestNewDOMRemoveNodeArgs(t *testing.T) {
	args := NewDOMRemoveNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMRemoveNodeArgs returned nil args")
	}
}

func TestNewDOMSetAttributeValueArgs(t *testing.T) {
	args := NewDOMSetAttributeValueArgs(func() (nodeID cdptype.DOMNodeID, name string, value string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetAttributeValueArgs returned nil args")
	}
}

func TestNewDOMSetAttributesAsTextArgs(t *testing.T) {
	args := NewDOMSetAttributesAsTextArgs(func() (nodeID cdptype.DOMNodeID, text string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetAttributesAsTextArgs returned nil args")
	}
}

func TestNewDOMRemoveAttributeArgs(t *testing.T) {
	args := NewDOMRemoveAttributeArgs(func() (nodeID cdptype.DOMNodeID, name string) { return }())
	if args == nil {
		t.Errorf("NewDOMRemoveAttributeArgs returned nil args")
	}
}

func TestNewDOMGetOuterHTMLArgs(t *testing.T) {
	args := NewDOMGetOuterHTMLArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMGetOuterHTMLArgs returned nil args")
	}
}

func TestNewDOMSetOuterHTMLArgs(t *testing.T) {
	args := NewDOMSetOuterHTMLArgs(func() (nodeID cdptype.DOMNodeID, outerHTML string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetOuterHTMLArgs returned nil args")
	}
}

func TestNewDOMPerformSearchArgs(t *testing.T) {
	args := NewDOMPerformSearchArgs(func() (query string) { return }())
	if args == nil {
		t.Errorf("NewDOMPerformSearchArgs returned nil args")
	}
}

func TestNewDOMGetSearchResultsArgs(t *testing.T) {
	args := NewDOMGetSearchResultsArgs(func() (searchID string, fromIndex int, toIndex int) { return }())
	if args == nil {
		t.Errorf("NewDOMGetSearchResultsArgs returned nil args")
	}
}

func TestNewDOMDiscardSearchResultsArgs(t *testing.T) {
	args := NewDOMDiscardSearchResultsArgs(func() (searchID string) { return }())
	if args == nil {
		t.Errorf("NewDOMDiscardSearchResultsArgs returned nil args")
	}
}

func TestNewDOMRequestNodeArgs(t *testing.T) {
	args := NewDOMRequestNodeArgs(func() (objectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewDOMRequestNodeArgs returned nil args")
	}
}

func TestNewDOMPushNodeByPathToFrontendArgs(t *testing.T) {
	args := NewDOMPushNodeByPathToFrontendArgs(func() (path string) { return }())
	if args == nil {
		t.Errorf("NewDOMPushNodeByPathToFrontendArgs returned nil args")
	}
}

func TestNewDOMPushNodesByBackendIdsToFrontendArgs(t *testing.T) {
	args := NewDOMPushNodesByBackendIdsToFrontendArgs(func() (backendNodeIDs []cdptype.DOMBackendNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMPushNodesByBackendIdsToFrontendArgs returned nil args")
	}
}

func TestNewDOMSetInspectedNodeArgs(t *testing.T) {
	args := NewDOMSetInspectedNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMSetInspectedNodeArgs returned nil args")
	}
}

func TestNewDOMResolveNodeArgs(t *testing.T) {
	args := NewDOMResolveNodeArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMResolveNodeArgs returned nil args")
	}
}

func TestNewDOMGetAttributesArgs(t *testing.T) {
	args := NewDOMGetAttributesArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMGetAttributesArgs returned nil args")
	}
}

func TestNewDOMCopyToArgs(t *testing.T) {
	args := NewDOMCopyToArgs(func() (nodeID cdptype.DOMNodeID, targetNodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMCopyToArgs returned nil args")
	}
}

func TestNewDOMMoveToArgs(t *testing.T) {
	args := NewDOMMoveToArgs(func() (nodeID cdptype.DOMNodeID, targetNodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMMoveToArgs returned nil args")
	}
}

func TestNewDOMFocusArgs(t *testing.T) {
	args := NewDOMFocusArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMFocusArgs returned nil args")
	}
}

func TestNewDOMSetFileInputFilesArgs(t *testing.T) {
	args := NewDOMSetFileInputFilesArgs(func() (nodeID cdptype.DOMNodeID, files []string) { return }())
	if args == nil {
		t.Errorf("NewDOMSetFileInputFilesArgs returned nil args")
	}
}

func TestNewDOMGetBoxModelArgs(t *testing.T) {
	args := NewDOMGetBoxModelArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMGetBoxModelArgs returned nil args")
	}
}

func TestNewDOMGetNodeForLocationArgs(t *testing.T) {
	args := NewDOMGetNodeForLocationArgs(func() (x int, y int) { return }())
	if args == nil {
		t.Errorf("NewDOMGetNodeForLocationArgs returned nil args")
	}
}

func TestNewDOMGetRelayoutBoundaryArgs(t *testing.T) {
	args := NewDOMGetRelayoutBoundaryArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewDOMGetRelayoutBoundaryArgs returned nil args")
	}
}

func TestNewDOMDebuggerSetDOMBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerSetDOMBreakpointArgs(func() (nodeID cdptype.DOMNodeID, typ cdptype.DOMDebuggerDOMBreakpointType) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerSetDOMBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerRemoveDOMBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerRemoveDOMBreakpointArgs(func() (nodeID cdptype.DOMNodeID, typ cdptype.DOMDebuggerDOMBreakpointType) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerRemoveDOMBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerSetEventListenerBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerSetEventListenerBreakpointArgs(func() (eventName string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerSetEventListenerBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerRemoveEventListenerBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerRemoveEventListenerBreakpointArgs(func() (eventName string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerRemoveEventListenerBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerSetInstrumentationBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerSetInstrumentationBreakpointArgs(func() (eventName string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerSetInstrumentationBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerRemoveInstrumentationBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerRemoveInstrumentationBreakpointArgs(func() (eventName string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerRemoveInstrumentationBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerSetXHRBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerSetXHRBreakpointArgs(func() (url string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerSetXHRBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerRemoveXHRBreakpointArgs(t *testing.T) {
	args := NewDOMDebuggerRemoveXHRBreakpointArgs(func() (url string) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerRemoveXHRBreakpointArgs returned nil args")
	}
}

func TestNewDOMDebuggerGetEventListenersArgs(t *testing.T) {
	args := NewDOMDebuggerGetEventListenersArgs(func() (objectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewDOMDebuggerGetEventListenersArgs returned nil args")
	}
}

func TestNewDOMStorageClearArgs(t *testing.T) {
	args := NewDOMStorageClearArgs(func() (storageID cdptype.DOMStorageStorageID) { return }())
	if args == nil {
		t.Errorf("NewDOMStorageClearArgs returned nil args")
	}
}

func TestNewDOMStorageGetDOMStorageItemsArgs(t *testing.T) {
	args := NewDOMStorageGetDOMStorageItemsArgs(func() (storageID cdptype.DOMStorageStorageID) { return }())
	if args == nil {
		t.Errorf("NewDOMStorageGetDOMStorageItemsArgs returned nil args")
	}
}

func TestNewDOMStorageSetDOMStorageItemArgs(t *testing.T) {
	args := NewDOMStorageSetDOMStorageItemArgs(func() (storageID cdptype.DOMStorageStorageID, key string, value string) { return }())
	if args == nil {
		t.Errorf("NewDOMStorageSetDOMStorageItemArgs returned nil args")
	}
}

func TestNewDOMStorageRemoveDOMStorageItemArgs(t *testing.T) {
	args := NewDOMStorageRemoveDOMStorageItemArgs(func() (storageID cdptype.DOMStorageStorageID, key string) { return }())
	if args == nil {
		t.Errorf("NewDOMStorageRemoveDOMStorageItemArgs returned nil args")
	}
}

func TestNewDatabaseGetDatabaseTableNamesArgs(t *testing.T) {
	args := NewDatabaseGetDatabaseTableNamesArgs(func() (databaseID cdptype.DatabaseID) { return }())
	if args == nil {
		t.Errorf("NewDatabaseGetDatabaseTableNamesArgs returned nil args")
	}
}

func TestNewDatabaseExecuteSQLArgs(t *testing.T) {
	args := NewDatabaseExecuteSQLArgs(func() (databaseID cdptype.DatabaseID, query string) { return }())
	if args == nil {
		t.Errorf("NewDatabaseExecuteSQLArgs returned nil args")
	}
}

func TestNewDebuggerSetBreakpointsActiveArgs(t *testing.T) {
	args := NewDebuggerSetBreakpointsActiveArgs(func() (active bool) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetBreakpointsActiveArgs returned nil args")
	}
}

func TestNewDebuggerSetSkipAllPausesArgs(t *testing.T) {
	args := NewDebuggerSetSkipAllPausesArgs(func() (skip bool) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetSkipAllPausesArgs returned nil args")
	}
}

func TestNewDebuggerSetBreakpointByURLArgs(t *testing.T) {
	args := NewDebuggerSetBreakpointByURLArgs(func() (lineNumber int) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetBreakpointByURLArgs returned nil args")
	}
}

func TestNewDebuggerSetBreakpointArgs(t *testing.T) {
	args := NewDebuggerSetBreakpointArgs(func() (location cdptype.DebuggerLocation) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetBreakpointArgs returned nil args")
	}
}

func TestNewDebuggerRemoveBreakpointArgs(t *testing.T) {
	args := NewDebuggerRemoveBreakpointArgs(func() (breakpointID cdptype.DebuggerBreakpointID) { return }())
	if args == nil {
		t.Errorf("NewDebuggerRemoveBreakpointArgs returned nil args")
	}
}

func TestNewDebuggerGetPossibleBreakpointsArgs(t *testing.T) {
	args := NewDebuggerGetPossibleBreakpointsArgs(func() (start cdptype.DebuggerLocation) { return }())
	if args == nil {
		t.Errorf("NewDebuggerGetPossibleBreakpointsArgs returned nil args")
	}
}

func TestNewDebuggerContinueToLocationArgs(t *testing.T) {
	args := NewDebuggerContinueToLocationArgs(func() (location cdptype.DebuggerLocation) { return }())
	if args == nil {
		t.Errorf("NewDebuggerContinueToLocationArgs returned nil args")
	}
}

func TestNewDebuggerSearchInContentArgs(t *testing.T) {
	args := NewDebuggerSearchInContentArgs(func() (scriptID cdptype.RuntimeScriptID, query string) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSearchInContentArgs returned nil args")
	}
}

func TestNewDebuggerSetScriptSourceArgs(t *testing.T) {
	args := NewDebuggerSetScriptSourceArgs(func() (scriptID cdptype.RuntimeScriptID, scriptSource string) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetScriptSourceArgs returned nil args")
	}
}

func TestNewDebuggerRestartFrameArgs(t *testing.T) {
	args := NewDebuggerRestartFrameArgs(func() (callFrameID cdptype.DebuggerCallFrameID) { return }())
	if args == nil {
		t.Errorf("NewDebuggerRestartFrameArgs returned nil args")
	}
}

func TestNewDebuggerGetScriptSourceArgs(t *testing.T) {
	args := NewDebuggerGetScriptSourceArgs(func() (scriptID cdptype.RuntimeScriptID) { return }())
	if args == nil {
		t.Errorf("NewDebuggerGetScriptSourceArgs returned nil args")
	}
}

func TestNewDebuggerSetPauseOnExceptionsArgs(t *testing.T) {
	args := NewDebuggerSetPauseOnExceptionsArgs(func() (state string) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetPauseOnExceptionsArgs returned nil args")
	}
}

func TestNewDebuggerEvaluateOnCallFrameArgs(t *testing.T) {
	args := NewDebuggerEvaluateOnCallFrameArgs(func() (callFrameID cdptype.DebuggerCallFrameID, expression string) { return }())
	if args == nil {
		t.Errorf("NewDebuggerEvaluateOnCallFrameArgs returned nil args")
	}
}

func TestNewDebuggerSetVariableValueArgs(t *testing.T) {
	args := NewDebuggerSetVariableValueArgs(func() (scopeNumber int, variableName string, newValue cdptype.RuntimeCallArgument, callFrameID cdptype.DebuggerCallFrameID) {
		return
	}())
	if args == nil {
		t.Errorf("NewDebuggerSetVariableValueArgs returned nil args")
	}
}

func TestNewDebuggerSetAsyncCallStackDepthArgs(t *testing.T) {
	args := NewDebuggerSetAsyncCallStackDepthArgs(func() (maxDepth int) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetAsyncCallStackDepthArgs returned nil args")
	}
}

func TestNewDebuggerSetBlackboxPatternsArgs(t *testing.T) {
	args := NewDebuggerSetBlackboxPatternsArgs(func() (patterns []string) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetBlackboxPatternsArgs returned nil args")
	}
}

func TestNewDebuggerSetBlackboxedRangesArgs(t *testing.T) {
	args := NewDebuggerSetBlackboxedRangesArgs(func() (scriptID cdptype.RuntimeScriptID, positions []cdptype.DebuggerScriptPosition) { return }())
	if args == nil {
		t.Errorf("NewDebuggerSetBlackboxedRangesArgs returned nil args")
	}
}

func TestNewDeviceOrientationSetDeviceOrientationOverrideArgs(t *testing.T) {
	args := NewDeviceOrientationSetDeviceOrientationOverrideArgs(func() (alpha float64, beta float64, gamma float64) { return }())
	if args == nil {
		t.Errorf("NewDeviceOrientationSetDeviceOrientationOverrideArgs returned nil args")
	}
}

func TestNewEmulationSetDeviceMetricsOverrideArgs(t *testing.T) {
	args := NewEmulationSetDeviceMetricsOverrideArgs(func() (width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetDeviceMetricsOverrideArgs returned nil args")
	}
}

func TestNewEmulationForceViewportArgs(t *testing.T) {
	args := NewEmulationForceViewportArgs(func() (x float64, y float64, scale float64) { return }())
	if args == nil {
		t.Errorf("NewEmulationForceViewportArgs returned nil args")
	}
}

func TestNewEmulationSetPageScaleFactorArgs(t *testing.T) {
	args := NewEmulationSetPageScaleFactorArgs(func() (pageScaleFactor float64) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetPageScaleFactorArgs returned nil args")
	}
}

func TestNewEmulationSetVisibleSizeArgs(t *testing.T) {
	args := NewEmulationSetVisibleSizeArgs(func() (width int, height int) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetVisibleSizeArgs returned nil args")
	}
}

func TestNewEmulationSetScriptExecutionDisabledArgs(t *testing.T) {
	args := NewEmulationSetScriptExecutionDisabledArgs(func() (value bool) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetScriptExecutionDisabledArgs returned nil args")
	}
}

func TestNewEmulationSetGeolocationOverrideArgs(t *testing.T) {
	args := NewEmulationSetGeolocationOverrideArgs()
	if args == nil {
		t.Errorf("NewEmulationSetGeolocationOverrideArgs returned nil args")
	}
}

func TestNewEmulationSetTouchEmulationEnabledArgs(t *testing.T) {
	args := NewEmulationSetTouchEmulationEnabledArgs(func() (enabled bool) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetTouchEmulationEnabledArgs returned nil args")
	}
}

func TestNewEmulationSetEmulatedMediaArgs(t *testing.T) {
	args := NewEmulationSetEmulatedMediaArgs(func() (media string) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetEmulatedMediaArgs returned nil args")
	}
}

func TestNewEmulationSetCPUThrottlingRateArgs(t *testing.T) {
	args := NewEmulationSetCPUThrottlingRateArgs(func() (rate float64) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetCPUThrottlingRateArgs returned nil args")
	}
}

func TestNewEmulationSetVirtualTimePolicyArgs(t *testing.T) {
	args := NewEmulationSetVirtualTimePolicyArgs(func() (policy cdptype.EmulationVirtualTimePolicy) { return }())
	if args == nil {
		t.Errorf("NewEmulationSetVirtualTimePolicyArgs returned nil args")
	}
}

func TestNewEmulationSetDefaultBackgroundColorOverrideArgs(t *testing.T) {
	args := NewEmulationSetDefaultBackgroundColorOverrideArgs()
	if args == nil {
		t.Errorf("NewEmulationSetDefaultBackgroundColorOverrideArgs returned nil args")
	}
}

func TestNewHeapProfilerStartTrackingHeapObjectsArgs(t *testing.T) {
	args := NewHeapProfilerStartTrackingHeapObjectsArgs()
	if args == nil {
		t.Errorf("NewHeapProfilerStartTrackingHeapObjectsArgs returned nil args")
	}
}

func TestNewHeapProfilerStopTrackingHeapObjectsArgs(t *testing.T) {
	args := NewHeapProfilerStopTrackingHeapObjectsArgs()
	if args == nil {
		t.Errorf("NewHeapProfilerStopTrackingHeapObjectsArgs returned nil args")
	}
}

func TestNewHeapProfilerTakeHeapSnapshotArgs(t *testing.T) {
	args := NewHeapProfilerTakeHeapSnapshotArgs()
	if args == nil {
		t.Errorf("NewHeapProfilerTakeHeapSnapshotArgs returned nil args")
	}
}

func TestNewHeapProfilerGetObjectByHeapObjectIDArgs(t *testing.T) {
	args := NewHeapProfilerGetObjectByHeapObjectIDArgs(func() (objectID cdptype.HeapProfilerHeapSnapshotObjectID) { return }())
	if args == nil {
		t.Errorf("NewHeapProfilerGetObjectByHeapObjectIDArgs returned nil args")
	}
}

func TestNewHeapProfilerAddInspectedHeapObjectArgs(t *testing.T) {
	args := NewHeapProfilerAddInspectedHeapObjectArgs(func() (heapObjectID cdptype.HeapProfilerHeapSnapshotObjectID) { return }())
	if args == nil {
		t.Errorf("NewHeapProfilerAddInspectedHeapObjectArgs returned nil args")
	}
}

func TestNewHeapProfilerGetHeapObjectIDArgs(t *testing.T) {
	args := NewHeapProfilerGetHeapObjectIDArgs(func() (objectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewHeapProfilerGetHeapObjectIDArgs returned nil args")
	}
}

func TestNewHeapProfilerStartSamplingArgs(t *testing.T) {
	args := NewHeapProfilerStartSamplingArgs()
	if args == nil {
		t.Errorf("NewHeapProfilerStartSamplingArgs returned nil args")
	}
}

func TestNewIOReadArgs(t *testing.T) {
	args := NewIOReadArgs(func() (handle cdptype.IOStreamHandle) { return }())
	if args == nil {
		t.Errorf("NewIOReadArgs returned nil args")
	}
}

func TestNewIOCloseArgs(t *testing.T) {
	args := NewIOCloseArgs(func() (handle cdptype.IOStreamHandle) { return }())
	if args == nil {
		t.Errorf("NewIOCloseArgs returned nil args")
	}
}

func TestNewIndexedDBRequestDatabaseNamesArgs(t *testing.T) {
	args := NewIndexedDBRequestDatabaseNamesArgs(func() (securityOrigin string) { return }())
	if args == nil {
		t.Errorf("NewIndexedDBRequestDatabaseNamesArgs returned nil args")
	}
}

func TestNewIndexedDBRequestDatabaseArgs(t *testing.T) {
	args := NewIndexedDBRequestDatabaseArgs(func() (securityOrigin string, databaseName string) { return }())
	if args == nil {
		t.Errorf("NewIndexedDBRequestDatabaseArgs returned nil args")
	}
}

func TestNewIndexedDBRequestDataArgs(t *testing.T) {
	args := NewIndexedDBRequestDataArgs(func() (securityOrigin string, databaseName string, objectStoreName string, indexName string, skipCount int, pageSize int) {
		return
	}())
	if args == nil {
		t.Errorf("NewIndexedDBRequestDataArgs returned nil args")
	}
}

func TestNewIndexedDBClearObjectStoreArgs(t *testing.T) {
	args := NewIndexedDBClearObjectStoreArgs(func() (securityOrigin string, databaseName string, objectStoreName string) { return }())
	if args == nil {
		t.Errorf("NewIndexedDBClearObjectStoreArgs returned nil args")
	}
}

func TestNewIndexedDBDeleteDatabaseArgs(t *testing.T) {
	args := NewIndexedDBDeleteDatabaseArgs(func() (securityOrigin string, databaseName string) { return }())
	if args == nil {
		t.Errorf("NewIndexedDBDeleteDatabaseArgs returned nil args")
	}
}

func TestNewInputSetIgnoreInputEventsArgs(t *testing.T) {
	args := NewInputSetIgnoreInputEventsArgs(func() (ignore bool) { return }())
	if args == nil {
		t.Errorf("NewInputSetIgnoreInputEventsArgs returned nil args")
	}
}

func TestNewInputDispatchKeyEventArgs(t *testing.T) {
	args := NewInputDispatchKeyEventArgs(func() (typ string) { return }())
	if args == nil {
		t.Errorf("NewInputDispatchKeyEventArgs returned nil args")
	}
}

func TestNewInputDispatchMouseEventArgs(t *testing.T) {
	args := NewInputDispatchMouseEventArgs(func() (typ string, x int, y int) { return }())
	if args == nil {
		t.Errorf("NewInputDispatchMouseEventArgs returned nil args")
	}
}

func TestNewInputDispatchTouchEventArgs(t *testing.T) {
	args := NewInputDispatchTouchEventArgs(func() (typ string, touchPoints []cdptype.InputTouchPoint) { return }())
	if args == nil {
		t.Errorf("NewInputDispatchTouchEventArgs returned nil args")
	}
}

func TestNewInputEmulateTouchFromMouseEventArgs(t *testing.T) {
	args := NewInputEmulateTouchFromMouseEventArgs(func() (typ string, x int, y int, timestamp cdptype.Timestamp, button string) { return }())
	if args == nil {
		t.Errorf("NewInputEmulateTouchFromMouseEventArgs returned nil args")
	}
}

func TestNewInputSynthesizePinchGestureArgs(t *testing.T) {
	args := NewInputSynthesizePinchGestureArgs(func() (x int, y int, scaleFactor float64) { return }())
	if args == nil {
		t.Errorf("NewInputSynthesizePinchGestureArgs returned nil args")
	}
}

func TestNewInputSynthesizeScrollGestureArgs(t *testing.T) {
	args := NewInputSynthesizeScrollGestureArgs(func() (x int, y int) { return }())
	if args == nil {
		t.Errorf("NewInputSynthesizeScrollGestureArgs returned nil args")
	}
}

func TestNewInputSynthesizeTapGestureArgs(t *testing.T) {
	args := NewInputSynthesizeTapGestureArgs(func() (x int, y int) { return }())
	if args == nil {
		t.Errorf("NewInputSynthesizeTapGestureArgs returned nil args")
	}
}

func TestNewLayerTreeCompositingReasonsArgs(t *testing.T) {
	args := NewLayerTreeCompositingReasonsArgs(func() (layerID cdptype.LayerTreeLayerID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeCompositingReasonsArgs returned nil args")
	}
}

func TestNewLayerTreeMakeSnapshotArgs(t *testing.T) {
	args := NewLayerTreeMakeSnapshotArgs(func() (layerID cdptype.LayerTreeLayerID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeMakeSnapshotArgs returned nil args")
	}
}

func TestNewLayerTreeLoadSnapshotArgs(t *testing.T) {
	args := NewLayerTreeLoadSnapshotArgs(func() (tiles []cdptype.LayerTreePictureTile) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeLoadSnapshotArgs returned nil args")
	}
}

func TestNewLayerTreeReleaseSnapshotArgs(t *testing.T) {
	args := NewLayerTreeReleaseSnapshotArgs(func() (snapshotID cdptype.LayerTreeSnapshotID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeReleaseSnapshotArgs returned nil args")
	}
}

func TestNewLayerTreeProfileSnapshotArgs(t *testing.T) {
	args := NewLayerTreeProfileSnapshotArgs(func() (snapshotID cdptype.LayerTreeSnapshotID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeProfileSnapshotArgs returned nil args")
	}
}

func TestNewLayerTreeReplaySnapshotArgs(t *testing.T) {
	args := NewLayerTreeReplaySnapshotArgs(func() (snapshotID cdptype.LayerTreeSnapshotID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeReplaySnapshotArgs returned nil args")
	}
}

func TestNewLayerTreeSnapshotCommandLogArgs(t *testing.T) {
	args := NewLayerTreeSnapshotCommandLogArgs(func() (snapshotID cdptype.LayerTreeSnapshotID) { return }())
	if args == nil {
		t.Errorf("NewLayerTreeSnapshotCommandLogArgs returned nil args")
	}
}

func TestNewLogStartViolationsReportArgs(t *testing.T) {
	args := NewLogStartViolationsReportArgs(func() (config []cdptype.LogViolationSetting) { return }())
	if args == nil {
		t.Errorf("NewLogStartViolationsReportArgs returned nil args")
	}
}

func TestNewMemorySetPressureNotificationsSuppressedArgs(t *testing.T) {
	args := NewMemorySetPressureNotificationsSuppressedArgs(func() (suppressed bool) { return }())
	if args == nil {
		t.Errorf("NewMemorySetPressureNotificationsSuppressedArgs returned nil args")
	}
}

func TestNewMemorySimulatePressureNotificationArgs(t *testing.T) {
	args := NewMemorySimulatePressureNotificationArgs(func() (level cdptype.MemoryPressureLevel) { return }())
	if args == nil {
		t.Errorf("NewMemorySimulatePressureNotificationArgs returned nil args")
	}
}

func TestNewNetworkEnableArgs(t *testing.T) {
	args := NewNetworkEnableArgs()
	if args == nil {
		t.Errorf("NewNetworkEnableArgs returned nil args")
	}
}

func TestNewNetworkSetUserAgentOverrideArgs(t *testing.T) {
	args := NewNetworkSetUserAgentOverrideArgs(func() (userAgent string) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetUserAgentOverrideArgs returned nil args")
	}
}

func TestNewNetworkSetExtraHTTPHeadersArgs(t *testing.T) {
	args := NewNetworkSetExtraHTTPHeadersArgs(func() (headers cdptype.NetworkHeaders) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetExtraHTTPHeadersArgs returned nil args")
	}
}

func TestNewNetworkGetResponseBodyArgs(t *testing.T) {
	args := NewNetworkGetResponseBodyArgs(func() (requestID cdptype.NetworkRequestID) { return }())
	if args == nil {
		t.Errorf("NewNetworkGetResponseBodyArgs returned nil args")
	}
}

func TestNewNetworkSetBlockedURLsArgs(t *testing.T) {
	args := NewNetworkSetBlockedURLsArgs(func() (urls []string) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetBlockedURLsArgs returned nil args")
	}
}

func TestNewNetworkReplayXHRArgs(t *testing.T) {
	args := NewNetworkReplayXHRArgs(func() (requestID cdptype.NetworkRequestID) { return }())
	if args == nil {
		t.Errorf("NewNetworkReplayXHRArgs returned nil args")
	}
}

func TestNewNetworkGetCookiesArgs(t *testing.T) {
	args := NewNetworkGetCookiesArgs()
	if args == nil {
		t.Errorf("NewNetworkGetCookiesArgs returned nil args")
	}
}

func TestNewNetworkDeleteCookieArgs(t *testing.T) {
	args := NewNetworkDeleteCookieArgs(func() (cookieName string, url string) { return }())
	if args == nil {
		t.Errorf("NewNetworkDeleteCookieArgs returned nil args")
	}
}

func TestNewNetworkSetCookieArgs(t *testing.T) {
	args := NewNetworkSetCookieArgs(func() (url string, name string, value string) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetCookieArgs returned nil args")
	}
}

func TestNewNetworkEmulateNetworkConditionsArgs(t *testing.T) {
	args := NewNetworkEmulateNetworkConditionsArgs(func() (offline bool, latency float64, downloadThroughput float64, uploadThroughput float64) { return }())
	if args == nil {
		t.Errorf("NewNetworkEmulateNetworkConditionsArgs returned nil args")
	}
}

func TestNewNetworkSetCacheDisabledArgs(t *testing.T) {
	args := NewNetworkSetCacheDisabledArgs(func() (cacheDisabled bool) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetCacheDisabledArgs returned nil args")
	}
}

func TestNewNetworkSetBypassServiceWorkerArgs(t *testing.T) {
	args := NewNetworkSetBypassServiceWorkerArgs(func() (bypass bool) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetBypassServiceWorkerArgs returned nil args")
	}
}

func TestNewNetworkSetDataSizeLimitsForTestArgs(t *testing.T) {
	args := NewNetworkSetDataSizeLimitsForTestArgs(func() (maxTotalSize int, maxResourceSize int) { return }())
	if args == nil {
		t.Errorf("NewNetworkSetDataSizeLimitsForTestArgs returned nil args")
	}
}

func TestNewNetworkGetCertificateArgs(t *testing.T) {
	args := NewNetworkGetCertificateArgs(func() (origin string) { return }())
	if args == nil {
		t.Errorf("NewNetworkGetCertificateArgs returned nil args")
	}
}

func TestNewOverlaySetShowPaintRectsArgs(t *testing.T) {
	args := NewOverlaySetShowPaintRectsArgs(func() (result bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetShowPaintRectsArgs returned nil args")
	}
}

func TestNewOverlaySetShowDebugBordersArgs(t *testing.T) {
	args := NewOverlaySetShowDebugBordersArgs(func() (show bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetShowDebugBordersArgs returned nil args")
	}
}

func TestNewOverlaySetShowFPSCounterArgs(t *testing.T) {
	args := NewOverlaySetShowFPSCounterArgs(func() (show bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetShowFPSCounterArgs returned nil args")
	}
}

func TestNewOverlaySetShowScrollBottleneckRectsArgs(t *testing.T) {
	args := NewOverlaySetShowScrollBottleneckRectsArgs(func() (show bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetShowScrollBottleneckRectsArgs returned nil args")
	}
}

func TestNewOverlaySetShowViewportSizeOnResizeArgs(t *testing.T) {
	args := NewOverlaySetShowViewportSizeOnResizeArgs(func() (show bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetShowViewportSizeOnResizeArgs returned nil args")
	}
}

func TestNewOverlaySetPausedInDebuggerMessageArgs(t *testing.T) {
	args := NewOverlaySetPausedInDebuggerMessageArgs()
	if args == nil {
		t.Errorf("NewOverlaySetPausedInDebuggerMessageArgs returned nil args")
	}
}

func TestNewOverlaySetSuspendedArgs(t *testing.T) {
	args := NewOverlaySetSuspendedArgs(func() (suspended bool) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetSuspendedArgs returned nil args")
	}
}

func TestNewOverlaySetInspectModeArgs(t *testing.T) {
	args := NewOverlaySetInspectModeArgs(func() (mode cdptype.OverlayInspectMode) { return }())
	if args == nil {
		t.Errorf("NewOverlaySetInspectModeArgs returned nil args")
	}
}

func TestNewOverlayHighlightRectArgs(t *testing.T) {
	args := NewOverlayHighlightRectArgs(func() (x int, y int, width int, height int) { return }())
	if args == nil {
		t.Errorf("NewOverlayHighlightRectArgs returned nil args")
	}
}

func TestNewOverlayHighlightQuadArgs(t *testing.T) {
	args := NewOverlayHighlightQuadArgs(func() (quad cdptype.DOMQuad) { return }())
	if args == nil {
		t.Errorf("NewOverlayHighlightQuadArgs returned nil args")
	}
}

func TestNewOverlayHighlightNodeArgs(t *testing.T) {
	args := NewOverlayHighlightNodeArgs(func() (highlightConfig cdptype.OverlayHighlightConfig) { return }())
	if args == nil {
		t.Errorf("NewOverlayHighlightNodeArgs returned nil args")
	}
}

func TestNewOverlayHighlightFrameArgs(t *testing.T) {
	args := NewOverlayHighlightFrameArgs(func() (frameID cdptype.PageFrameID) { return }())
	if args == nil {
		t.Errorf("NewOverlayHighlightFrameArgs returned nil args")
	}
}

func TestNewOverlayGetHighlightObjectForTestArgs(t *testing.T) {
	args := NewOverlayGetHighlightObjectForTestArgs(func() (nodeID cdptype.DOMNodeID) { return }())
	if args == nil {
		t.Errorf("NewOverlayGetHighlightObjectForTestArgs returned nil args")
	}
}

func TestNewPageAddScriptToEvaluateOnLoadArgs(t *testing.T) {
	args := NewPageAddScriptToEvaluateOnLoadArgs(func() (scriptSource string) { return }())
	if args == nil {
		t.Errorf("NewPageAddScriptToEvaluateOnLoadArgs returned nil args")
	}
}

func TestNewPageRemoveScriptToEvaluateOnLoadArgs(t *testing.T) {
	args := NewPageRemoveScriptToEvaluateOnLoadArgs(func() (identifier cdptype.PageScriptIdentifier) { return }())
	if args == nil {
		t.Errorf("NewPageRemoveScriptToEvaluateOnLoadArgs returned nil args")
	}
}

func TestNewPageSetAutoAttachToCreatedPagesArgs(t *testing.T) {
	args := NewPageSetAutoAttachToCreatedPagesArgs(func() (autoAttach bool) { return }())
	if args == nil {
		t.Errorf("NewPageSetAutoAttachToCreatedPagesArgs returned nil args")
	}
}

func TestNewPageReloadArgs(t *testing.T) {
	args := NewPageReloadArgs()
	if args == nil {
		t.Errorf("NewPageReloadArgs returned nil args")
	}
}

func TestNewPageNavigateArgs(t *testing.T) {
	args := NewPageNavigateArgs(func() (url string) { return }())
	if args == nil {
		t.Errorf("NewPageNavigateArgs returned nil args")
	}
}

func TestNewPageNavigateToHistoryEntryArgs(t *testing.T) {
	args := NewPageNavigateToHistoryEntryArgs(func() (entryID int) { return }())
	if args == nil {
		t.Errorf("NewPageNavigateToHistoryEntryArgs returned nil args")
	}
}

func TestNewPageDeleteCookieArgs(t *testing.T) {
	args := NewPageDeleteCookieArgs(func() (cookieName string, url string) { return }())
	if args == nil {
		t.Errorf("NewPageDeleteCookieArgs returned nil args")
	}
}

func TestNewPageGetResourceContentArgs(t *testing.T) {
	args := NewPageGetResourceContentArgs(func() (frameID cdptype.PageFrameID, url string) { return }())
	if args == nil {
		t.Errorf("NewPageGetResourceContentArgs returned nil args")
	}
}

func TestNewPageSearchInResourceArgs(t *testing.T) {
	args := NewPageSearchInResourceArgs(func() (frameID cdptype.PageFrameID, url string, query string) { return }())
	if args == nil {
		t.Errorf("NewPageSearchInResourceArgs returned nil args")
	}
}

func TestNewPageSetDocumentContentArgs(t *testing.T) {
	args := NewPageSetDocumentContentArgs(func() (frameID cdptype.PageFrameID, html string) { return }())
	if args == nil {
		t.Errorf("NewPageSetDocumentContentArgs returned nil args")
	}
}

func TestNewPageSetDeviceMetricsOverrideArgs(t *testing.T) {
	args := NewPageSetDeviceMetricsOverrideArgs(func() (width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool) { return }())
	if args == nil {
		t.Errorf("NewPageSetDeviceMetricsOverrideArgs returned nil args")
	}
}

func TestNewPageSetGeolocationOverrideArgs(t *testing.T) {
	args := NewPageSetGeolocationOverrideArgs()
	if args == nil {
		t.Errorf("NewPageSetGeolocationOverrideArgs returned nil args")
	}
}

func TestNewPageSetDeviceOrientationOverrideArgs(t *testing.T) {
	args := NewPageSetDeviceOrientationOverrideArgs(func() (alpha float64, beta float64, gamma float64) { return }())
	if args == nil {
		t.Errorf("NewPageSetDeviceOrientationOverrideArgs returned nil args")
	}
}

func TestNewPageSetTouchEmulationEnabledArgs(t *testing.T) {
	args := NewPageSetTouchEmulationEnabledArgs(func() (enabled bool) { return }())
	if args == nil {
		t.Errorf("NewPageSetTouchEmulationEnabledArgs returned nil args")
	}
}

func TestNewPageCaptureScreenshotArgs(t *testing.T) {
	args := NewPageCaptureScreenshotArgs()
	if args == nil {
		t.Errorf("NewPageCaptureScreenshotArgs returned nil args")
	}
}

func TestNewPagePrintToPDFArgs(t *testing.T) {
	args := NewPagePrintToPDFArgs()
	if args == nil {
		t.Errorf("NewPagePrintToPDFArgs returned nil args")
	}
}

func TestNewPageStartScreencastArgs(t *testing.T) {
	args := NewPageStartScreencastArgs()
	if args == nil {
		t.Errorf("NewPageStartScreencastArgs returned nil args")
	}
}

func TestNewPageScreencastFrameAckArgs(t *testing.T) {
	args := NewPageScreencastFrameAckArgs(func() (sessionID int) { return }())
	if args == nil {
		t.Errorf("NewPageScreencastFrameAckArgs returned nil args")
	}
}

func TestNewPageHandleJavaScriptDialogArgs(t *testing.T) {
	args := NewPageHandleJavaScriptDialogArgs(func() (accept bool) { return }())
	if args == nil {
		t.Errorf("NewPageHandleJavaScriptDialogArgs returned nil args")
	}
}

func TestNewPageSetControlNavigationsArgs(t *testing.T) {
	args := NewPageSetControlNavigationsArgs(func() (enabled bool) { return }())
	if args == nil {
		t.Errorf("NewPageSetControlNavigationsArgs returned nil args")
	}
}

func TestNewPageProcessNavigationArgs(t *testing.T) {
	args := NewPageProcessNavigationArgs(func() (response cdptype.PageNavigationResponse, navigationID int) { return }())
	if args == nil {
		t.Errorf("NewPageProcessNavigationArgs returned nil args")
	}
}

func TestNewPageCreateIsolatedWorldArgs(t *testing.T) {
	args := NewPageCreateIsolatedWorldArgs(func() (frameID cdptype.PageFrameID) { return }())
	if args == nil {
		t.Errorf("NewPageCreateIsolatedWorldArgs returned nil args")
	}
}

func TestNewProfilerSetSamplingIntervalArgs(t *testing.T) {
	args := NewProfilerSetSamplingIntervalArgs(func() (interval int) { return }())
	if args == nil {
		t.Errorf("NewProfilerSetSamplingIntervalArgs returned nil args")
	}
}

func TestNewProfilerStartPreciseCoverageArgs(t *testing.T) {
	args := NewProfilerStartPreciseCoverageArgs()
	if args == nil {
		t.Errorf("NewProfilerStartPreciseCoverageArgs returned nil args")
	}
}

func TestNewRuntimeEvaluateArgs(t *testing.T) {
	args := NewRuntimeEvaluateArgs(func() (expression string) { return }())
	if args == nil {
		t.Errorf("NewRuntimeEvaluateArgs returned nil args")
	}
}

func TestNewRuntimeAwaitPromiseArgs(t *testing.T) {
	args := NewRuntimeAwaitPromiseArgs(func() (promiseObjectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewRuntimeAwaitPromiseArgs returned nil args")
	}
}

func TestNewRuntimeCallFunctionOnArgs(t *testing.T) {
	args := NewRuntimeCallFunctionOnArgs(func() (objectID cdptype.RuntimeRemoteObjectID, functionDeclaration string) { return }())
	if args == nil {
		t.Errorf("NewRuntimeCallFunctionOnArgs returned nil args")
	}
}

func TestNewRuntimeGetPropertiesArgs(t *testing.T) {
	args := NewRuntimeGetPropertiesArgs(func() (objectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewRuntimeGetPropertiesArgs returned nil args")
	}
}

func TestNewRuntimeReleaseObjectArgs(t *testing.T) {
	args := NewRuntimeReleaseObjectArgs(func() (objectID cdptype.RuntimeRemoteObjectID) { return }())
	if args == nil {
		t.Errorf("NewRuntimeReleaseObjectArgs returned nil args")
	}
}

func TestNewRuntimeReleaseObjectGroupArgs(t *testing.T) {
	args := NewRuntimeReleaseObjectGroupArgs(func() (objectGroup string) { return }())
	if args == nil {
		t.Errorf("NewRuntimeReleaseObjectGroupArgs returned nil args")
	}
}

func TestNewRuntimeSetCustomObjectFormatterEnabledArgs(t *testing.T) {
	args := NewRuntimeSetCustomObjectFormatterEnabledArgs(func() (enabled bool) { return }())
	if args == nil {
		t.Errorf("NewRuntimeSetCustomObjectFormatterEnabledArgs returned nil args")
	}
}

func TestNewRuntimeCompileScriptArgs(t *testing.T) {
	args := NewRuntimeCompileScriptArgs(func() (expression string, sourceURL string, persistScript bool) { return }())
	if args == nil {
		t.Errorf("NewRuntimeCompileScriptArgs returned nil args")
	}
}

func TestNewRuntimeRunScriptArgs(t *testing.T) {
	args := NewRuntimeRunScriptArgs(func() (scriptID cdptype.RuntimeScriptID) { return }())
	if args == nil {
		t.Errorf("NewRuntimeRunScriptArgs returned nil args")
	}
}

func TestNewSecurityHandleCertificateErrorArgs(t *testing.T) {
	args := NewSecurityHandleCertificateErrorArgs(func() (eventID int, action cdptype.SecurityCertificateErrorAction) { return }())
	if args == nil {
		t.Errorf("NewSecurityHandleCertificateErrorArgs returned nil args")
	}
}

func TestNewSecuritySetOverrideCertificateErrorsArgs(t *testing.T) {
	args := NewSecuritySetOverrideCertificateErrorsArgs(func() (override bool) { return }())
	if args == nil {
		t.Errorf("NewSecuritySetOverrideCertificateErrorsArgs returned nil args")
	}
}

func TestNewServiceWorkerUnregisterArgs(t *testing.T) {
	args := NewServiceWorkerUnregisterArgs(func() (scopeURL string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerUnregisterArgs returned nil args")
	}
}

func TestNewServiceWorkerUpdateRegistrationArgs(t *testing.T) {
	args := NewServiceWorkerUpdateRegistrationArgs(func() (scopeURL string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerUpdateRegistrationArgs returned nil args")
	}
}

func TestNewServiceWorkerStartWorkerArgs(t *testing.T) {
	args := NewServiceWorkerStartWorkerArgs(func() (scopeURL string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerStartWorkerArgs returned nil args")
	}
}

func TestNewServiceWorkerSkipWaitingArgs(t *testing.T) {
	args := NewServiceWorkerSkipWaitingArgs(func() (scopeURL string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerSkipWaitingArgs returned nil args")
	}
}

func TestNewServiceWorkerStopWorkerArgs(t *testing.T) {
	args := NewServiceWorkerStopWorkerArgs(func() (versionID string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerStopWorkerArgs returned nil args")
	}
}

func TestNewServiceWorkerInspectWorkerArgs(t *testing.T) {
	args := NewServiceWorkerInspectWorkerArgs(func() (versionID string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerInspectWorkerArgs returned nil args")
	}
}

func TestNewServiceWorkerSetForceUpdateOnPageLoadArgs(t *testing.T) {
	args := NewServiceWorkerSetForceUpdateOnPageLoadArgs(func() (forceUpdateOnPageLoad bool) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerSetForceUpdateOnPageLoadArgs returned nil args")
	}
}

func TestNewServiceWorkerDeliverPushMessageArgs(t *testing.T) {
	args := NewServiceWorkerDeliverPushMessageArgs(func() (origin string, registrationID string, data string) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerDeliverPushMessageArgs returned nil args")
	}
}

func TestNewServiceWorkerDispatchSyncEventArgs(t *testing.T) {
	args := NewServiceWorkerDispatchSyncEventArgs(func() (origin string, registrationID string, tag string, lastChance bool) { return }())
	if args == nil {
		t.Errorf("NewServiceWorkerDispatchSyncEventArgs returned nil args")
	}
}

func TestNewStorageClearDataForOriginArgs(t *testing.T) {
	args := NewStorageClearDataForOriginArgs(func() (origin string, storageTypes string) { return }())
	if args == nil {
		t.Errorf("NewStorageClearDataForOriginArgs returned nil args")
	}
}

func TestNewTargetSetDiscoverTargetsArgs(t *testing.T) {
	args := NewTargetSetDiscoverTargetsArgs(func() (discover bool) { return }())
	if args == nil {
		t.Errorf("NewTargetSetDiscoverTargetsArgs returned nil args")
	}
}

func TestNewTargetSetAutoAttachArgs(t *testing.T) {
	args := NewTargetSetAutoAttachArgs(func() (autoAttach bool, waitForDebuggerOnStart bool) { return }())
	if args == nil {
		t.Errorf("NewTargetSetAutoAttachArgs returned nil args")
	}
}

func TestNewTargetSetAttachToFramesArgs(t *testing.T) {
	args := NewTargetSetAttachToFramesArgs(func() (value bool) { return }())
	if args == nil {
		t.Errorf("NewTargetSetAttachToFramesArgs returned nil args")
	}
}

func TestNewTargetSetRemoteLocationsArgs(t *testing.T) {
	args := NewTargetSetRemoteLocationsArgs(func() (locations []cdptype.TargetRemoteLocation) { return }())
	if args == nil {
		t.Errorf("NewTargetSetRemoteLocationsArgs returned nil args")
	}
}

func TestNewTargetSendMessageToTargetArgs(t *testing.T) {
	args := NewTargetSendMessageToTargetArgs(func() (targetID cdptype.TargetID, message string) { return }())
	if args == nil {
		t.Errorf("NewTargetSendMessageToTargetArgs returned nil args")
	}
}

func TestNewTargetGetTargetInfoArgs(t *testing.T) {
	args := NewTargetGetTargetInfoArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewTargetGetTargetInfoArgs returned nil args")
	}
}

func TestNewTargetActivateTargetArgs(t *testing.T) {
	args := NewTargetActivateTargetArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewTargetActivateTargetArgs returned nil args")
	}
}

func TestNewTargetCloseTargetArgs(t *testing.T) {
	args := NewTargetCloseTargetArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewTargetCloseTargetArgs returned nil args")
	}
}

func TestNewTargetAttachToTargetArgs(t *testing.T) {
	args := NewTargetAttachToTargetArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewTargetAttachToTargetArgs returned nil args")
	}
}

func TestNewTargetDetachFromTargetArgs(t *testing.T) {
	args := NewTargetDetachFromTargetArgs(func() (targetID cdptype.TargetID) { return }())
	if args == nil {
		t.Errorf("NewTargetDetachFromTargetArgs returned nil args")
	}
}

func TestNewTargetDisposeBrowserContextArgs(t *testing.T) {
	args := NewTargetDisposeBrowserContextArgs(func() (browserContextID cdptype.TargetBrowserContextID) { return }())
	if args == nil {
		t.Errorf("NewTargetDisposeBrowserContextArgs returned nil args")
	}
}

func TestNewTargetCreateTargetArgs(t *testing.T) {
	args := NewTargetCreateTargetArgs(func() (url string) { return }())
	if args == nil {
		t.Errorf("NewTargetCreateTargetArgs returned nil args")
	}
}

func TestNewTetheringBindArgs(t *testing.T) {
	args := NewTetheringBindArgs(func() (port int) { return }())
	if args == nil {
		t.Errorf("NewTetheringBindArgs returned nil args")
	}
}

func TestNewTetheringUnbindArgs(t *testing.T) {
	args := NewTetheringUnbindArgs(func() (port int) { return }())
	if args == nil {
		t.Errorf("NewTetheringUnbindArgs returned nil args")
	}
}

func TestNewTracingStartArgs(t *testing.T) {
	args := NewTracingStartArgs()
	if args == nil {
		t.Errorf("NewTracingStartArgs returned nil args")
	}
}

func TestNewTracingRecordClockSyncMarkerArgs(t *testing.T) {
	args := NewTracingRecordClockSyncMarkerArgs(func() (syncID string) { return }())
	if args == nil {
		t.Errorf("NewTracingRecordClockSyncMarkerArgs returned nil args")
	}
}
