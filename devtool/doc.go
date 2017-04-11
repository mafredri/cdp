/*

Package devtool provides methods for interacting with a DevTools
endpoint.

To activate the DevTools endpoint, a browser (or other debug target)
should be started with debugging enabled:

	chromium --remote-debugging-port=9222
	./<path>/EdgeDiagnosticsAdapter.exe --port 9223
	node --inspect=9224

Create a new DevTools instance that interacts with the given URL:

	devt := devtool.New("http://127.0.0.1:9222")

Get the active page or create a new one:

	devt := devtool.New("http://127.0.0.1:9222")
	page, err := devt.Get(context.Background(), devtool.Page)
	if err != nil {
		page, err = devt.Create(context.Background())
		if err != nil {
			// Handle error.
		}
	}
	// ...

Set request timeouts via contexts:

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	devt := devtool.New("http://127.0.0.1:9222")
	list, err := devt.List(ctx)
	if err != nil {
		// Handle error.
	}
	// ...

*/
package devtool
