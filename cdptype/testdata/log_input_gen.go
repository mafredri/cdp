package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/cdpcmd"
	"github.com/mafredri/cdp/cdptype"
	"github.com/mafredri/cdp/rpcc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	chromium := "ws://localhost:9222/devtools/page/45a887ba-c92a-4cff-9194-d9398cc87e2c"
	conn, err := rpcc.Dial(chromium)
	if err != nil {
		panic(err)
	}
	c := cdp.NewClient(conn)

	input, err := os.Create(filepath.Join(dir, "log.input"))
	if err != nil {
		panic(err)
	}
	defer input.Close()

	consoleAPICalled, err := c.Runtime.ConsoleAPICalled(ctx)
	if err != nil {
		panic(err)
	}
	go func() {
		defer consoleAPICalled.Close()
		for {
			ev, err := consoleAPICalled.Recv()
			if err != nil {
				return
			}
			// Reset fields that would cause noise in diffs.
			ev.ExecutionContextID = 0
			ev.Timestamp = 0
			ev.StackTrace = nil
			for i, arg := range ev.Args {
				arg.ObjectID = nil
				ev.Args[i] = arg
			}
			if err = json.NewEncoder(input).Encode(ev); err != nil {
				panic(err)
			}
		}
	}()

	domLoadTimeout := 5 * time.Second

	// First page load is to trigger console log behavior without object
	// previews.
	_, err = navigate(c.Page, "file:///"+dir+"/log.html", domLoadTimeout)
	if err != nil {
		panic(err)
	}

	// Enable console log events.
	if err = c.Runtime.Enable(ctx); err != nil {
		panic(err)
	}

	// Re-load the page to receive console logs with previews.
	_, err = navigate(c.Page, "file:///"+dir+"/log.html", domLoadTimeout)
	if err != nil {
		panic(err)
	}

	time.Sleep(250 * time.Millisecond)
}

// navigate to the URL and wait for DOMContentEventFired. An error is
// returned if timeout happens before DOMContentEventFired.
func navigate(page cdp.Page, url string, timeout time.Duration) (frame cdptype.PageFrameID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Enable the Page domain events.
	if err := page.Enable(ctx); err != nil {
		panic(err)
	}

	// Open client for DOMContentEventFired to pause execution until
	// DOM has fully loaded.
	domContentEventFired, err := page.DOMContentEventFired(ctx)
	if err != nil {
		return frame, err
	}
	defer domContentEventFired.Close()

	nav, err := page.Navigate(ctx, cdpcmd.NewPageNavigateArgs(url))
	if err != nil {
		return frame, err
	}

	if _, err = domContentEventFired.Recv(); err != nil {
		return frame, err
	}

	return nav.FrameID, nil
}
