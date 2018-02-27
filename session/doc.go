/*

Package session implements a session Manager for establishing session
connections to targets (via the Target domain). Session connections allow a
single websocket connection (from the provided cdp.Client) to be used for
communicating with multiple targets.

Initialize a new session Manager.

	c := cdp.NewClient(conn) // cdp.Client with websocket connection.

	m, err := session.NewManager(c)
	if err != nil {
		// Handle error.
	}
	defer m.Close() // Cleanup.

Establish a new session connection to targetID.

	pageConn, err := m.Dial(context.TODO(), targetID)
	if err != nil {
		// Handle error.
	}
	defer pageConn.Close()

Use the session connection.

	pageClient := cdp.NewClient(pageConn)
	err = pageClient.Page.Enable(context.TODO())
	// ...

*/
package session
