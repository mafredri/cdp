package cdp_test

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/coder/websocket"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"
)

func Example_dial_using_alternative_websocket_implementation() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	devt := devtool.New("http://localhost:9222")
	page, err := devt.CreateURL(ctx, "about:blank")
	if err != nil {
		log.Println(err)
	}

	// Dial using an alternative websocket implementation.
	//
	// Note that this disables functionality like:
	//
	// - Safety measure against writing fragmented websocket messages
	// - Setting compression level after dial
	dialer := rpcc.WithDialer(func(dialCtx context.Context, addr string) (io.ReadWriteCloser, error) {
		log.Println(addr)
		conn, _, err := websocket.Dial(dialCtx, addr, &websocket.DialOptions{
			CompressionMode: websocket.CompressionContextTakeover,
		})
		if err != nil {
			return nil, err
		}
		// Note that we cannot use dialCtx here since websocket.NetConn
		// binds to the lifetime of ctx.
		return websocket.NetConn(ctx, conn, websocket.MessageText), nil
	})
	conn, err := rpcc.Dial(page.WebSocketDebuggerURL, dialer)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// Use the connection that uses nhooyr.io/websocket underneath.
	c := cdp.NewClient(conn)

	if err = c.Runtime.Enable(ctx); err != nil {
		log.Println(err)
	}
	eval, err := c.Runtime.Evaluate(ctx, runtime.NewEvaluateArgs(`document.location.href`).SetReturnByValue(true))
	if err == nil && eval.ExceptionDetails != nil {
		err = eval.ExceptionDetails
	}
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eval.Result.String())

	err = devt.Close(ctx, page)
	if err != nil {
		log.Println(err)
	}
	// Output:
	// "about:blank"
}
