# cdp

[![Build Status][travis-badge]][travis] [![Coverage Status][codecov-badge]][codecov] [![Go Report Card][goreportcard-badge]][goreportcard] [![GoDoc][godoc-badge]][godoc]

Package `cdp` provides type-safe bindings for the [Chrome Debugging Protocol][devtool-protocol] (CDP), written in the Go programming language. The bindings are generated (by [cdpgen][cdpgen]) from the latest [tip-of-tree (tot)][tip-of-tree] protocol definitions and are mainly intended for use with Google Chrome or Chromium, however, they can be used with any debug target ([Node.js][node-debugging], [Edge][edge-diagnostics-adapter], [Safari][ios-webkit-debug-proxy], etc.) that implement the protocol.

This package can be used for any kind of browser automation, scripting or debugging via the Chrome Debugging Protocol.

A big motivation for `cdp` is to expose the full functionality of the Chrome Debugging Protocol and provide it in a discoverable and self-documenting manner.

Providing high-level browser automation is a non-goal for this project. That being said, `cdp` hopes to improve the ergonomics of working with the protocol by providing primitives better suited for Go and automating repetitive tasks.

## Features

* Discoverable API for the Chrome Debugging Protocol (GoDoc, autocomplete friendly)
* Contexts as a first-class citizen (for timeouts and cancellation)
* Simple and synchronous event handling (no callbacks)
* Concurrently safe
* No silent or hidden errors
* Do what the user expects
* Match CDP types to Go types wherever possible
* Separation of concerns (avoid mixing CDP and RPC)

## Installation

```console
$ go get -u github.com/mafredri/cdp
```

## Documentation

See [API documentation][godoc] for package, API descriptions and examples. Examples can also be found in this repository, see the [simple][simple-example], [advanced][advanced-example] and [logging][logging-example] examples.

## Usage

The main packages are `cdp` and `rpcc`, the former provides the CDP bindings and the latter handles the RPC communication with the debugging target.

To connect to a debug target, a WebSocket debugger URL is needed. For example, if Chrome is running with `--remote-debugging-port=9222` the debugger URL can be found at [localhost:9222/json](http://localhost:9222/json). The `devtool` package can also be used to query the DevTools JSON API (see example below).

Here is an example of using `cdp`:

[embedmd]:# (example_test.go)
```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func main() {
	err := run(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

func run(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return err
		}
	}

	// Initiate a new RPC connection to the Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return err
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	// Open a DOMContentEventFired client to buffer this event.
	domContent, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		return err
	}
	defer domContent.Close()

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = c.Page.Enable(ctx); err != nil {
		return err
	}

	// Create the Navigate arguments with the optional Referrer field set.
	navArgs := page.NewNavigateArgs("https://www.google.com").
		SetReferrer("https://duckduckgo.com")
	nav, err := c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return err
	}

	// Wait until we have a DOMContentEventFired event.
	if _, err = domContent.Recv(); err != nil {
		return err
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	// Fetch the document root node. We can pass nil here
	// since this method only takes optional arguments.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		return err
	}

	// Get the outer HTML for the page.
	result, err := c.DOM.GetOuterHTML(ctx, &dom.GetOuterHTMLArgs{
		NodeID: &doc.Root.NodeID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("HTML: %s\n", result.OuterHTML)

	// Capture a screenshot of the current page.
	screenshotName := "screenshot.jpg"
	screenshotArgs := page.NewCaptureScreenshotArgs().
		SetFormat("jpeg").
		SetQuality(80)
	screenshot, err := c.Page.CaptureScreenshot(ctx, screenshotArgs)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0644); err != nil {
		return err
	}

	fmt.Printf("Saved screenshot: %s\n", screenshotName)

	return nil
}
```

For more information, consult the [documentation](#documentation).

## Acknowledgements

The Go implementation of gRPC ([grpc-go](https://github.com/grpc/grpc-go)) has been a source of inspiration for some of the design descisions made in the `cdp` and `rpcc` packages. Some ideas have also been borrowed from the `net/rpc` package from the standard library.

## Resources

* [Chrome DevTools Protocol][devtool-protocol]
    * [Viewer (latest tip-of-tree)][tip-of-tree] official protocol API docs
    * [Repository (GitHub)](https://github.com/chromedevtools/devtools-protocol) please [file issues](https://github.com/ChromeDevTools/devtools-protocol/issues) at this repo if you have concerns or problems with the Chrome Debugging Protocol
    * [Mailing list](https://groups.google.com/forum/#!forum/chrome-debugging-protocol)
* [Getting Started with Headless Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome)
* [Awesome chrome-devtools: Chrome DevTools Protocol](https://github.com/ChromeDevTools/awesome-chrome-devtools#chrome-devtools-protocol)
* [RemoteDebug](http://remotedebug.org/) is an initiative to bring remote debugging (e.g. CDP) to all modern browsers
    * [Protocol Compatibility Tables](https://compatibility.remotedebug.org/)

[cdpgen]: https://github.com/mafredri/cdp/tree/master/cmd/cdpgen
[simple-example]: https://github.com/mafredri/cdp/blob/master/example_test.go
[advanced-example]: https://github.com/mafredri/cdp/blob/master/example_advanced_test.go
[logging-example]: https://github.com/mafredri/cdp/blob/master/example_logging_test.go

[devtool-protocol]: https://chromedevtools.github.io/devtools-protocol/
[tip-of-tree]: https://chromedevtools.github.io/devtools-protocol/tot/
[node-debugging]: https://nodejs.org/en/docs/guides/debugging-getting-started/
[edge-diagnostics-adapter]: https://github.com/Microsoft/edge-diagnostics-adapter
[ios-webkit-debug-proxy]: https://github.com/google/ios-webkit-debug-proxy

[travis]: https://travis-ci.org/mafredri/cdp
[travis-badge]: https://travis-ci.org/mafredri/cdp.svg
[codecov]: https://codecov.io/gh/mafredri/cdp
[codecov-badge]: https://codecov.io/gh/mafredri/cdp/branch/master/graph/badge.svg
[goreportcard]: https://goreportcard.com/report/github.com/mafredri/cdp
[goreportcard-badge]: https://goreportcard.com/badge/github.com/mafredri/cdp
[godoc]: https://godoc.org/github.com/mafredri/cdp
[godoc-badge]: https://godoc.org/mafredri/cdp?status.svg
