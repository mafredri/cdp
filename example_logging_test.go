package cdp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/rpcc"
)

// LogCodec captures the output from writing RPC requests and reading
// responses on the connection. It implements rpcc.Codec via
// WriteRequest and ReadResponse.
type LogCodec struct{ conn io.ReadWriter }

// WriteRequest marshals v into a buffer, writes its contents onto the
// connection and logs it.
func (c *LogCodec) WriteRequest(req *rpcc.Request) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return err
	}
	fmt.Printf("SEND: %s", buf.Bytes())
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// ReadResponse unmarshals from the connection into v whilst echoing
// what is read into a buffer for logging.
func (c *LogCodec) ReadResponse(resp *rpcc.Response) error {
	var buf bytes.Buffer
	if err := json.NewDecoder(io.TeeReader(c.conn, &buf)).Decode(resp); err != nil {
		return err
	}
	fmt.Printf("RECV: %s\n", buf.String())
	return nil
}

func Example_logging() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newLogCodec := func(conn io.ReadWriter) rpcc.Codec {
		return &LogCodec{conn: conn}
	}
	conn, err := rpcc.Dial("ws://"+TestSockSrv+"/example_logging", rpcc.WithCodec(newLogCodec))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := cdp.NewClient(conn)

	if err = c.Network.Enable(ctx, nil); err != nil {
		fmt.Println(err)
	}
	// Output:
	// SEND: {"id":1,"method":"Network.enable"}
	// RECV: {"id":1,"result":{}}
}
