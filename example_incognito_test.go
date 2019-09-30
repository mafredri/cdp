package cdp_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp/session"
)

func Example_incognito() {
	if !*testBrowser {
		return	
	}
	err := func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Fetch the websocket URL for the browser endpoint.
		bver, err := devtool.New(fmt.Sprintf("http://localhost:%d", *remoteDebuggingPort)).Version(ctx)
		if err != nil {
			return err
		}
		bconn, err := rpcc.DialContext(ctx, bver.WebSocketDebuggerURL)
		if err != nil {
			return err
		}
		defer bconn.Close()

		// Initialize the browser CDP client.
		bc := cdp.NewClient(bconn)

		// Initialize session manager for connecting to targets.
		sess, err := session.NewManager(bc)
		if err != nil {
			return err
		}
		defer sess.Close()

		// Create the new browser context, similar to a new incognito
		// window.
		createCtx, err := bc.Target.CreateBrowserContext(ctx)
		if err != nil {
			return err
		}

		// Create a new target belonging to the browser context, similar
		// to opening a new tab in an incognito window.
		createTargetArgs := target.NewCreateTargetArgs("about:blank").
			SetBrowserContextID(createCtx.BrowserContextID)
		createTarget, err := bc.Target.CreateTarget(ctx, createTargetArgs)
		if err != nil {
			return err
		}

		// Connect to target using the existing websocket connection.
		conn, err := sess.Dial(ctx, createTarget.TargetID)
		if err != nil {
			return err
		}
		defer conn.Close()

		// This cdp client controls the "incognito tab".
		c := cdp.NewClient(conn)

		err = c.Page.Enable(ctx)
		if err != nil {
			return err
		}

		url := "https://github.com/mafredri/cdp"
		nav, err := c.Page.Navigate(ctx, page.NewNavigateArgs(url))
		if err != nil {
			return err
		}

		if nav.ErrorText != nil {
			return errors.New("navigation failed: " + *nav.ErrorText)
		}

		// Check the window.location for validation.
		eval, err := c.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs("window.location.toString();"))
		if err != nil {
			return err
		}
		if eval.ExceptionDetails != nil {
			return eval.ExceptionDetails
		}
		var winloc string
		err = json.Unmarshal(eval.Result.Value, &winloc)
		if err != nil {
			return err
		}

		fmt.Printf("Navigated to %s inside an incognito tab!\n", winloc)

		// Close the tab when we are done with it (this is a requirement
		// for closing the browser context).
		closeReply, err := bc.Target.CloseTarget(ctx, target.NewCloseTargetArgs(createTarget.TargetID))
		if err != nil {
			return err
		}
		if !closeReply.Success {
			return errors.New("could not close target: " + string(createTarget.TargetID))
		}

		// Dispose of browser context (a.k.a. incognito window), this
		// will fail if not all targets in this context are closed
		// beforehand.
		err = bc.Target.DisposeBrowserContext(ctx, target.NewDisposeBrowserContextArgs(createCtx.BrowserContextID))
		if err != nil {
			return err
		}

		return nil
	}()
	if err != nil {
		fmt.Println(err)
	}
	// Output: Navigated to https://github.com/mafredri/cdp inside an incognito tab!
}
