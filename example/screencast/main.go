package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	devt := devtool.New("http://localhost:9222")

	pageTarget, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		return err
	}

	conn, err := rpcc.DialContext(ctx, pageTarget.WebSocketDebuggerURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	c := cdp.NewClient(conn)

	err = c.Page.Enable(ctx)
	if err != nil {
		return err
	}

	// Navigate to GitHub, block until ready.
	loadEventFired, err := c.Page.LoadEventFired(ctx)
	if err != nil {
		return err
	}

	_, err = c.Page.Navigate(ctx, page.NewNavigateArgs("https://github.com"))
	if err != nil {
		return err
	}

	_, err = loadEventFired.Recv()
	if err != nil {
		return err
	}
	loadEventFired.Close()

	// Start listening to ScreencastFrame events.
	screencastFrame, err := c.Page.ScreencastFrame(ctx)
	if err != nil {
		return err
	}

	go func() {
		defer screencastFrame.Close()

		for {
			ev, err := screencastFrame.Recv()
			if err != nil {
				log.Printf("Failed to receive ScreencastFrame: %v", err)
				return
			}
			log.Printf("Got frame with sessionID: %d: %+v", ev.SessionID, ev.Metadata)

			err = c.Page.ScreencastFrameAck(ctx, page.NewScreencastFrameAckArgs(ev.SessionID))
			if err != nil {
				log.Printf("Failed to ack ScreencastFrame: %v", err)
				return
			}

			// Write to screencast_frame-[timestamp].png.
			name := fmt.Sprintf("screencast_frame-%d.png", ev.Metadata.Timestamp.Time().Unix())

			// Write the frame to file (without blocking).
			go func() {
				err = ioutil.WriteFile(name, ev.Data, 0644)
				if err != nil {
					log.Printf("Failed to write ScreencastFrame to %q: %v", name, err)
				}
			}()
		}
	}()

	screencastArgs := page.NewStartScreencastArgs().
		SetEveryNthFrame(1).
		SetFormat("png")
	err = c.Page.StartScreencast(ctx, screencastArgs)
	if err != nil {
		return err
	}

	// Random delay for our screencast.
	time.Sleep(30 * time.Second)

	err = c.Page.StopScreencast(ctx)
	if err != nil {
		return err
	}

	return nil
}
