package cdp_test

import (
	"context"
	"fmt"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/rpcc"
)

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to WebSocket URL (page) that speaks the Chrome Debugging Protocol.
	conn, err := rpcc.DialContext(ctx, "ws://localhost:9222/devtools/page/45a887ba-c92a-4cff-9194-d9398cc87e2c")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // Close connection when we're done.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	// Enable the Page domain (enables events for the domain).
	if err = c.Page.Enable(ctx); err != nil {
		panic(err)
	}

	// Open client for DOMContentEventFired to block until DOM has fully loaded.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		panic(err)
	}
	defer domContentEventFired.Close()

	// Initialize the Navigate args with the optional argument Referrer set.
	navigateArgs := cdp.NewPageNavigateArgs("https://www.google.com").SetReferrer("https://www.google.com")
	nav, err := c.Page.Navigate(ctx, navigateArgs)
	if err != nil {
		panic(err)
	}

	// Wait for the event.
	if _, err = domContentEventFired.Recv(); err != nil {
		panic(err)
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	// Fetch the document root node. DOM GetDocument only
	// has optional arguments, we can pass nil here.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Get the HTML for the page.
	outer, err := c.DOM.GetOuterHTML(ctx, cdp.NewDOMGetOuterHTMLArgs(doc.Root.NodeID))
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTML: %s\n", outer.OuterHTML)
}
