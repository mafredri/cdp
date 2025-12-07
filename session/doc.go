/*
Package session implements a session Manager for establishing session
connections to targets (via the Target domain). Session connections allow a
single websocket connection (from the provided cdp.Client) to be used for
communicating with multiple targets.

By default, sessions use flattened mode (Chrome 77+) where the session ID is
included directly in CDP messages.

Initialize a new session Manager:

	c := cdp.NewClient(conn) // cdp.Client with websocket connection.

	m, err := session.NewManager(c)
	if err != nil {
		// Handle error.
	}
	defer m.Close() // Cleanup.

For older Chrome versions, use legacy mode:

	m, err := session.NewManager(c, session.WithNoFlatten())

Establish a new session connection to targetID:

	pageConn, err := m.Dial(ctx, targetID)
	if err != nil {
		// Handle error.
	}
	defer pageConn.Close()

Use the session connection:

	pageClient := cdp.NewClient(pageConn)
	err = pageClient.Page.Enable(ctx)
	// ...

If session connections are behaving unexpectedly, you can debug the session
Manager by checking the error channel:

	go func() {
		for err := range m.Err() {
			log.Println(err)
		}
		// Manager is closed.
	}()

# Manual session management (flattened protocol)

For more control, flattened sessions can be created directly via rpcc.NewSession
without using the Manager.

	reply, err := c.Target.AttachToTarget(ctx, target.NewAttachToTargetArgs(targetID).SetFlatten(true))
	if err != nil {
		// Handle error.
	}

	sessConn, err := rpcc.NewSession(conn, string(reply.SessionID), rpcc.WithSessionClose(func(ctx context.Context) error {
		args := target.NewDetachFromTargetArgs().SetSessionID(reply.SessionID)
		return c.Target.DetachFromTarget(ctx, args)
	}))
	if err != nil {
		// Handle error.
	}
	defer sessConn.Close()

	sessClient := cdp.NewClient(sessConn)
	// ...

The WithSessionClose function handles detachment when you close the session.
Note that Chrome may also detach sessions (e.g. when a target is destroyed).
For short-lived operations this can be ignored, but for long-lived sessions
you should listen for Target.DetachedFromTarget events or use the Manager.
*/
package session
