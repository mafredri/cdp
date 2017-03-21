package cdp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/rpcc"
)

// LogCodec captures the output from writing RPC requests and reading
// responses on the connection. It implements rpcc.Codec via
// WriteRequest and ReadResponse.
type LogCodec struct{ conn net.Conn }

// WriteRequest marshals v into a buffer, writes its contents onto the
// connection and logs it.
func (c *LogCodec) WriteRequest(v interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		return err
	}
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	fmt.Printf("SEND: %s", buf.Bytes())
	return nil
}

// ReadResponse unmarshals from the connection into v whilst echoing
// what is read into a buffer for logging.
func (c *LogCodec) ReadResponse(v interface{}) error {
	var buf bytes.Buffer
	if err := json.NewDecoder(io.TeeReader(c.conn, &buf)).Decode(v); err != nil {
		return err
	}
	fmt.Printf("RECV: %s\n", buf.String())
	return nil
}

func Example_logging() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newLogCodec := func(conn net.Conn) rpcc.Codec {
		return &LogCodec{conn: conn}
	}
	conn, err := rpcc.Dial("ws://"+TestSockSrv+"/example_logging", rpcc.WithCodec(newLogCodec))
	if err != nil {
		fmt.Println(err)
	}
	c := cdp.NewClient(conn)

	if err = c.Network.Enable(ctx, nil); err != nil {
		fmt.Println(err)
	}
	// Output:
	// SEND: {"id":1,"method":"Network.enable"}
	// RECV: {"id":1,"result":{}}
}
