// Code generated by cdpgen. DO NOT EDIT.

// Package input implements the Input domain.
package input

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Input domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Input domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// DispatchDragEvent invokes the Input method. Dispatches a drag event into
// the page.
func (d *domainClient) DispatchDragEvent(ctx context.Context, args *DispatchDragEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchDragEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchDragEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchDragEvent", Err: err}
	}
	return
}

// DispatchKeyEvent invokes the Input method. Dispatches a key event to the
// page.
func (d *domainClient) DispatchKeyEvent(ctx context.Context, args *DispatchKeyEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchKeyEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchKeyEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchKeyEvent", Err: err}
	}
	return
}

// InsertText invokes the Input method. This method emulates inserting text
// that doesn't come from a key press, for example an emoji keyboard or an IME.
func (d *domainClient) InsertText(ctx context.Context, args *InsertTextArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.insertText", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.insertText", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "InsertText", Err: err}
	}
	return
}

// ImeSetComposition invokes the Input method. This method sets the current
// candidate text for ime. Use imeCommitComposition to commit the final text.
// Use imeSetComposition with empty string as text to cancel composition.
func (d *domainClient) ImeSetComposition(ctx context.Context, args *ImeSetCompositionArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.imeSetComposition", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.imeSetComposition", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "ImeSetComposition", Err: err}
	}
	return
}

// DispatchMouseEvent invokes the Input method. Dispatches a mouse event to
// the page.
func (d *domainClient) DispatchMouseEvent(ctx context.Context, args *DispatchMouseEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchMouseEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchMouseEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchMouseEvent", Err: err}
	}
	return
}

// DispatchTouchEvent invokes the Input method. Dispatches a touch event to
// the page.
func (d *domainClient) DispatchTouchEvent(ctx context.Context, args *DispatchTouchEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchTouchEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchTouchEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchTouchEvent", Err: err}
	}
	return
}

// EmulateTouchFromMouseEvent invokes the Input method. Emulates touch event
// from the mouse event parameters.
func (d *domainClient) EmulateTouchFromMouseEvent(ctx context.Context, args *EmulateTouchFromMouseEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.emulateTouchFromMouseEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.emulateTouchFromMouseEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "EmulateTouchFromMouseEvent", Err: err}
	}
	return
}

// SetIgnoreInputEvents invokes the Input method. Ignores input events (useful
// while auditing page).
func (d *domainClient) SetIgnoreInputEvents(ctx context.Context, args *SetIgnoreInputEventsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.setIgnoreInputEvents", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.setIgnoreInputEvents", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SetIgnoreInputEvents", Err: err}
	}
	return
}

// SetInterceptDrags invokes the Input method. Prevents default drag and drop
// behavior and instead emits `Input.dragIntercepted` events. Drag and drop
// behavior can be directly controlled via `Input.dispatchDragEvent`.
func (d *domainClient) SetInterceptDrags(ctx context.Context, args *SetInterceptDragsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.setInterceptDrags", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.setInterceptDrags", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SetInterceptDrags", Err: err}
	}
	return
}

// SynthesizePinchGesture invokes the Input method. Synthesizes a pinch
// gesture over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizePinchGesture(ctx context.Context, args *SynthesizePinchGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizePinchGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizePinchGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizePinchGesture", Err: err}
	}
	return
}

// SynthesizeScrollGesture invokes the Input method. Synthesizes a scroll
// gesture over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizeScrollGesture(ctx context.Context, args *SynthesizeScrollGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizeScrollGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizeScrollGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizeScrollGesture", Err: err}
	}
	return
}

// SynthesizeTapGesture invokes the Input method. Synthesizes a tap gesture
// over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizeTapGesture(ctx context.Context, args *SynthesizeTapGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizeTapGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizeTapGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizeTapGesture", Err: err}
	}
	return
}

func (d *domainClient) DragIntercepted(ctx context.Context) (DragInterceptedClient, error) {
	s, err := rpcc.NewStream(ctx, "Input.dragIntercepted", d.conn)
	if err != nil {
		return nil, err
	}
	return &dragInterceptedClient{Stream: s}, nil
}

type dragInterceptedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *dragInterceptedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *dragInterceptedClient) Recv() (*DragInterceptedReply, error) {
	event := new(DragInterceptedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Input", Op: "DragIntercepted Recv", Err: err}
	}
	return event, nil
}
