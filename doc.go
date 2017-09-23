/*

Package cdp provides type-safe bindings for the Chrome Debugging
Protocol (CDP) and can be used with any debug target that implements it.

The cdp Client requires an rpcc connection (*rpcc.Conn):

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	conn, err := rpcc.DialContext(ctx, "ws://127.0.0.1:9222/f39a3624-e972-4a77-8a5f-6f8c42ef5129")
	if err != nil {
		// Handle error.
	}
	defer conn.Close()

	c := cdp.NewClient(conn)
	// ...

The devtool package can be used for finding the websocket URL (see
devtool documentation for more):

	devt := devtool.New("http://127.0.0.1:9222")
	pg, err := devtool.Get(ctx, devtool.Page)
	if err != nil {
		// Handle error.
	}
	conn, err := rpcc.Dial(pg.WebSocketDebuggerURL)
	// ...

Domain methods

Domain methods are used to perform actions or request data over the
Chrome Debugging Protocol.

Methods can be invoked from the Client:

	c := cdp.NewClient(conn)
	nav, err := c.Page.Navigate(ctx, page.NewNavigateArgs("https://www.google.com"))
	if err != nil {
		// Handle error.
	}
	// ...

Domain events

Event clients are used to handle events sent over the protocol. A client
will buffer all events, preserving order, after creation until it is
closed, context done or connection closed. Under the hood, an event
client is a rpcc.Stream.

Create an event client for the DOMContentEventFired event. Call Close
when the client is no longer used to avoid leaking memory. The client
will remain active for the duration of the context or until it is
closed:

	// DOMContentEventFired = DOMContentLoaded.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		// Handle error.
	}
	defer domContentEventFired.Close()
	// ...

Enable (if available) must be called before events are transmitted over
the Chrome Debugging Protocol:

	err := c.Page.Enable(ctx)
	if err != nil {
		// Handle error.
	}
	// ...

Calling Enable can result in immediate event transmissions. If these
events are important, an event client should be created before calling
Enable.

Wait for an event by calling Recv:

	ev, err := domContentEventFired.Recv()
	if err != nil {
		// Handle error.
	}
	// ...

The Ready channel can be used to check for pending events or
coordinating between multiple event handlers:

	go func() {
		select {
		case <-domContentEventFired.Ready():
			_, err := domContentEventFired.Recv() // Does not block here.
			if err != nil {
				// Handle error.
			}
		case <-loadEventFired.Ready():
			// Handle loadEventFired.
		}
	}()
	// ...

Ready must not be called concurrently while relying on the non-blocking
behavior of Recv.

Event clients can be synchronized, relative to each other, when the order of
events is important:

	err := cdp.Sync(domContentEventFired, loadEventFired)
	if err != nil {
		// Handle error.
	}

Use the Ready channel to detect which synchronized event client is ready to
Recv.

*/
package cdp

// Generate protcol definition using cdpgen.
//go:generate go install github.com/mafredri/cdp/cmd/cdpgen
//go:generate cdpgen -dest-pkg github.com/mafredri/cdp -browser-proto ./cmd/cdpgen/protodef/browser_protocol.json -js-proto ./cmd/cdpgen/protodef/js_protocol.json

// Update code samples in README.
//go:generate embedmd -w README.md
//go:generate -command sed sed -i ""
//go:generate sed -e "s/^package .*_test$/package main/" README.md
//go:generate sed -e "s/^func Example\\(_[^)]*\\)*() {$/func main() {/" README.md
