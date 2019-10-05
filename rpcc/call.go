package rpcc

import "context"

type rpcCall struct {
	Method    string
	SessionID string
	Args      interface{}
	Reply     interface{}
	Error     chan error
}

func (c *rpcCall) done(err error) {
	c.Error <- err
}

// Deprecated: Invoke is kept for compatibility, but will be removed
// in a later release. Please call InvokeRPC instead.
func Invoke(ctx context.Context, method string, args, reply interface{}, conn *Conn) error {
	return InvokeRPC(ctx, method, "", args, reply, conn)
}

// InvokeRPC sends an RPC request and blocks until the response is received.
// This function is called by generated code but can be used to issue
// requests manually.
func InvokeRPC(ctx context.Context, method, sessionID string, args, reply interface{}, conn *Conn) error {
	if ctx == nil {
		ctx = context.Background()
	}

	call := &rpcCall{
		Method:    method,
		SessionID: sessionID,
		Args:      args,
		Reply:     reply,
		Error:     make(chan error, 1), // Do not block.
	}

	err := conn.send(ctx, call)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err = <-call.Error:
		return err
	}
}
