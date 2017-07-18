package cdp_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use the DevTools json API to get the current page.
	devt := devtool.New("http://127.0.0.1:9222")
	pageTarget, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pageTarget, err = devt.Create(ctx)
		if err != nil {
			panic(err)
		}
	}

	// Connect to Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pageTarget.WebSocketDebuggerURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Must be closed when we are done.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	// Enable events on the Page domain.
	if err = c.Page.Enable(ctx); err != nil {
		panic(err)
	}

	// New DOMContentEventFired client will receive and buffer
	// ContentEventFired events from now on.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		panic(err)
	}
	defer domContentEventFired.Close()

	// Create the Navigate arguments with the optional Referrer field set.
	navArgs := page.NewNavigateArgs("https://www.google.com").SetReferrer("https://duckduckgo.com")
	nav, err := c.Page.Navigate(ctx, navArgs)
	if err != nil {
		panic(err)
	}

	// Block until a DOM ContentEventFired event is triggered.
	if _, err = domContentEventFired.Recv(); err != nil {
		panic(err)
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	// Fetch the document root node. We can pass nil here
	// since this method only takes optional arguments.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Get the outer HTML for the page.
	result, err := c.DOM.GetOuterHTML(ctx, dom.NewGetOuterHTMLArgs(doc.Root.NodeID))
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTML: %s\n", result.OuterHTML)

	// Capture a screenshot of the current page.
	screenshotName := "screenshot.jpg"
	screenshotArgs := page.NewCaptureScreenshotArgs().
		SetFormat("jpeg").
		SetQuality(80)
	screenshot, err := c.Page.CaptureScreenshot(ctx, screenshotArgs)
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
