/*

Package rpcc provides an RPC client connection with support for the
JSON-RPC 2.0 specification, not including Batch requests. Server side
RPC notifications are also supported.

Dial connects to an RPC server listening on a websocket using the
gorilla/websocket package.

	conn, err := rpcc.Dial("ws://127.0.0.1:9999/f39a3624-e972-4a77-8a5f-6f8c42ef5129")
	// ...

The user must close the connection when finnished with it:

	conn, err := rpcc.Dial("ws://127.0.0.1:9999/f39a3624-e972-4a77-8a5f-6f8c42ef5129")
	if err != nil {
		// Handle error.
	}
	defer conn.Close()
	// ...

A custom dialer can be used to change the websocket lib or communicate
over other protocols.

	netDial := func(ctx context.Context, addr string) (net.Conn, error) {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// Handle error.
		}
		// Wrap connection to handle writing JSON.
		// ...
		return conn, nil
	}
	conn, err := rpcc.Dial("127.0.0.1:9999", rpcc.WithDialer(netDial))
	// ...

Communicating with the server

Send a request using Invoke:

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := rpcc.Invoke(ctx, "Domain.method", args, reply, conn)
	// ...

Receive a notification using NewStream:

	stream, err := rpcc.NewStream(ctx, "Domain.event", conn)
	if err != nil {
		// Handle error.
	}
	err = stream.RecvMsg(&reply)
	if err != nil {
		// Handle error.
	}

The stream should be closed when it is no longer used to avoid leaking
memory:

	stream, err := rpcc.NewStream(ctx, "Domain.event", conn)
	if err != nil {
		// Handle error.
	}
	defer stream.Close()

When order is important, two streams can be synchronized with Sync:

	err := rpcc.Sync(stream1, stream2)
	if err != nil {
		// Handle error.
	}

*/
package rpcc
