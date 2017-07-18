package devtool

import (
	"context"
	"errors"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// headlessCreateURL tries to create a new target for Headless Chrome that does
// not support the json new endpoint. A rpcc connection is established to the
// "/devtools/browser" endpoint and "Target.createTarget" is issued.
func headlessCreateURL(ctx context.Context, d *DevTools, openURL string) (*Target, error) {
	// Context must be set, rpcc DialContext panics on nil context.
	if ctx == nil {
		ctx = context.Background()
	}

	// Headless Chrome requires a non-empty URL for CreateTarget.
	if openURL == "" {
		openURL = "about:blank"
	}

	wsURL := "ws://" + httpRe.ReplaceAllString(d.url, "") + "/devtools/browser"
	conn, err := rpcc.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := cdp.NewClient(conn)
	t, err := c.Target.CreateTarget(ctx, target.NewCreateTargetArgs(openURL))
	if err != nil {
		return nil, err
	}

	// List must be called after CreateTarget (headless bug):
	// https://bugs.chromium.org/p/chromium/issues/detail?id=704503
	list, err := d.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, tt := range list {
		if tt.ID == string(t.TargetID) {
			return tt, nil
		}
	}

	return nil, errors.New("devtool: headlessCreateURL: could not create target")
}
