/*

Package cdp provides type-safe bindings for the Chrome Debugging
Protocol (CDP) and can be used with any debug target that implements it.

The cdp Client requires an rpcc connection (*rpcc.Conn):

	conn, err := rpcc.Dial("ws://127.0.0.1:9222/f39a3624-e972-4a77-8a5f-6f8c42ef5129")
	if err != nil {
		// Handle error.
	}
	defer conn.Close()
	c := cdp.NewClient(conn)
	// ...

Domain methods

Methods can be invoked from the Client:

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := cdp.NewClient(conn)
	nav, err := c.Page.Navigate(ctx, cdp.NewPageNavigateArgs("https://www.google.com"))
	if err != nil {
		// Handle error.
	}
	// ...

Domain events

Events are received with an event client:

	// DOMContentEventFired = DOMContentLoaded.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		// Handle error.
	}
	ev, err := domContentEventFired.Recv()
	if err != nil {
		// Handle error.
	}
	// ...

Enable must be called before events are triggered for the domain:

	err := c.Page.Enable(ctx)
	if err != nil {
		// Handle error.
	}
	// ...

Some events are sent immediately after the call to Enable, it is a good
idea to create event clients beforehand. The rpcc.Stream will buffer the
events until they are ready to be received via Recv.

*/
package cdp
