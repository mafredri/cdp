// +build go1.9

package cdp_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"

	"golang.org/x/sync/errgroup"
)

// Cookie represents a browser cookie.
type Cookie struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// DocumentInfo contains information about the document.
type DocumentInfo struct {
	Title string `json:"title"`
}

var (
	MyURL   = "https://google.com"
	Cookies = []Cookie{
		{MyURL, "myauth", "myvalue"},
		{MyURL, "mysetting1", "myvalue1"},
		{MyURL, "mysetting2", "myvalue2"},
		{MyURL, "mysetting3", "myvalue3"},
	}
)

func Example_advanced() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to WebSocket URL (page) that speaks the Chrome Debugging Protocol.
	conn, err := rpcc.DialContext(ctx, "ws://localhost:9222/devtools/page/45a887ba-c92a-4cff-9194-d9398cc87e2c")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // Cleanup.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	// Give enough capacity to avoid blocking any event listeners
	abort := make(chan error, 2)

	// Watch the abort channel.
	go func() {
		select {
		case <-ctx.Done():
		case err := <-abort:
			fmt.Printf("aborted: %s\n", err.Error())
			cancel()
		}
	}()

	// Setup event handlers early because domain events can be sent as
	// soon as Enable is called on the domain.
	if err = catchExceptionThrown(ctx, c.Runtime, abort); err != nil {
		fmt.Println(err)
		return
	}
	if err = catchLoadingFailed(ctx, c.Network, abort); err != nil {
		fmt.Println(err)
		return
	}

	if err = runBatch(
		// Enable all the domain events that we're interested in.
		func() error { return c.DOM.Enable(ctx) },
		func() error { return c.Network.Enable(ctx, nil) },
		func() error { return c.Page.Enable(ctx) },
		func() error { return c.Runtime.Enable(ctx) },

		func() error { return setCookies(ctx, c.Network, Cookies...) },
	); err != nil {
		fmt.Println(err)
		return
	}

	domLoadTimeout := 5 * time.Second
	frameID, err := navigate(ctx, c.Page, MyURL, domLoadTimeout)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Navigating with frame ID: %v\n", frameID)

	// Parse information from the document by evaluating JavaScript.
	expression := `
		new Promise((resolve, reject) => {
			setTimeout(() => {
				const title = document.querySelector('title').innerText;
				resolve({title});
			}, 500);
		});
	`
	evalArgs := runtime.NewEvaluateArgs(expression).SetAwaitPromise(true).SetReturnByValue(true)
	eval, err := c.Runtime.Evaluate(ctx, evalArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	var info DocumentInfo
	if err = json.Unmarshal(eval.Result.Value, &info); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Document title: %q\n", info.Title)

	// Fetch the document root node.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch all <script> and <noscript> elements so we can delete them.
	scriptIDs, err := c.DOM.QuerySelectorAll(ctx, dom.NewQuerySelectorAllArgs(doc.Root.NodeID, "script, noscript"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = removeNodes(ctx, c.DOM, scriptIDs.NodeIDs...); err != nil {
		fmt.Println(err)
		return
	}
}

func catchExceptionThrown(ctx context.Context, runtime cdp.Runtime, abort chan<- error) error {
	// Listen to exceptions so we can abort as soon as one is encountered.
	exceptionThrown, err := runtime.ExceptionThrown(ctx)
	if err != nil {
		// Connection is closed.
		return err
	}
	go func() {
		defer exceptionThrown.Close() // Cleanup.
		for {
			ev, err := exceptionThrown.Recv()
			if err != nil {
				// This could be any one of: connection closed,
				// context deadline or unmarshal failed.
				abort <- err
				return
			}

			// Ruh-roh! Let the caller know something went wrong.
			abort <- ev.ExceptionDetails
		}
	}()
	return nil
}

func catchLoadingFailed(ctx context.Context, net cdp.Network, abort chan<- error) error {
	// Check for non-canceled resources that failed to load.
	loadingFailed, err := net.LoadingFailed(ctx)
	if err != nil {
		return err
	}
	go func() {
		defer loadingFailed.Close()
		for {
			ev, err := loadingFailed.Recv()
			if err != nil {
				abort <- err
				return
			}

			// For now, most optional fields are pointers and must be
			// checked for nil.
			canceled := ev.Canceled != nil && *ev.Canceled

			if !canceled {
				abort <- fmt.Errorf("request %s failed: %s", ev.RequestID, ev.ErrorText)
			}
		}
	}()
	return nil
}

// setCookies sets all the provided cookies.
func setCookies(ctx context.Context, net cdp.Network, cookies ...Cookie) error {
	var cmds []runBatchFunc
	for _, c := range cookies {
		args := network.NewSetCookieArgs(c.Name, c.Value).SetURL(c.URL)
		cmds = append(cmds, func() error {
			reply, err := net.SetCookie(ctx, args)
			if err != nil {
				return err
			}
			if !reply.Success {
				return errors.New("could not set cookie")
			}
			return nil
		})
	}
	return runBatch(cmds...)
}

// navigate to the URL and wait for DOMContentEventFired. An error is
// returned if timeout happens before DOMContentEventFired.
func navigate(ctx context.Context, pageClient cdp.Page, url string, timeout time.Duration) (frame page.FrameID, err error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	// Make sure Page events are enabled.
	if err = pageClient.Enable(ctx); err != nil {
		return frame, err
	}

	// Open client for DOMContentEventFired to block until DOM has fully loaded.
	domContentEventFired, err := pageClient.DOMContentEventFired(ctx)
	if err != nil {
		return frame, err
	}
	defer domContentEventFired.Close()

	nav, err := pageClient.Navigate(ctx, page.NewNavigateArgs(url))
	if err != nil {
		return frame, err
	}

	_, err = domContentEventFired.Recv()
	return nav.FrameID, err
}

// removeNodes deletes all provided nodeIDs from the DOM.
func removeNodes(ctx context.Context, domClient cdp.DOM, nodes ...dom.NodeID) error {
	var rmNodes []runBatchFunc
	for _, id := range nodes {
		arg := dom.NewRemoveNodeArgs(id)
		rmNodes = append(rmNodes, func() error { return domClient.RemoveNode(ctx, arg) })
	}
	return runBatch(rmNodes...)
}

// runBatchFunc is the function signature for runBatch.
type runBatchFunc func() error

// runBatch runs all functions simultaneously and waits until
// execution has completed or an error is encountered.
func runBatch(fn ...runBatchFunc) error {
	eg := errgroup.Group{}
	for _, f := range fn {
		eg.Go(f)
	}
	return eg.Wait()
}
