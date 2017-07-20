package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(context.TODO(), dir); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, dir string) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	devt := devtool.New("http://localhost:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		return err
	}

	conn, err := rpcc.Dial(pt.WebSocketDebuggerURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	c := cdp.NewClient(conn)

	input, err := os.Create(filepath.Join(dir, "log.input"))
	if err != nil {
		return err
	}

	consoleAPICalled, err := c.Runtime.ConsoleAPICalled(ctx)
	if err != nil {
		return err
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
				log.Println(err)
				return
			}
		}
	}()

	domLoadTimeout := 5 * time.Second

	// First page load is to trigger console log behavior without object
	// previews.
	if err := navigate(c.Page, "file:///"+dir+"/log.html", domLoadTimeout); err != nil {
		return err
	}

	// Enable console log events.
	if err := c.Runtime.Enable(ctx); err != nil {
		return err
	}

	// Re-load the page to receive console logs with previews.
	if err := navigate(c.Page, "file:///"+dir+"/log.html", domLoadTimeout); err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond)

	if err := input.Close(); err != nil {
		return err
	}

	return nil
}

// navigate to the URL and wait for DOMContentEventFired. An error is
// returned if timeout happens before DOMContentEventFired.
func navigate(pc cdp.Page, url string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Enable the Page domain events.
	if err := pc.Enable(ctx); err != nil {
		return err
	}

	// Open client for DOMContentEventFired to pause execution until
	// DOM has fully loaded.
	domContentEventFired, err := pc.DOMContentEventFired(ctx)
	if err != nil {
		return err
	}
	defer domContentEventFired.Close()

	_, err = pc.Navigate(ctx, page.NewNavigateArgs(url))
	if err != nil {
		return err
	}

	_, err = domContentEventFired.Recv()
	return err
}
