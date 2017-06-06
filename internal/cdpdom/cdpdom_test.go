package cdpdom

import (
	"errors"
	"testing"

	"github.com/mafredri/cdp/cdpcmd"
	"github.com/mafredri/cdp/cdpevent"
	"github.com/mafredri/cdp/rpcc"
)

func TestAccessibility_GetPartialAXTree(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAccessibility(conn)
	var err error

	// Test nil args.
	_, err = dom.GetPartialAXTree(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetPartialAXTree(nil, &cdpcmd.AccessibilityGetPartialAXTreeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetPartialAXTree(nil, &cdpcmd.AccessibilityGetPartialAXTreeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_GetPlaybackRate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	_, err = dom.GetPlaybackRate(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetPlaybackRate(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_SetPlaybackRate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	err = dom.SetPlaybackRate(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPlaybackRate(nil, &cdpcmd.AnimationSetPlaybackRateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPlaybackRate(nil, &cdpcmd.AnimationSetPlaybackRateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_GetCurrentTime(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	_, err = dom.GetCurrentTime(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetCurrentTime(nil, &cdpcmd.AnimationGetCurrentTimeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetCurrentTime(nil, &cdpcmd.AnimationGetCurrentTimeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_SetPaused(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	err = dom.SetPaused(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPaused(nil, &cdpcmd.AnimationSetPausedArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPaused(nil, &cdpcmd.AnimationSetPausedArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_SetTiming(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	err = dom.SetTiming(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetTiming(nil, &cdpcmd.AnimationSetTimingArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetTiming(nil, &cdpcmd.AnimationSetTimingArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_SeekAnimations(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	err = dom.SeekAnimations(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SeekAnimations(nil, &cdpcmd.AnimationSeekAnimationsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SeekAnimations(nil, &cdpcmd.AnimationSeekAnimationsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_ReleaseAnimations(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	err = dom.ReleaseAnimations(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ReleaseAnimations(nil, &cdpcmd.AnimationReleaseAnimationsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ReleaseAnimations(nil, &cdpcmd.AnimationReleaseAnimationsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_ResolveAnimation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)
	var err error

	// Test nil args.
	_, err = dom.ResolveAnimation(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.ResolveAnimation(nil, &cdpcmd.AnimationResolveAnimationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.ResolveAnimation(nil, &cdpcmd.AnimationResolveAnimationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestAnimation_AnimationCreated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)

	stream, err := dom.AnimationCreated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.AnimationCreated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AnimationCreated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestAnimation_AnimationStarted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)

	stream, err := dom.AnimationStarted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.AnimationStarted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AnimationStarted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestAnimation_AnimationCanceled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewAnimation(conn)

	stream, err := dom.AnimationCanceled(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.AnimationCanceled.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AnimationCanceled(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestApplicationCache_GetFramesWithManifests(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)
	var err error

	_, err = dom.GetFramesWithManifests(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetFramesWithManifests(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestApplicationCache_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestApplicationCache_GetManifestForFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)
	var err error

	// Test nil args.
	_, err = dom.GetManifestForFrame(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetManifestForFrame(nil, &cdpcmd.ApplicationCacheGetManifestForFrameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetManifestForFrame(nil, &cdpcmd.ApplicationCacheGetManifestForFrameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestApplicationCache_GetApplicationCacheForFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)
	var err error

	// Test nil args.
	_, err = dom.GetApplicationCacheForFrame(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetApplicationCacheForFrame(nil, &cdpcmd.ApplicationCacheGetApplicationCacheForFrameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetApplicationCacheForFrame(nil, &cdpcmd.ApplicationCacheGetApplicationCacheForFrameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestApplicationCache_ApplicationCacheStatusUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)

	stream, err := dom.ApplicationCacheStatusUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ApplicationCacheStatusUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ApplicationCacheStatusUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestApplicationCache_NetworkStateUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewApplicationCache(conn)

	stream, err := dom.NetworkStateUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ApplicationCacheNetworkStateUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.NetworkStateUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestBrowser_GetWindowForTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewBrowser(conn)
	var err error

	// Test nil args.
	_, err = dom.GetWindowForTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetWindowForTarget(nil, &cdpcmd.BrowserGetWindowForTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetWindowForTarget(nil, &cdpcmd.BrowserGetWindowForTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestBrowser_SetWindowBounds(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewBrowser(conn)
	var err error

	// Test nil args.
	err = dom.SetWindowBounds(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetWindowBounds(nil, &cdpcmd.BrowserSetWindowBoundsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetWindowBounds(nil, &cdpcmd.BrowserSetWindowBoundsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestBrowser_GetWindowBounds(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewBrowser(conn)
	var err error

	// Test nil args.
	_, err = dom.GetWindowBounds(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetWindowBounds(nil, &cdpcmd.BrowserGetWindowBoundsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetWindowBounds(nil, &cdpcmd.BrowserGetWindowBoundsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetMatchedStylesForNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetMatchedStylesForNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetMatchedStylesForNode(nil, &cdpcmd.CSSGetMatchedStylesForNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetMatchedStylesForNode(nil, &cdpcmd.CSSGetMatchedStylesForNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetInlineStylesForNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetInlineStylesForNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetInlineStylesForNode(nil, &cdpcmd.CSSGetInlineStylesForNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetInlineStylesForNode(nil, &cdpcmd.CSSGetInlineStylesForNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetComputedStyleForNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetComputedStyleForNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetComputedStyleForNode(nil, &cdpcmd.CSSGetComputedStyleForNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetComputedStyleForNode(nil, &cdpcmd.CSSGetComputedStyleForNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetPlatformFontsForNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetPlatformFontsForNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetPlatformFontsForNode(nil, &cdpcmd.CSSGetPlatformFontsForNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetPlatformFontsForNode(nil, &cdpcmd.CSSGetPlatformFontsForNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetStyleSheetText(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetStyleSheetText(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetStyleSheetText(nil, &cdpcmd.CSSGetStyleSheetTextArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetStyleSheetText(nil, &cdpcmd.CSSGetStyleSheetTextArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_CollectClassNames(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.CollectClassNames(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CollectClassNames(nil, &cdpcmd.CSSCollectClassNamesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CollectClassNames(nil, &cdpcmd.CSSCollectClassNamesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetStyleSheetText(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.SetStyleSheetText(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetStyleSheetText(nil, &cdpcmd.CSSSetStyleSheetTextArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetStyleSheetText(nil, &cdpcmd.CSSSetStyleSheetTextArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetRuleSelector(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.SetRuleSelector(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetRuleSelector(nil, &cdpcmd.CSSSetRuleSelectorArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetRuleSelector(nil, &cdpcmd.CSSSetRuleSelectorArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetKeyframeKey(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.SetKeyframeKey(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetKeyframeKey(nil, &cdpcmd.CSSSetKeyframeKeyArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetKeyframeKey(nil, &cdpcmd.CSSSetKeyframeKeyArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetStyleTexts(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.SetStyleTexts(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetStyleTexts(nil, &cdpcmd.CSSSetStyleTextsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetStyleTexts(nil, &cdpcmd.CSSSetStyleTextsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetMediaText(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.SetMediaText(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetMediaText(nil, &cdpcmd.CSSSetMediaTextArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetMediaText(nil, &cdpcmd.CSSSetMediaTextArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_CreateStyleSheet(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.CreateStyleSheet(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CreateStyleSheet(nil, &cdpcmd.CSSCreateStyleSheetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CreateStyleSheet(nil, &cdpcmd.CSSCreateStyleSheetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_AddRule(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.AddRule(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.AddRule(nil, &cdpcmd.CSSAddRuleArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.AddRule(nil, &cdpcmd.CSSAddRuleArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_ForcePseudoState(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	err = dom.ForcePseudoState(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ForcePseudoState(nil, &cdpcmd.CSSForcePseudoStateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ForcePseudoState(nil, &cdpcmd.CSSForcePseudoStateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetMediaQueries(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	_, err = dom.GetMediaQueries(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetMediaQueries(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_SetEffectivePropertyValueForNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	err = dom.SetEffectivePropertyValueForNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetEffectivePropertyValueForNode(nil, &cdpcmd.CSSSetEffectivePropertyValueForNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetEffectivePropertyValueForNode(nil, &cdpcmd.CSSSetEffectivePropertyValueForNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetBackgroundColors(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetBackgroundColors(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetBackgroundColors(nil, &cdpcmd.CSSGetBackgroundColorsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetBackgroundColors(nil, &cdpcmd.CSSGetBackgroundColorsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_GetLayoutTreeAndStyles(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	// Test nil args.
	_, err = dom.GetLayoutTreeAndStyles(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetLayoutTreeAndStyles(nil, &cdpcmd.CSSGetLayoutTreeAndStylesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetLayoutTreeAndStyles(nil, &cdpcmd.CSSGetLayoutTreeAndStylesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_StartRuleUsageTracking(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	err = dom.StartRuleUsageTracking(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartRuleUsageTracking(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_TakeCoverageDelta(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	_, err = dom.TakeCoverageDelta(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.TakeCoverageDelta(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_StopRuleUsageTracking(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)
	var err error

	_, err = dom.StopRuleUsageTracking(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.StopRuleUsageTracking(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCSS_MediaQueryResultChanged(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)

	stream, err := dom.MediaQueryResultChanged(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.CSSMediaQueryResultChanged.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.MediaQueryResultChanged(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestCSS_FontsUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)

	stream, err := dom.FontsUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.CSSFontsUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FontsUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestCSS_StyleSheetChanged(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)

	stream, err := dom.StyleSheetChanged(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.CSSStyleSheetChanged.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.StyleSheetChanged(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestCSS_StyleSheetAdded(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)

	stream, err := dom.StyleSheetAdded(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.CSSStyleSheetAdded.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.StyleSheetAdded(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestCSS_StyleSheetRemoved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCSS(conn)

	stream, err := dom.StyleSheetRemoved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.CSSStyleSheetRemoved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.StyleSheetRemoved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestCacheStorage_RequestCacheNames(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCacheStorage(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestCacheNames(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestCacheNames(nil, &cdpcmd.CacheStorageRequestCacheNamesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestCacheNames(nil, &cdpcmd.CacheStorageRequestCacheNamesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCacheStorage_RequestEntries(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCacheStorage(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestEntries(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestEntries(nil, &cdpcmd.CacheStorageRequestEntriesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestEntries(nil, &cdpcmd.CacheStorageRequestEntriesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCacheStorage_DeleteCache(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCacheStorage(conn)
	var err error

	// Test nil args.
	err = dom.DeleteCache(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeleteCache(nil, &cdpcmd.CacheStorageDeleteCacheArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeleteCache(nil, &cdpcmd.CacheStorageDeleteCacheArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestCacheStorage_DeleteEntry(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewCacheStorage(conn)
	var err error

	// Test nil args.
	err = dom.DeleteEntry(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeleteEntry(nil, &cdpcmd.CacheStorageDeleteEntryArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeleteEntry(nil, &cdpcmd.CacheStorageDeleteEntryArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestConsole_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewConsole(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestConsole_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewConsole(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestConsole_ClearMessages(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewConsole(conn)
	var err error

	err = dom.ClearMessages(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearMessages(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestConsole_MessageAdded(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewConsole(conn)

	stream, err := dom.MessageAdded(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ConsoleMessageAdded.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.MessageAdded(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetDocument(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetDocument(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetDocument(nil, &cdpcmd.DOMGetDocumentArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetDocument(nil, &cdpcmd.DOMGetDocumentArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetFlattenedDocument(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetFlattenedDocument(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetFlattenedDocument(nil, &cdpcmd.DOMGetFlattenedDocumentArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetFlattenedDocument(nil, &cdpcmd.DOMGetFlattenedDocumentArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_CollectClassNamesFromSubtree(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.CollectClassNamesFromSubtree(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CollectClassNamesFromSubtree(nil, &cdpcmd.DOMCollectClassNamesFromSubtreeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CollectClassNamesFromSubtree(nil, &cdpcmd.DOMCollectClassNamesFromSubtreeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_RequestChildNodes(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.RequestChildNodes(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RequestChildNodes(nil, &cdpcmd.DOMRequestChildNodesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RequestChildNodes(nil, &cdpcmd.DOMRequestChildNodesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_QuerySelector(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.QuerySelector(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.QuerySelector(nil, &cdpcmd.DOMQuerySelectorArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.QuerySelector(nil, &cdpcmd.DOMQuerySelectorArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_QuerySelectorAll(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.QuerySelectorAll(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.QuerySelectorAll(nil, &cdpcmd.DOMQuerySelectorAllArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.QuerySelectorAll(nil, &cdpcmd.DOMQuerySelectorAllArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetNodeName(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.SetNodeName(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetNodeName(nil, &cdpcmd.DOMSetNodeNameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetNodeName(nil, &cdpcmd.DOMSetNodeNameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetNodeValue(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetNodeValue(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetNodeValue(nil, &cdpcmd.DOMSetNodeValueArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetNodeValue(nil, &cdpcmd.DOMSetNodeValueArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_RemoveNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.RemoveNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveNode(nil, &cdpcmd.DOMRemoveNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveNode(nil, &cdpcmd.DOMRemoveNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetAttributeValue(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetAttributeValue(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAttributeValue(nil, &cdpcmd.DOMSetAttributeValueArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAttributeValue(nil, &cdpcmd.DOMSetAttributeValueArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetAttributesAsText(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetAttributesAsText(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAttributesAsText(nil, &cdpcmd.DOMSetAttributesAsTextArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAttributesAsText(nil, &cdpcmd.DOMSetAttributesAsTextArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_RemoveAttribute(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.RemoveAttribute(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveAttribute(nil, &cdpcmd.DOMRemoveAttributeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveAttribute(nil, &cdpcmd.DOMRemoveAttributeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetOuterHTML(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetOuterHTML(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetOuterHTML(nil, &cdpcmd.DOMGetOuterHTMLArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetOuterHTML(nil, &cdpcmd.DOMGetOuterHTMLArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetOuterHTML(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetOuterHTML(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetOuterHTML(nil, &cdpcmd.DOMSetOuterHTMLArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetOuterHTML(nil, &cdpcmd.DOMSetOuterHTMLArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_PerformSearch(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.PerformSearch(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.PerformSearch(nil, &cdpcmd.DOMPerformSearchArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.PerformSearch(nil, &cdpcmd.DOMPerformSearchArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetSearchResults(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetSearchResults(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetSearchResults(nil, &cdpcmd.DOMGetSearchResultsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetSearchResults(nil, &cdpcmd.DOMGetSearchResultsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_DiscardSearchResults(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.DiscardSearchResults(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DiscardSearchResults(nil, &cdpcmd.DOMDiscardSearchResultsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DiscardSearchResults(nil, &cdpcmd.DOMDiscardSearchResultsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_RequestNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestNode(nil, &cdpcmd.DOMRequestNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestNode(nil, &cdpcmd.DOMRequestNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_HighlightRect(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.HighlightRect(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightRect(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_HighlightNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.HighlightNode(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightNode(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_HideHighlight(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.HideHighlight(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HideHighlight(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_PushNodeByPathToFrontend(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.PushNodeByPathToFrontend(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.PushNodeByPathToFrontend(nil, &cdpcmd.DOMPushNodeByPathToFrontendArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.PushNodeByPathToFrontend(nil, &cdpcmd.DOMPushNodeByPathToFrontendArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_PushNodesByBackendIdsToFrontend(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.PushNodesByBackendIdsToFrontend(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.PushNodesByBackendIdsToFrontend(nil, &cdpcmd.DOMPushNodesByBackendIdsToFrontendArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.PushNodesByBackendIdsToFrontend(nil, &cdpcmd.DOMPushNodesByBackendIdsToFrontendArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetInspectedNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetInspectedNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetInspectedNode(nil, &cdpcmd.DOMSetInspectedNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetInspectedNode(nil, &cdpcmd.DOMSetInspectedNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_ResolveNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.ResolveNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.ResolveNode(nil, &cdpcmd.DOMResolveNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.ResolveNode(nil, &cdpcmd.DOMResolveNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetAttributes(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetAttributes(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetAttributes(nil, &cdpcmd.DOMGetAttributesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetAttributes(nil, &cdpcmd.DOMGetAttributesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_CopyTo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.CopyTo(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CopyTo(nil, &cdpcmd.DOMCopyToArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CopyTo(nil, &cdpcmd.DOMCopyToArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_MoveTo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.MoveTo(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.MoveTo(nil, &cdpcmd.DOMMoveToArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.MoveTo(nil, &cdpcmd.DOMMoveToArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_Undo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.Undo(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Undo(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_Redo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.Redo(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Redo(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_MarkUndoableState(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	err = dom.MarkUndoableState(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.MarkUndoableState(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_Focus(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.Focus(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Focus(nil, &cdpcmd.DOMFocusArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Focus(nil, &cdpcmd.DOMFocusArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_SetFileInputFiles(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	err = dom.SetFileInputFiles(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetFileInputFiles(nil, &cdpcmd.DOMSetFileInputFilesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetFileInputFiles(nil, &cdpcmd.DOMSetFileInputFilesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetBoxModel(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetBoxModel(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetBoxModel(nil, &cdpcmd.DOMGetBoxModelArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetBoxModel(nil, &cdpcmd.DOMGetBoxModelArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetNodeForLocation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetNodeForLocation(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetNodeForLocation(nil, &cdpcmd.DOMGetNodeForLocationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetNodeForLocation(nil, &cdpcmd.DOMGetNodeForLocationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_GetRelayoutBoundary(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)
	var err error

	// Test nil args.
	_, err = dom.GetRelayoutBoundary(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetRelayoutBoundary(nil, &cdpcmd.DOMGetRelayoutBoundaryArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetRelayoutBoundary(nil, &cdpcmd.DOMGetRelayoutBoundaryArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOM_DocumentUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.DocumentUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMDocumentUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DocumentUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_SetChildNodes(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.SetChildNodes(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMSetChildNodes.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.SetChildNodes(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_AttributeModified(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.AttributeModified(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMAttributeModified.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AttributeModified(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_AttributeRemoved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.AttributeRemoved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMAttributeRemoved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AttributeRemoved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_InlineStyleInvalidated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.InlineStyleInvalidated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMInlineStyleInvalidated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.InlineStyleInvalidated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_CharacterDataModified(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.CharacterDataModified(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMCharacterDataModified.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.CharacterDataModified(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_ChildNodeCountUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.ChildNodeCountUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMChildNodeCountUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ChildNodeCountUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_ChildNodeInserted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.ChildNodeInserted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMChildNodeInserted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ChildNodeInserted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_ChildNodeRemoved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.ChildNodeRemoved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMChildNodeRemoved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ChildNodeRemoved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_ShadowRootPushed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.ShadowRootPushed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMShadowRootPushed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ShadowRootPushed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_ShadowRootPopped(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.ShadowRootPopped(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMShadowRootPopped.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ShadowRootPopped(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_PseudoElementAdded(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.PseudoElementAdded(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMPseudoElementAdded.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.PseudoElementAdded(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_PseudoElementRemoved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.PseudoElementRemoved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMPseudoElementRemoved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.PseudoElementRemoved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOM_DistributedNodesUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOM(conn)

	stream, err := dom.DistributedNodesUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMDistributedNodesUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DistributedNodesUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOMDebugger_SetDOMBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetDOMBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDOMBreakpoint(nil, &cdpcmd.DOMDebuggerSetDOMBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDOMBreakpoint(nil, &cdpcmd.DOMDebuggerSetDOMBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_RemoveDOMBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.RemoveDOMBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveDOMBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveDOMBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveDOMBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveDOMBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_SetEventListenerBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetEventListenerBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetEventListenerBreakpoint(nil, &cdpcmd.DOMDebuggerSetEventListenerBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetEventListenerBreakpoint(nil, &cdpcmd.DOMDebuggerSetEventListenerBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_RemoveEventListenerBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.RemoveEventListenerBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveEventListenerBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveEventListenerBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveEventListenerBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveEventListenerBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_SetInstrumentationBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetInstrumentationBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetInstrumentationBreakpoint(nil, &cdpcmd.DOMDebuggerSetInstrumentationBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetInstrumentationBreakpoint(nil, &cdpcmd.DOMDebuggerSetInstrumentationBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_RemoveInstrumentationBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.RemoveInstrumentationBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveInstrumentationBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveInstrumentationBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveInstrumentationBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveInstrumentationBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_SetXHRBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetXHRBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetXHRBreakpoint(nil, &cdpcmd.DOMDebuggerSetXHRBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetXHRBreakpoint(nil, &cdpcmd.DOMDebuggerSetXHRBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_RemoveXHRBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	err = dom.RemoveXHRBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveXHRBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveXHRBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveXHRBreakpoint(nil, &cdpcmd.DOMDebuggerRemoveXHRBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMDebugger_GetEventListeners(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.GetEventListeners(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetEventListeners(nil, &cdpcmd.DOMDebuggerGetEventListenersArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetEventListeners(nil, &cdpcmd.DOMDebuggerGetEventListenersArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_Clear(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	// Test nil args.
	err = dom.Clear(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Clear(nil, &cdpcmd.DOMStorageClearArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Clear(nil, &cdpcmd.DOMStorageClearArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_GetDOMStorageItems(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	// Test nil args.
	_, err = dom.GetDOMStorageItems(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetDOMStorageItems(nil, &cdpcmd.DOMStorageGetDOMStorageItemsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetDOMStorageItems(nil, &cdpcmd.DOMStorageGetDOMStorageItemsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_SetDOMStorageItem(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	// Test nil args.
	err = dom.SetDOMStorageItem(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDOMStorageItem(nil, &cdpcmd.DOMStorageSetDOMStorageItemArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDOMStorageItem(nil, &cdpcmd.DOMStorageSetDOMStorageItemArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_RemoveDOMStorageItem(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)
	var err error

	// Test nil args.
	err = dom.RemoveDOMStorageItem(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveDOMStorageItem(nil, &cdpcmd.DOMStorageRemoveDOMStorageItemArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveDOMStorageItem(nil, &cdpcmd.DOMStorageRemoveDOMStorageItemArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDOMStorage_DOMStorageItemsCleared(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)

	stream, err := dom.DOMStorageItemsCleared(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMStorageItemsCleared.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DOMStorageItemsCleared(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOMStorage_DOMStorageItemRemoved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)

	stream, err := dom.DOMStorageItemRemoved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMStorageItemRemoved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DOMStorageItemRemoved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOMStorage_DOMStorageItemAdded(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)

	stream, err := dom.DOMStorageItemAdded(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMStorageItemAdded.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DOMStorageItemAdded(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDOMStorage_DOMStorageItemUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDOMStorage(conn)

	stream, err := dom.DOMStorageItemUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DOMStorageItemUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DOMStorageItemUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDatabase_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDatabase(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDatabase_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDatabase(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDatabase_GetDatabaseTableNames(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDatabase(conn)
	var err error

	// Test nil args.
	_, err = dom.GetDatabaseTableNames(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetDatabaseTableNames(nil, &cdpcmd.DatabaseGetDatabaseTableNamesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetDatabaseTableNames(nil, &cdpcmd.DatabaseGetDatabaseTableNamesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDatabase_ExecuteSQL(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDatabase(conn)
	var err error

	// Test nil args.
	_, err = dom.ExecuteSQL(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.ExecuteSQL(nil, &cdpcmd.DatabaseExecuteSQLArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.ExecuteSQL(nil, &cdpcmd.DatabaseExecuteSQLArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDatabase_AddDatabase(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDatabase(conn)

	stream, err := dom.AddDatabase(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DatabaseAddDatabase.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AddDatabase(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDebugger_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetBreakpointsActive(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetBreakpointsActive(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetBreakpointsActive(nil, &cdpcmd.DebuggerSetBreakpointsActiveArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetBreakpointsActive(nil, &cdpcmd.DebuggerSetBreakpointsActiveArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetSkipAllPauses(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetSkipAllPauses(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetSkipAllPauses(nil, &cdpcmd.DebuggerSetSkipAllPausesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetSkipAllPauses(nil, &cdpcmd.DebuggerSetSkipAllPausesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetBreakpointByURL(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.SetBreakpointByURL(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetBreakpointByURL(nil, &cdpcmd.DebuggerSetBreakpointByURLArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetBreakpointByURL(nil, &cdpcmd.DebuggerSetBreakpointByURLArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.SetBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetBreakpoint(nil, &cdpcmd.DebuggerSetBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetBreakpoint(nil, &cdpcmd.DebuggerSetBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_RemoveBreakpoint(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.RemoveBreakpoint(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveBreakpoint(nil, &cdpcmd.DebuggerRemoveBreakpointArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveBreakpoint(nil, &cdpcmd.DebuggerRemoveBreakpointArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_GetPossibleBreakpoints(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.GetPossibleBreakpoints(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetPossibleBreakpoints(nil, &cdpcmd.DebuggerGetPossibleBreakpointsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetPossibleBreakpoints(nil, &cdpcmd.DebuggerGetPossibleBreakpointsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_ContinueToLocation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.ContinueToLocation(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ContinueToLocation(nil, &cdpcmd.DebuggerContinueToLocationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ContinueToLocation(nil, &cdpcmd.DebuggerContinueToLocationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_StepOver(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.StepOver(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StepOver(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_StepInto(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.StepInto(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StepInto(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_StepOut(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.StepOut(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StepOut(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_Pause(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.Pause(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Pause(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_ScheduleStepIntoAsync(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.ScheduleStepIntoAsync(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ScheduleStepIntoAsync(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_Resume(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	err = dom.Resume(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Resume(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SearchInContent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.SearchInContent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SearchInContent(nil, &cdpcmd.DebuggerSearchInContentArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SearchInContent(nil, &cdpcmd.DebuggerSearchInContentArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetScriptSource(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.SetScriptSource(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetScriptSource(nil, &cdpcmd.DebuggerSetScriptSourceArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetScriptSource(nil, &cdpcmd.DebuggerSetScriptSourceArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_RestartFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.RestartFrame(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RestartFrame(nil, &cdpcmd.DebuggerRestartFrameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RestartFrame(nil, &cdpcmd.DebuggerRestartFrameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_GetScriptSource(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.GetScriptSource(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetScriptSource(nil, &cdpcmd.DebuggerGetScriptSourceArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetScriptSource(nil, &cdpcmd.DebuggerGetScriptSourceArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetPauseOnExceptions(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetPauseOnExceptions(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPauseOnExceptions(nil, &cdpcmd.DebuggerSetPauseOnExceptionsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPauseOnExceptions(nil, &cdpcmd.DebuggerSetPauseOnExceptionsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_EvaluateOnCallFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	_, err = dom.EvaluateOnCallFrame(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.EvaluateOnCallFrame(nil, &cdpcmd.DebuggerEvaluateOnCallFrameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.EvaluateOnCallFrame(nil, &cdpcmd.DebuggerEvaluateOnCallFrameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetVariableValue(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetVariableValue(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetVariableValue(nil, &cdpcmd.DebuggerSetVariableValueArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetVariableValue(nil, &cdpcmd.DebuggerSetVariableValueArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetAsyncCallStackDepth(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetAsyncCallStackDepth(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAsyncCallStackDepth(nil, &cdpcmd.DebuggerSetAsyncCallStackDepthArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAsyncCallStackDepth(nil, &cdpcmd.DebuggerSetAsyncCallStackDepthArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetBlackboxPatterns(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetBlackboxPatterns(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetBlackboxPatterns(nil, &cdpcmd.DebuggerSetBlackboxPatternsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetBlackboxPatterns(nil, &cdpcmd.DebuggerSetBlackboxPatternsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_SetBlackboxedRanges(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)
	var err error

	// Test nil args.
	err = dom.SetBlackboxedRanges(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetBlackboxedRanges(nil, &cdpcmd.DebuggerSetBlackboxedRangesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetBlackboxedRanges(nil, &cdpcmd.DebuggerSetBlackboxedRangesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDebugger_ScriptParsed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)

	stream, err := dom.ScriptParsed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DebuggerScriptParsed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ScriptParsed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDebugger_ScriptFailedToParse(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)

	stream, err := dom.ScriptFailedToParse(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DebuggerScriptFailedToParse.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ScriptFailedToParse(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDebugger_BreakpointResolved(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)

	stream, err := dom.BreakpointResolved(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DebuggerBreakpointResolved.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.BreakpointResolved(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDebugger_Paused(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)

	stream, err := dom.Paused(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DebuggerPaused.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.Paused(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDebugger_Resumed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDebugger(conn)

	stream, err := dom.Resumed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.DebuggerResumed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.Resumed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestDeviceOrientation_SetDeviceOrientationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDeviceOrientation(conn)
	var err error

	// Test nil args.
	err = dom.SetDeviceOrientationOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDeviceOrientationOverride(nil, &cdpcmd.DeviceOrientationSetDeviceOrientationOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDeviceOrientationOverride(nil, &cdpcmd.DeviceOrientationSetDeviceOrientationOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestDeviceOrientation_ClearDeviceOrientationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewDeviceOrientation(conn)
	var err error

	err = dom.ClearDeviceOrientationOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearDeviceOrientationOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetDeviceMetricsOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetDeviceMetricsOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDeviceMetricsOverride(nil, &cdpcmd.EmulationSetDeviceMetricsOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDeviceMetricsOverride(nil, &cdpcmd.EmulationSetDeviceMetricsOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_ClearDeviceMetricsOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	err = dom.ClearDeviceMetricsOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearDeviceMetricsOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_ForceViewport(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.ForceViewport(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ForceViewport(nil, &cdpcmd.EmulationForceViewportArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ForceViewport(nil, &cdpcmd.EmulationForceViewportArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_ResetViewport(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	err = dom.ResetViewport(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ResetViewport(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_ResetPageScaleFactor(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	err = dom.ResetPageScaleFactor(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ResetPageScaleFactor(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetPageScaleFactor(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetPageScaleFactor(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPageScaleFactor(nil, &cdpcmd.EmulationSetPageScaleFactorArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPageScaleFactor(nil, &cdpcmd.EmulationSetPageScaleFactorArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetVisibleSize(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetVisibleSize(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetVisibleSize(nil, &cdpcmd.EmulationSetVisibleSizeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetVisibleSize(nil, &cdpcmd.EmulationSetVisibleSizeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetScriptExecutionDisabled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetScriptExecutionDisabled(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetScriptExecutionDisabled(nil, &cdpcmd.EmulationSetScriptExecutionDisabledArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetScriptExecutionDisabled(nil, &cdpcmd.EmulationSetScriptExecutionDisabledArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetGeolocationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetGeolocationOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetGeolocationOverride(nil, &cdpcmd.EmulationSetGeolocationOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetGeolocationOverride(nil, &cdpcmd.EmulationSetGeolocationOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_ClearGeolocationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	err = dom.ClearGeolocationOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearGeolocationOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetTouchEmulationEnabled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetTouchEmulationEnabled(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetTouchEmulationEnabled(nil, &cdpcmd.EmulationSetTouchEmulationEnabledArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetTouchEmulationEnabled(nil, &cdpcmd.EmulationSetTouchEmulationEnabledArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetEmulatedMedia(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetEmulatedMedia(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetEmulatedMedia(nil, &cdpcmd.EmulationSetEmulatedMediaArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetEmulatedMedia(nil, &cdpcmd.EmulationSetEmulatedMediaArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetCPUThrottlingRate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetCPUThrottlingRate(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetCPUThrottlingRate(nil, &cdpcmd.EmulationSetCPUThrottlingRateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetCPUThrottlingRate(nil, &cdpcmd.EmulationSetCPUThrottlingRateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_CanEmulate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	_, err = dom.CanEmulate(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CanEmulate(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetVirtualTimePolicy(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetVirtualTimePolicy(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetVirtualTimePolicy(nil, &cdpcmd.EmulationSetVirtualTimePolicyArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetVirtualTimePolicy(nil, &cdpcmd.EmulationSetVirtualTimePolicyArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_SetDefaultBackgroundColorOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)
	var err error

	// Test nil args.
	err = dom.SetDefaultBackgroundColorOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDefaultBackgroundColorOverride(nil, &cdpcmd.EmulationSetDefaultBackgroundColorOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDefaultBackgroundColorOverride(nil, &cdpcmd.EmulationSetDefaultBackgroundColorOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestEmulation_VirtualTimeBudgetExpired(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewEmulation(conn)

	stream, err := dom.VirtualTimeBudgetExpired(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.EmulationVirtualTimeBudgetExpired.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.VirtualTimeBudgetExpired(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestHeapProfiler_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_StartTrackingHeapObjects(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	err = dom.StartTrackingHeapObjects(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartTrackingHeapObjects(nil, &cdpcmd.HeapProfilerStartTrackingHeapObjectsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartTrackingHeapObjects(nil, &cdpcmd.HeapProfilerStartTrackingHeapObjectsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_StopTrackingHeapObjects(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	err = dom.StopTrackingHeapObjects(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StopTrackingHeapObjects(nil, &cdpcmd.HeapProfilerStopTrackingHeapObjectsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopTrackingHeapObjects(nil, &cdpcmd.HeapProfilerStopTrackingHeapObjectsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_TakeHeapSnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	err = dom.TakeHeapSnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.TakeHeapSnapshot(nil, &cdpcmd.HeapProfilerTakeHeapSnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.TakeHeapSnapshot(nil, &cdpcmd.HeapProfilerTakeHeapSnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_CollectGarbage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	err = dom.CollectGarbage(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.CollectGarbage(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_GetObjectByHeapObjectID(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	_, err = dom.GetObjectByHeapObjectID(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetObjectByHeapObjectID(nil, &cdpcmd.HeapProfilerGetObjectByHeapObjectIDArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetObjectByHeapObjectID(nil, &cdpcmd.HeapProfilerGetObjectByHeapObjectIDArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_AddInspectedHeapObject(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	err = dom.AddInspectedHeapObject(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.AddInspectedHeapObject(nil, &cdpcmd.HeapProfilerAddInspectedHeapObjectArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.AddInspectedHeapObject(nil, &cdpcmd.HeapProfilerAddInspectedHeapObjectArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_GetHeapObjectID(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	_, err = dom.GetHeapObjectID(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetHeapObjectID(nil, &cdpcmd.HeapProfilerGetHeapObjectIDArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetHeapObjectID(nil, &cdpcmd.HeapProfilerGetHeapObjectIDArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_StartSampling(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	// Test nil args.
	err = dom.StartSampling(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartSampling(nil, &cdpcmd.HeapProfilerStartSamplingArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartSampling(nil, &cdpcmd.HeapProfilerStartSamplingArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_StopSampling(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)
	var err error

	_, err = dom.StopSampling(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.StopSampling(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestHeapProfiler_AddHeapSnapshotChunk(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)

	stream, err := dom.AddHeapSnapshotChunk(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.HeapProfilerAddHeapSnapshotChunk.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AddHeapSnapshotChunk(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestHeapProfiler_ResetProfiles(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)

	stream, err := dom.ResetProfiles(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.HeapProfilerResetProfiles.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ResetProfiles(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestHeapProfiler_ReportHeapSnapshotProgress(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)

	stream, err := dom.ReportHeapSnapshotProgress(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.HeapProfilerReportHeapSnapshotProgress.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ReportHeapSnapshotProgress(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestHeapProfiler_LastSeenObjectID(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)

	stream, err := dom.LastSeenObjectID(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.HeapProfilerLastSeenObjectID.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LastSeenObjectID(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestHeapProfiler_HeapStatsUpdate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewHeapProfiler(conn)

	stream, err := dom.HeapStatsUpdate(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.HeapProfilerHeapStatsUpdate.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.HeapStatsUpdate(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestIO_Read(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIO(conn)
	var err error

	// Test nil args.
	_, err = dom.Read(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.Read(nil, &cdpcmd.IOReadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.Read(nil, &cdpcmd.IOReadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIO_Close(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIO(conn)
	var err error

	// Test nil args.
	err = dom.Close(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Close(nil, &cdpcmd.IOCloseArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Close(nil, &cdpcmd.IOCloseArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_RequestDatabaseNames(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestDatabaseNames(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestDatabaseNames(nil, &cdpcmd.IndexedDBRequestDatabaseNamesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestDatabaseNames(nil, &cdpcmd.IndexedDBRequestDatabaseNamesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_RequestDatabase(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestDatabase(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestDatabase(nil, &cdpcmd.IndexedDBRequestDatabaseArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestDatabase(nil, &cdpcmd.IndexedDBRequestDatabaseArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_RequestData(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	// Test nil args.
	_, err = dom.RequestData(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RequestData(nil, &cdpcmd.IndexedDBRequestDataArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestData(nil, &cdpcmd.IndexedDBRequestDataArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_ClearObjectStore(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	// Test nil args.
	err = dom.ClearObjectStore(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ClearObjectStore(nil, &cdpcmd.IndexedDBClearObjectStoreArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearObjectStore(nil, &cdpcmd.IndexedDBClearObjectStoreArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestIndexedDB_DeleteDatabase(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewIndexedDB(conn)
	var err error

	// Test nil args.
	err = dom.DeleteDatabase(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeleteDatabase(nil, &cdpcmd.IndexedDBDeleteDatabaseArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeleteDatabase(nil, &cdpcmd.IndexedDBDeleteDatabaseArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_SetIgnoreInputEvents(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.SetIgnoreInputEvents(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetIgnoreInputEvents(nil, &cdpcmd.InputSetIgnoreInputEventsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetIgnoreInputEvents(nil, &cdpcmd.InputSetIgnoreInputEventsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_DispatchKeyEvent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.DispatchKeyEvent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DispatchKeyEvent(nil, &cdpcmd.InputDispatchKeyEventArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DispatchKeyEvent(nil, &cdpcmd.InputDispatchKeyEventArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_DispatchMouseEvent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.DispatchMouseEvent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DispatchMouseEvent(nil, &cdpcmd.InputDispatchMouseEventArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DispatchMouseEvent(nil, &cdpcmd.InputDispatchMouseEventArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_DispatchTouchEvent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.DispatchTouchEvent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DispatchTouchEvent(nil, &cdpcmd.InputDispatchTouchEventArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DispatchTouchEvent(nil, &cdpcmd.InputDispatchTouchEventArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_EmulateTouchFromMouseEvent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.EmulateTouchFromMouseEvent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.EmulateTouchFromMouseEvent(nil, &cdpcmd.InputEmulateTouchFromMouseEventArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.EmulateTouchFromMouseEvent(nil, &cdpcmd.InputEmulateTouchFromMouseEventArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_SynthesizePinchGesture(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.SynthesizePinchGesture(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SynthesizePinchGesture(nil, &cdpcmd.InputSynthesizePinchGestureArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SynthesizePinchGesture(nil, &cdpcmd.InputSynthesizePinchGestureArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_SynthesizeScrollGesture(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.SynthesizeScrollGesture(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SynthesizeScrollGesture(nil, &cdpcmd.InputSynthesizeScrollGestureArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SynthesizeScrollGesture(nil, &cdpcmd.InputSynthesizeScrollGestureArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInput_SynthesizeTapGesture(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInput(conn)
	var err error

	// Test nil args.
	err = dom.SynthesizeTapGesture(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SynthesizeTapGesture(nil, &cdpcmd.InputSynthesizeTapGestureArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SynthesizeTapGesture(nil, &cdpcmd.InputSynthesizeTapGestureArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInspector_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInspector(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInspector_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInspector(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestInspector_Detached(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInspector(conn)

	stream, err := dom.Detached(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.InspectorDetached.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.Detached(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestInspector_TargetCrashed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewInspector(conn)

	stream, err := dom.TargetCrashed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.InspectorTargetCrashed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.TargetCrashed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestLayerTree_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_CompositingReasons(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.CompositingReasons(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CompositingReasons(nil, &cdpcmd.LayerTreeCompositingReasonsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CompositingReasons(nil, &cdpcmd.LayerTreeCompositingReasonsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_MakeSnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.MakeSnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.MakeSnapshot(nil, &cdpcmd.LayerTreeMakeSnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.MakeSnapshot(nil, &cdpcmd.LayerTreeMakeSnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_LoadSnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.LoadSnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.LoadSnapshot(nil, &cdpcmd.LayerTreeLoadSnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.LoadSnapshot(nil, &cdpcmd.LayerTreeLoadSnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_ReleaseSnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	err = dom.ReleaseSnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ReleaseSnapshot(nil, &cdpcmd.LayerTreeReleaseSnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ReleaseSnapshot(nil, &cdpcmd.LayerTreeReleaseSnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_ProfileSnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.ProfileSnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.ProfileSnapshot(nil, &cdpcmd.LayerTreeProfileSnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.ProfileSnapshot(nil, &cdpcmd.LayerTreeProfileSnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_ReplaySnapshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.ReplaySnapshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.ReplaySnapshot(nil, &cdpcmd.LayerTreeReplaySnapshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.ReplaySnapshot(nil, &cdpcmd.LayerTreeReplaySnapshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_SnapshotCommandLog(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)
	var err error

	// Test nil args.
	_, err = dom.SnapshotCommandLog(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SnapshotCommandLog(nil, &cdpcmd.LayerTreeSnapshotCommandLogArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SnapshotCommandLog(nil, &cdpcmd.LayerTreeSnapshotCommandLogArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLayerTree_LayerTreeDidChange(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)

	stream, err := dom.LayerTreeDidChange(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.LayerTreeDidChange.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LayerTreeDidChange(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestLayerTree_LayerPainted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLayerTree(conn)

	stream, err := dom.LayerPainted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.LayerTreeLayerPainted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LayerPainted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestLog_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLog_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLog_Clear(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)
	var err error

	err = dom.Clear(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Clear(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLog_StartViolationsReport(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)
	var err error

	// Test nil args.
	err = dom.StartViolationsReport(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartViolationsReport(nil, &cdpcmd.LogStartViolationsReportArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartViolationsReport(nil, &cdpcmd.LogStartViolationsReportArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLog_StopViolationsReport(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)
	var err error

	err = dom.StopViolationsReport(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopViolationsReport(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestLog_EntryAdded(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewLog(conn)

	stream, err := dom.EntryAdded(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.LogEntryAdded.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.EntryAdded(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestMemory_GetDOMCounters(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewMemory(conn)
	var err error

	_, err = dom.GetDOMCounters(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetDOMCounters(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestMemory_SetPressureNotificationsSuppressed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewMemory(conn)
	var err error

	// Test nil args.
	err = dom.SetPressureNotificationsSuppressed(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPressureNotificationsSuppressed(nil, &cdpcmd.MemorySetPressureNotificationsSuppressedArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPressureNotificationsSuppressed(nil, &cdpcmd.MemorySetPressureNotificationsSuppressedArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestMemory_SimulatePressureNotification(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewMemory(conn)
	var err error

	// Test nil args.
	err = dom.SimulatePressureNotification(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SimulatePressureNotification(nil, &cdpcmd.MemorySimulatePressureNotificationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SimulatePressureNotification(nil, &cdpcmd.MemorySimulatePressureNotificationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.Enable(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Enable(nil, &cdpcmd.NetworkEnableArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil, &cdpcmd.NetworkEnableArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetUserAgentOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetUserAgentOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetUserAgentOverride(nil, &cdpcmd.NetworkSetUserAgentOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetUserAgentOverride(nil, &cdpcmd.NetworkSetUserAgentOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetExtraHTTPHeaders(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetExtraHTTPHeaders(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetExtraHTTPHeaders(nil, &cdpcmd.NetworkSetExtraHTTPHeadersArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetExtraHTTPHeaders(nil, &cdpcmd.NetworkSetExtraHTTPHeadersArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_GetResponseBody(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	_, err = dom.GetResponseBody(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetResponseBody(nil, &cdpcmd.NetworkGetResponseBodyArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetResponseBody(nil, &cdpcmd.NetworkGetResponseBodyArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetBlockedURLs(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetBlockedURLs(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetBlockedURLs(nil, &cdpcmd.NetworkSetBlockedURLsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetBlockedURLs(nil, &cdpcmd.NetworkSetBlockedURLsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_ReplayXHR(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.ReplayXHR(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ReplayXHR(nil, &cdpcmd.NetworkReplayXHRArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ReplayXHR(nil, &cdpcmd.NetworkReplayXHRArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_CanClearBrowserCache(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	_, err = dom.CanClearBrowserCache(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CanClearBrowserCache(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_ClearBrowserCache(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	err = dom.ClearBrowserCache(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearBrowserCache(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_CanClearBrowserCookies(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	_, err = dom.CanClearBrowserCookies(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CanClearBrowserCookies(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_ClearBrowserCookies(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	err = dom.ClearBrowserCookies(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearBrowserCookies(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_GetCookies(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	_, err = dom.GetCookies(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetCookies(nil, &cdpcmd.NetworkGetCookiesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetCookies(nil, &cdpcmd.NetworkGetCookiesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_GetAllCookies(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	_, err = dom.GetAllCookies(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetAllCookies(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_DeleteCookie(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.DeleteCookie(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeleteCookie(nil, &cdpcmd.NetworkDeleteCookieArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeleteCookie(nil, &cdpcmd.NetworkDeleteCookieArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetCookie(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	_, err = dom.SetCookie(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SetCookie(nil, &cdpcmd.NetworkSetCookieArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SetCookie(nil, &cdpcmd.NetworkSetCookieArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_CanEmulateNetworkConditions(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	_, err = dom.CanEmulateNetworkConditions(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CanEmulateNetworkConditions(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_EmulateNetworkConditions(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.EmulateNetworkConditions(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.EmulateNetworkConditions(nil, &cdpcmd.NetworkEmulateNetworkConditionsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.EmulateNetworkConditions(nil, &cdpcmd.NetworkEmulateNetworkConditionsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetCacheDisabled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetCacheDisabled(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetCacheDisabled(nil, &cdpcmd.NetworkSetCacheDisabledArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetCacheDisabled(nil, &cdpcmd.NetworkSetCacheDisabledArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetBypassServiceWorker(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetBypassServiceWorker(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetBypassServiceWorker(nil, &cdpcmd.NetworkSetBypassServiceWorkerArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetBypassServiceWorker(nil, &cdpcmd.NetworkSetBypassServiceWorkerArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_SetDataSizeLimitsForTest(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.SetDataSizeLimitsForTest(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDataSizeLimitsForTest(nil, &cdpcmd.NetworkSetDataSizeLimitsForTestArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDataSizeLimitsForTest(nil, &cdpcmd.NetworkSetDataSizeLimitsForTestArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_GetCertificate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	_, err = dom.GetCertificate(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetCertificate(nil, &cdpcmd.NetworkGetCertificateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetCertificate(nil, &cdpcmd.NetworkGetCertificateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_EnableRequestInterception(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.EnableRequestInterception(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.EnableRequestInterception(nil, &cdpcmd.NetworkEnableRequestInterceptionArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.EnableRequestInterception(nil, &cdpcmd.NetworkEnableRequestInterceptionArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_ContinueInterceptedRequest(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)
	var err error

	// Test nil args.
	err = dom.ContinueInterceptedRequest(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ContinueInterceptedRequest(nil, &cdpcmd.NetworkContinueInterceptedRequestArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ContinueInterceptedRequest(nil, &cdpcmd.NetworkContinueInterceptedRequestArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestNetwork_ResourceChangedPriority(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.ResourceChangedPriority(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkResourceChangedPriority.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ResourceChangedPriority(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_RequestWillBeSent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.RequestWillBeSent(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkRequestWillBeSent.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.RequestWillBeSent(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_RequestServedFromCache(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.RequestServedFromCache(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkRequestServedFromCache.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.RequestServedFromCache(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_ResponseReceived(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.ResponseReceived(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkResponseReceived.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ResponseReceived(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_DataReceived(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.DataReceived(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkDataReceived.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DataReceived(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_LoadingFinished(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.LoadingFinished(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkLoadingFinished.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LoadingFinished(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_LoadingFailed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.LoadingFailed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkLoadingFailed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LoadingFailed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketWillSendHandshakeRequest(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketWillSendHandshakeRequest(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketWillSendHandshakeRequest.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketWillSendHandshakeRequest(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketHandshakeResponseReceived(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketHandshakeResponseReceived(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketHandshakeResponseReceived.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketHandshakeResponseReceived(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketCreated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketCreated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketCreated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketCreated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketClosed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketClosed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketClosed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketClosed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketFrameReceived(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketFrameReceived(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketFrameReceived.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketFrameReceived(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketFrameError(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketFrameError(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketFrameError.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketFrameError(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_WebSocketFrameSent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.WebSocketFrameSent(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkWebSocketFrameSent.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WebSocketFrameSent(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_EventSourceMessageReceived(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.EventSourceMessageReceived(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkEventSourceMessageReceived.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.EventSourceMessageReceived(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestNetwork_RequestIntercepted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewNetwork(conn)

	stream, err := dom.RequestIntercepted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.NetworkRequestIntercepted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.RequestIntercepted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestOverlay_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetShowPaintRects(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetShowPaintRects(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetShowPaintRects(nil, &cdpcmd.OverlaySetShowPaintRectsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetShowPaintRects(nil, &cdpcmd.OverlaySetShowPaintRectsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetShowDebugBorders(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetShowDebugBorders(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetShowDebugBorders(nil, &cdpcmd.OverlaySetShowDebugBordersArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetShowDebugBorders(nil, &cdpcmd.OverlaySetShowDebugBordersArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetShowFPSCounter(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetShowFPSCounter(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetShowFPSCounter(nil, &cdpcmd.OverlaySetShowFPSCounterArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetShowFPSCounter(nil, &cdpcmd.OverlaySetShowFPSCounterArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetShowScrollBottleneckRects(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetShowScrollBottleneckRects(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetShowScrollBottleneckRects(nil, &cdpcmd.OverlaySetShowScrollBottleneckRectsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetShowScrollBottleneckRects(nil, &cdpcmd.OverlaySetShowScrollBottleneckRectsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetShowViewportSizeOnResize(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetShowViewportSizeOnResize(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetShowViewportSizeOnResize(nil, &cdpcmd.OverlaySetShowViewportSizeOnResizeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetShowViewportSizeOnResize(nil, &cdpcmd.OverlaySetShowViewportSizeOnResizeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetPausedInDebuggerMessage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetPausedInDebuggerMessage(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetPausedInDebuggerMessage(nil, &cdpcmd.OverlaySetPausedInDebuggerMessageArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetPausedInDebuggerMessage(nil, &cdpcmd.OverlaySetPausedInDebuggerMessageArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetSuspended(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetSuspended(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetSuspended(nil, &cdpcmd.OverlaySetSuspendedArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetSuspended(nil, &cdpcmd.OverlaySetSuspendedArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_SetInspectMode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.SetInspectMode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetInspectMode(nil, &cdpcmd.OverlaySetInspectModeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetInspectMode(nil, &cdpcmd.OverlaySetInspectModeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_HighlightRect(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.HighlightRect(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HighlightRect(nil, &cdpcmd.OverlayHighlightRectArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightRect(nil, &cdpcmd.OverlayHighlightRectArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_HighlightQuad(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.HighlightQuad(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HighlightQuad(nil, &cdpcmd.OverlayHighlightQuadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightQuad(nil, &cdpcmd.OverlayHighlightQuadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_HighlightNode(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.HighlightNode(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HighlightNode(nil, &cdpcmd.OverlayHighlightNodeArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightNode(nil, &cdpcmd.OverlayHighlightNodeArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_HighlightFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	err = dom.HighlightFrame(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HighlightFrame(nil, &cdpcmd.OverlayHighlightFrameArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HighlightFrame(nil, &cdpcmd.OverlayHighlightFrameArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_HideHighlight(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	err = dom.HideHighlight(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HideHighlight(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_GetHighlightObjectForTest(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)
	var err error

	// Test nil args.
	_, err = dom.GetHighlightObjectForTest(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetHighlightObjectForTest(nil, &cdpcmd.OverlayGetHighlightObjectForTestArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetHighlightObjectForTest(nil, &cdpcmd.OverlayGetHighlightObjectForTestArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestOverlay_NodeHighlightRequested(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)

	stream, err := dom.NodeHighlightRequested(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.OverlayNodeHighlightRequested.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.NodeHighlightRequested(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestOverlay_InspectNodeRequested(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewOverlay(conn)

	stream, err := dom.InspectNodeRequested(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.OverlayInspectNodeRequested.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.InspectNodeRequested(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_AddScriptToEvaluateOnLoad(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.AddScriptToEvaluateOnLoad(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.AddScriptToEvaluateOnLoad(nil, &cdpcmd.PageAddScriptToEvaluateOnLoadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.AddScriptToEvaluateOnLoad(nil, &cdpcmd.PageAddScriptToEvaluateOnLoadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_RemoveScriptToEvaluateOnLoad(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.RemoveScriptToEvaluateOnLoad(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RemoveScriptToEvaluateOnLoad(nil, &cdpcmd.PageRemoveScriptToEvaluateOnLoadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RemoveScriptToEvaluateOnLoad(nil, &cdpcmd.PageRemoveScriptToEvaluateOnLoadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetAutoAttachToCreatedPages(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetAutoAttachToCreatedPages(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAutoAttachToCreatedPages(nil, &cdpcmd.PageSetAutoAttachToCreatedPagesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAutoAttachToCreatedPages(nil, &cdpcmd.PageSetAutoAttachToCreatedPagesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_Reload(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.Reload(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Reload(nil, &cdpcmd.PageReloadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Reload(nil, &cdpcmd.PageReloadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_Navigate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.Navigate(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.Navigate(nil, &cdpcmd.PageNavigateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.Navigate(nil, &cdpcmd.PageNavigateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_StopLoading(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.StopLoading(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopLoading(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetNavigationHistory(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	_, err = dom.GetNavigationHistory(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetNavigationHistory(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_NavigateToHistoryEntry(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.NavigateToHistoryEntry(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.NavigateToHistoryEntry(nil, &cdpcmd.PageNavigateToHistoryEntryArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.NavigateToHistoryEntry(nil, &cdpcmd.PageNavigateToHistoryEntryArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetCookies(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	_, err = dom.GetCookies(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetCookies(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_DeleteCookie(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.DeleteCookie(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeleteCookie(nil, &cdpcmd.PageDeleteCookieArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeleteCookie(nil, &cdpcmd.PageDeleteCookieArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetResourceTree(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	_, err = dom.GetResourceTree(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetResourceTree(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetResourceContent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.GetResourceContent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetResourceContent(nil, &cdpcmd.PageGetResourceContentArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetResourceContent(nil, &cdpcmd.PageGetResourceContentArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SearchInResource(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.SearchInResource(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.SearchInResource(nil, &cdpcmd.PageSearchInResourceArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.SearchInResource(nil, &cdpcmd.PageSearchInResourceArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetDocumentContent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetDocumentContent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDocumentContent(nil, &cdpcmd.PageSetDocumentContentArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDocumentContent(nil, &cdpcmd.PageSetDocumentContentArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetDeviceMetricsOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetDeviceMetricsOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDeviceMetricsOverride(nil, &cdpcmd.PageSetDeviceMetricsOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDeviceMetricsOverride(nil, &cdpcmd.PageSetDeviceMetricsOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_ClearDeviceMetricsOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.ClearDeviceMetricsOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearDeviceMetricsOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetGeolocationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetGeolocationOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetGeolocationOverride(nil, &cdpcmd.PageSetGeolocationOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetGeolocationOverride(nil, &cdpcmd.PageSetGeolocationOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_ClearGeolocationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.ClearGeolocationOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearGeolocationOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetDeviceOrientationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetDeviceOrientationOverride(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDeviceOrientationOverride(nil, &cdpcmd.PageSetDeviceOrientationOverrideArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDeviceOrientationOverride(nil, &cdpcmd.PageSetDeviceOrientationOverrideArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_ClearDeviceOrientationOverride(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.ClearDeviceOrientationOverride(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearDeviceOrientationOverride(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetTouchEmulationEnabled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetTouchEmulationEnabled(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetTouchEmulationEnabled(nil, &cdpcmd.PageSetTouchEmulationEnabledArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetTouchEmulationEnabled(nil, &cdpcmd.PageSetTouchEmulationEnabledArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_CaptureScreenshot(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.CaptureScreenshot(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CaptureScreenshot(nil, &cdpcmd.PageCaptureScreenshotArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CaptureScreenshot(nil, &cdpcmd.PageCaptureScreenshotArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_PrintToPDF(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	_, err = dom.PrintToPDF(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.PrintToPDF(nil, &cdpcmd.PagePrintToPDFArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.PrintToPDF(nil, &cdpcmd.PagePrintToPDFArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_StartScreencast(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.StartScreencast(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartScreencast(nil, &cdpcmd.PageStartScreencastArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartScreencast(nil, &cdpcmd.PageStartScreencastArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_StopScreencast(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.StopScreencast(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopScreencast(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_ScreencastFrameAck(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.ScreencastFrameAck(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ScreencastFrameAck(nil, &cdpcmd.PageScreencastFrameAckArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ScreencastFrameAck(nil, &cdpcmd.PageScreencastFrameAckArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_HandleJavaScriptDialog(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.HandleJavaScriptDialog(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HandleJavaScriptDialog(nil, &cdpcmd.PageHandleJavaScriptDialogArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HandleJavaScriptDialog(nil, &cdpcmd.PageHandleJavaScriptDialogArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetAppManifest(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	_, err = dom.GetAppManifest(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetAppManifest(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_RequestAppBanner(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	err = dom.RequestAppBanner(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RequestAppBanner(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_SetControlNavigations(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.SetControlNavigations(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetControlNavigations(nil, &cdpcmd.PageSetControlNavigationsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetControlNavigations(nil, &cdpcmd.PageSetControlNavigationsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_ProcessNavigation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.ProcessNavigation(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ProcessNavigation(nil, &cdpcmd.PageProcessNavigationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ProcessNavigation(nil, &cdpcmd.PageProcessNavigationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_GetLayoutMetrics(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	_, err = dom.GetLayoutMetrics(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetLayoutMetrics(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_CreateIsolatedWorld(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)
	var err error

	// Test nil args.
	err = dom.CreateIsolatedWorld(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.CreateIsolatedWorld(nil, &cdpcmd.PageCreateIsolatedWorldArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.CreateIsolatedWorld(nil, &cdpcmd.PageCreateIsolatedWorldArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestPage_DOMContentEventFired(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.DOMContentEventFired(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageDOMContentEventFired.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DOMContentEventFired(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_LoadEventFired(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.LoadEventFired(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageLoadEventFired.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.LoadEventFired(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameAttached(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameAttached(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameAttached.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameAttached(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameNavigated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameNavigated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameNavigated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameNavigated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameDetached(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameDetached(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameDetached.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameDetached(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameStartedLoading(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameStartedLoading(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameStartedLoading.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameStartedLoading(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameStoppedLoading(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameStoppedLoading(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameStoppedLoading.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameStoppedLoading(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameScheduledNavigation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameScheduledNavigation(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameScheduledNavigation.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameScheduledNavigation(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameClearedScheduledNavigation(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameClearedScheduledNavigation(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameClearedScheduledNavigation.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameClearedScheduledNavigation(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_FrameResized(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.FrameResized(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageFrameResized.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.FrameResized(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_JavascriptDialogOpening(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.JavascriptDialogOpening(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageJavascriptDialogOpening.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.JavascriptDialogOpening(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_JavascriptDialogClosed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.JavascriptDialogClosed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageJavascriptDialogClosed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.JavascriptDialogClosed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_ScreencastFrame(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.ScreencastFrame(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageScreencastFrame.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ScreencastFrame(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_ScreencastVisibilityChanged(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.ScreencastVisibilityChanged(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageScreencastVisibilityChanged.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ScreencastVisibilityChanged(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_InterstitialShown(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.InterstitialShown(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageInterstitialShown.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.InterstitialShown(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_InterstitialHidden(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.InterstitialHidden(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageInterstitialHidden.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.InterstitialHidden(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestPage_NavigationRequested(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewPage(conn)

	stream, err := dom.NavigationRequested(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.PageNavigationRequested.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.NavigationRequested(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestProfiler_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_SetSamplingInterval(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	// Test nil args.
	err = dom.SetSamplingInterval(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetSamplingInterval(nil, &cdpcmd.ProfilerSetSamplingIntervalArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetSamplingInterval(nil, &cdpcmd.ProfilerSetSamplingIntervalArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_Start(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	err = dom.Start(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Start(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_Stop(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	_, err = dom.Stop(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.Stop(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_StartPreciseCoverage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	// Test nil args.
	err = dom.StartPreciseCoverage(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartPreciseCoverage(nil, &cdpcmd.ProfilerStartPreciseCoverageArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartPreciseCoverage(nil, &cdpcmd.ProfilerStartPreciseCoverageArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_StopPreciseCoverage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	err = dom.StopPreciseCoverage(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopPreciseCoverage(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_TakePreciseCoverage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	_, err = dom.TakePreciseCoverage(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.TakePreciseCoverage(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_GetBestEffortCoverage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)
	var err error

	_, err = dom.GetBestEffortCoverage(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetBestEffortCoverage(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestProfiler_ConsoleProfileStarted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)

	stream, err := dom.ConsoleProfileStarted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ProfilerConsoleProfileStarted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ConsoleProfileStarted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestProfiler_ConsoleProfileFinished(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewProfiler(conn)

	stream, err := dom.ConsoleProfileFinished(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ProfilerConsoleProfileFinished.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ConsoleProfileFinished(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_Evaluate(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.Evaluate(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.Evaluate(nil, &cdpcmd.RuntimeEvaluateArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.Evaluate(nil, &cdpcmd.RuntimeEvaluateArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_AwaitPromise(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.AwaitPromise(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.AwaitPromise(nil, &cdpcmd.RuntimeAwaitPromiseArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.AwaitPromise(nil, &cdpcmd.RuntimeAwaitPromiseArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_CallFunctionOn(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.CallFunctionOn(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CallFunctionOn(nil, &cdpcmd.RuntimeCallFunctionOnArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CallFunctionOn(nil, &cdpcmd.RuntimeCallFunctionOnArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_GetProperties(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.GetProperties(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetProperties(nil, &cdpcmd.RuntimeGetPropertiesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetProperties(nil, &cdpcmd.RuntimeGetPropertiesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_ReleaseObject(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	err = dom.ReleaseObject(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ReleaseObject(nil, &cdpcmd.RuntimeReleaseObjectArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ReleaseObject(nil, &cdpcmd.RuntimeReleaseObjectArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_ReleaseObjectGroup(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	err = dom.ReleaseObjectGroup(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ReleaseObjectGroup(nil, &cdpcmd.RuntimeReleaseObjectGroupArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ReleaseObjectGroup(nil, &cdpcmd.RuntimeReleaseObjectGroupArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_RunIfWaitingForDebugger(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	err = dom.RunIfWaitingForDebugger(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RunIfWaitingForDebugger(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_DiscardConsoleEntries(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	err = dom.DiscardConsoleEntries(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DiscardConsoleEntries(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_SetCustomObjectFormatterEnabled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	err = dom.SetCustomObjectFormatterEnabled(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetCustomObjectFormatterEnabled(nil, &cdpcmd.RuntimeSetCustomObjectFormatterEnabledArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetCustomObjectFormatterEnabled(nil, &cdpcmd.RuntimeSetCustomObjectFormatterEnabledArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_CompileScript(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.CompileScript(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CompileScript(nil, &cdpcmd.RuntimeCompileScriptArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CompileScript(nil, &cdpcmd.RuntimeCompileScriptArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_RunScript(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)
	var err error

	// Test nil args.
	_, err = dom.RunScript(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.RunScript(nil, &cdpcmd.RuntimeRunScriptArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RunScript(nil, &cdpcmd.RuntimeRunScriptArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestRuntime_ExecutionContextCreated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ExecutionContextCreated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeExecutionContextCreated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ExecutionContextCreated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_ExecutionContextDestroyed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ExecutionContextDestroyed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeExecutionContextDestroyed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ExecutionContextDestroyed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_ExecutionContextsCleared(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ExecutionContextsCleared(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeExecutionContextsCleared.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ExecutionContextsCleared(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_ExceptionThrown(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ExceptionThrown(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeExceptionThrown.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ExceptionThrown(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_ExceptionRevoked(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ExceptionRevoked(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeExceptionRevoked.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ExceptionRevoked(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_ConsoleAPICalled(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.ConsoleAPICalled(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeConsoleAPICalled.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ConsoleAPICalled(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestRuntime_InspectRequested(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewRuntime(conn)

	stream, err := dom.InspectRequested(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.RuntimeInspectRequested.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.InspectRequested(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestSchema_GetDomains(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSchema(conn)
	var err error

	_, err = dom.GetDomains(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetDomains(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_ShowCertificateViewer(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)
	var err error

	err = dom.ShowCertificateViewer(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ShowCertificateViewer(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_HandleCertificateError(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)
	var err error

	// Test nil args.
	err = dom.HandleCertificateError(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.HandleCertificateError(nil, &cdpcmd.SecurityHandleCertificateErrorArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.HandleCertificateError(nil, &cdpcmd.SecurityHandleCertificateErrorArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_SetOverrideCertificateErrors(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)
	var err error

	// Test nil args.
	err = dom.SetOverrideCertificateErrors(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetOverrideCertificateErrors(nil, &cdpcmd.SecuritySetOverrideCertificateErrorsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetOverrideCertificateErrors(nil, &cdpcmd.SecuritySetOverrideCertificateErrorsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSecurity_SecurityStateChanged(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)

	stream, err := dom.SecurityStateChanged(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.SecurityStateChanged.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.SecurityStateChanged(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestSecurity_CertificateError(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSecurity(conn)

	stream, err := dom.CertificateError(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.SecurityCertificateError.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.CertificateError(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestServiceWorker_Enable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	err = dom.Enable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Enable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_Disable(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	err = dom.Disable(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Disable(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_Unregister(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.Unregister(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Unregister(nil, &cdpcmd.ServiceWorkerUnregisterArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Unregister(nil, &cdpcmd.ServiceWorkerUnregisterArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_UpdateRegistration(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.UpdateRegistration(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.UpdateRegistration(nil, &cdpcmd.ServiceWorkerUpdateRegistrationArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.UpdateRegistration(nil, &cdpcmd.ServiceWorkerUpdateRegistrationArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_StartWorker(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.StartWorker(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StartWorker(nil, &cdpcmd.ServiceWorkerStartWorkerArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StartWorker(nil, &cdpcmd.ServiceWorkerStartWorkerArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_SkipWaiting(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.SkipWaiting(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SkipWaiting(nil, &cdpcmd.ServiceWorkerSkipWaitingArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SkipWaiting(nil, &cdpcmd.ServiceWorkerSkipWaitingArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_StopWorker(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.StopWorker(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.StopWorker(nil, &cdpcmd.ServiceWorkerStopWorkerArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.StopWorker(nil, &cdpcmd.ServiceWorkerStopWorkerArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_InspectWorker(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.InspectWorker(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.InspectWorker(nil, &cdpcmd.ServiceWorkerInspectWorkerArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.InspectWorker(nil, &cdpcmd.ServiceWorkerInspectWorkerArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_SetForceUpdateOnPageLoad(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.SetForceUpdateOnPageLoad(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetForceUpdateOnPageLoad(nil, &cdpcmd.ServiceWorkerSetForceUpdateOnPageLoadArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetForceUpdateOnPageLoad(nil, &cdpcmd.ServiceWorkerSetForceUpdateOnPageLoadArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_DeliverPushMessage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.DeliverPushMessage(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DeliverPushMessage(nil, &cdpcmd.ServiceWorkerDeliverPushMessageArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DeliverPushMessage(nil, &cdpcmd.ServiceWorkerDeliverPushMessageArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_DispatchSyncEvent(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)
	var err error

	// Test nil args.
	err = dom.DispatchSyncEvent(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DispatchSyncEvent(nil, &cdpcmd.ServiceWorkerDispatchSyncEventArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DispatchSyncEvent(nil, &cdpcmd.ServiceWorkerDispatchSyncEventArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestServiceWorker_WorkerRegistrationUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)

	stream, err := dom.WorkerRegistrationUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ServiceWorkerWorkerRegistrationUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WorkerRegistrationUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestServiceWorker_WorkerVersionUpdated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)

	stream, err := dom.WorkerVersionUpdated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ServiceWorkerWorkerVersionUpdated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WorkerVersionUpdated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestServiceWorker_WorkerErrorReported(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewServiceWorker(conn)

	stream, err := dom.WorkerErrorReported(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.ServiceWorkerWorkerErrorReported.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.WorkerErrorReported(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestStorage_ClearDataForOrigin(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewStorage(conn)
	var err error

	// Test nil args.
	err = dom.ClearDataForOrigin(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ClearDataForOrigin(nil, &cdpcmd.StorageClearDataForOriginArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ClearDataForOrigin(nil, &cdpcmd.StorageClearDataForOriginArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestSystemInfo_GetInfo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewSystemInfo(conn)
	var err error

	_, err = dom.GetInfo(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetInfo(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_SetDiscoverTargets(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.SetDiscoverTargets(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetDiscoverTargets(nil, &cdpcmd.TargetSetDiscoverTargetsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetDiscoverTargets(nil, &cdpcmd.TargetSetDiscoverTargetsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_SetAutoAttach(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.SetAutoAttach(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAutoAttach(nil, &cdpcmd.TargetSetAutoAttachArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAutoAttach(nil, &cdpcmd.TargetSetAutoAttachArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_SetAttachToFrames(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.SetAttachToFrames(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetAttachToFrames(nil, &cdpcmd.TargetSetAttachToFramesArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetAttachToFrames(nil, &cdpcmd.TargetSetAttachToFramesArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_SetRemoteLocations(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.SetRemoteLocations(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SetRemoteLocations(nil, &cdpcmd.TargetSetRemoteLocationsArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SetRemoteLocations(nil, &cdpcmd.TargetSetRemoteLocationsArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_SendMessageToTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.SendMessageToTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.SendMessageToTarget(nil, &cdpcmd.TargetSendMessageToTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.SendMessageToTarget(nil, &cdpcmd.TargetSendMessageToTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_GetTargetInfo(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	_, err = dom.GetTargetInfo(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.GetTargetInfo(nil, &cdpcmd.TargetGetTargetInfoArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetTargetInfo(nil, &cdpcmd.TargetGetTargetInfoArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_ActivateTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.ActivateTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.ActivateTarget(nil, &cdpcmd.TargetActivateTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.ActivateTarget(nil, &cdpcmd.TargetActivateTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_CloseTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	_, err = dom.CloseTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CloseTarget(nil, &cdpcmd.TargetCloseTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CloseTarget(nil, &cdpcmd.TargetCloseTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_AttachToTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	_, err = dom.AttachToTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.AttachToTarget(nil, &cdpcmd.TargetAttachToTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.AttachToTarget(nil, &cdpcmd.TargetAttachToTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_DetachFromTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	err = dom.DetachFromTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.DetachFromTarget(nil, &cdpcmd.TargetDetachFromTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.DetachFromTarget(nil, &cdpcmd.TargetDetachFromTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_CreateBrowserContext(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	_, err = dom.CreateBrowserContext(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CreateBrowserContext(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_DisposeBrowserContext(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	_, err = dom.DisposeBrowserContext(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.DisposeBrowserContext(nil, &cdpcmd.TargetDisposeBrowserContextArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.DisposeBrowserContext(nil, &cdpcmd.TargetDisposeBrowserContextArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_CreateTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	// Test nil args.
	_, err = dom.CreateTarget(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	_, err = dom.CreateTarget(nil, &cdpcmd.TargetCreateTargetArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.CreateTarget(nil, &cdpcmd.TargetCreateTargetArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_GetTargets(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)
	var err error

	_, err = dom.GetTargets(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetTargets(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTarget_TargetCreated(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)

	stream, err := dom.TargetCreated(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TargetCreated.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.TargetCreated(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTarget_TargetDestroyed(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)

	stream, err := dom.TargetDestroyed(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TargetDestroyed.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.TargetDestroyed(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTarget_AttachedToTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)

	stream, err := dom.AttachedToTarget(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TargetAttachedToTarget.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.AttachedToTarget(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTarget_DetachedFromTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)

	stream, err := dom.DetachedFromTarget(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TargetDetachedFromTarget.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DetachedFromTarget(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTarget_ReceivedMessageFromTarget(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTarget(conn)

	stream, err := dom.ReceivedMessageFromTarget(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TargetReceivedMessageFromTarget.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.ReceivedMessageFromTarget(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTethering_Bind(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTethering(conn)
	var err error

	// Test nil args.
	err = dom.Bind(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Bind(nil, &cdpcmd.TetheringBindArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Bind(nil, &cdpcmd.TetheringBindArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTethering_Unbind(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTethering(conn)
	var err error

	// Test nil args.
	err = dom.Unbind(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Unbind(nil, &cdpcmd.TetheringUnbindArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Unbind(nil, &cdpcmd.TetheringUnbindArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTethering_Accepted(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTethering(conn)

	stream, err := dom.Accepted(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TetheringAccepted.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.Accepted(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTracing_Start(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)
	var err error

	// Test nil args.
	err = dom.Start(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.Start(nil, &cdpcmd.TracingStartArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.Start(nil, &cdpcmd.TracingStartArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTracing_End(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)
	var err error

	err = dom.End(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.End(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTracing_GetCategories(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)
	var err error

	_, err = dom.GetCategories(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.GetCategories(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTracing_RequestMemoryDump(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)
	var err error

	_, err = dom.RequestMemoryDump(nil)
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	_, err = dom.RequestMemoryDump(nil)
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTracing_RecordClockSyncMarker(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)
	var err error

	// Test nil args.
	err = dom.RecordClockSyncMarker(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	err = dom.RecordClockSyncMarker(nil, &cdpcmd.TracingRecordClockSyncMarkerArgs{})
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")
	err = dom.RecordClockSyncMarker(nil, &cdpcmd.TracingRecordClockSyncMarkerArgs{})
	if err == nil || err.(*opError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %v, want bad request", err)
	}
}

func TestTracing_DataCollected(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)

	stream, err := dom.DataCollected(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TracingDataCollected.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.DataCollected(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTracing_TracingComplete(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)

	stream, err := dom.TracingComplete(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TracingComplete.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.TracingComplete(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}

func TestTracing_BufferUsage(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := NewTracing(conn)

	stream, err := dom.BufferUsage(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = cdpevent.TracingBufferUsage.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*opError); !ok {
		t.Errorf("Recv() got %v, want opError", err)
	}

	conn.Close()
	stream, err = dom.BufferUsage(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}

}
