# cdp

Package `cdp` provides type-safe bindings for the [Chrome Debugging Protocol](https://developer.chrome.com/devtools/docs/debugger-protocol) (CDP), written in the Go programming language. The bindings are generated with the latest [js_protocol.json](https://chromium.googlesource.com/chromium/src/+/master/third_party/WebKit/Source/core/inspector/browser_protocol.json) and [browser_protocol.json](https://chromium.googlesource.com/v8/v8.git/+/master/src/inspector/js_protocol.json) from the Chromium repository using [cmd/cdpgen](https://github.com/mafredri/cdp/blob/master/cmd/cdpgen). These bindings can be used with any debug target that implements the protocol.

A big motivation for `cdp` is to expose the full functionality of the Chrome Debugging Protocol and provide it in a discoverable and self-documenting manner.

Providing high-level browser automation is a non-goal for this project. That being said, helpful helpers and useful utility functions could be added.

## Goals

* Discoverable API for the Chrome Debugging Protocol (GoDoc, autocomplete friendly)
* User should be able to control when and what events are triggered / listened to
* Concurrently safe without implementing concurrency patterns in the API
* No silent or hidden errors, this is why creating event clients and receiving events return errors
* Do what the user expects
* Match CDP types to Go types wherever possible
* Separation of concerns (avoid mixing CDP and RPC)

## Installation

```console
$ go get -u github.com/mafredri/cdp
```

## Documentation

See [API documentation](https://godoc.org/github.com/mafredri/cdp) for package, API descriptions and examples. Examples can also be found in this repository, see the [simple](https://github.com/mafredri/cdp/blob/master/example_test.go), [advanced](https://github.com/mafredri/cdp/blob/master/example_advanced_test.go) and [logging](https://github.com/mafredri/cdp/blob/master/example_logging_test.go) examples.

## Usage

The main packages are `cdp` and `rpcc`, the former provides the CDP bindings and the latter handles the RPC communication with the debugging target.

To connect to a debug target, a WebSocket debugger URL is needed. For example, if Chrome is running with `--remote-debugging-port=9222` the debugger URL can be found at [localhost:9222/json](http://localhost:9222/json) (see `webSocketDebuggerUrl`). This functionality might be provided in the future by either a separate package or a subpackage.

Here is an example of using `cdp`:

```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/rpcc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to Chrome Debugging Protocol target (webSocketDebuggerUrl).
	wsURL := "ws://localhost:9222/devtools/page/45a887ba-c92a-4cff-9194-d9398cc87e2c"
	conn, err := rpcc.DialContext(ctx, wsURL)
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
	navArgs := cdp.NewPageNavigateArgs("https://www.google.com").SetReferrer("https://duckduckgo.com")
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
	result, err := c.DOM.GetOuterHTML(ctx, cdp.NewDOMGetOuterHTMLArgs(doc.Root.NodeID))
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTML: %s\n", result.OuterHTML)

	// Capture a screenshot of the current page.
	screenshotName := "screenshot.jpg"
	screenshot, err := c.Page.CaptureScreenshot(ctx, cdp.NewPageCaptureScreenshotArgs().SetFormat("jpeg").SetQuality(80))
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
```

For more information, consult the [documentation](#documentation).

## Acknowledgements

The Go implementation of gRPC ([grpc-go](https://github.com/grpc/grpc-go)) has been a source of inspiration for some of the design descisions made in the `cdp` and `rpcc` packages. Some ideas have also been borrowed from the `net/rpc` package from the standard library.

## Other work

These are alternative implementations of the Chrome Debugging Protocol, written in Go:

* [gcd](https://github.com/wirepair/gcd): Low-level client library for communicating with Google Chrome
* [autogcd](https://github.com/wirepair/autogcd): Wrapper around gcd to enable browser automation
* [chromedp](https://github.com/knq/chromedp): High-level API for driving web browsers
