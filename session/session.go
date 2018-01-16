package session

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// session represents a session connection to a target.
type session struct {
	ID       target.SessionID
	TargetID target.ID
	conn     *rpcc.Conn
	recvC    chan []byte
	send     func([]byte) error
	detach   func()
}

// Ensure that session implements rpcc.Codec.
var _ rpcc.Codec = (*session)(nil)

// WriteRequest implements rpcc.Codec.
func (s *session) WriteRequest(r *rpcc.Request) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return s.send(data)
}

// ReadResponse implements rpcc.Codec.
func (s *session) ReadResponse(r *rpcc.Response) error {
	select {
	case m := <-s.recvC:
		return json.Unmarshal(m, r)
	case <-s.conn.Context().Done():
		s.detach()
		return s.conn.Context().Err()
	}
}

// Conn returns the underlying *rpcc.Conn that uses session as codec.
func (s *session) Conn() *rpcc.Conn { return s.conn }

// Write forwards a target message to the session connection.
func (s *session) Write(data []byte) {
	s.recvC <- data
}

// Close closes the underlying *rpcc.Conn.
func (s *session) Close() error {
	// Closing conn will trigger s.detach via ReadResponse.
	return s.conn.Close()
}

var (
	// We don't need to establish a connection because the codec does not
	// use conn for transmission.
	nilConn = rpcc.WithDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return nil, nil
	})
	codec = func(s *session) rpcc.DialOption {
		return rpcc.WithCodec(func(_ io.ReadWriter) rpcc.Codec {
			return s
		})
	}
)

// dial attaches to the target via the provided *cdp.Client and creates
// a lightweight RPC connection to the target. Communication is done via
// the underlying *rpcc.Conn for the provided *cdp.Client.
func dial(ctx context.Context, id target.ID, tc *cdp.Client) (s *session, err error) {
	args := target.NewAttachToTargetArgs(id)
	reply, err := tc.Target.AttachToTarget(ctx, args)
	if err != nil {
		return nil, err
	}

	s = &session{
		TargetID: id,
		ID:       reply.SessionID,
		recvC:    make(chan []byte, 1),
		send: func(data []byte) error {
			// TODO(maf): Use async invocation.
			// This context is unavailable until s.conn is assigned,
			// but no messages will be sent before that happens.
			return tc.Target.SendMessageToTarget(s.conn.Context(),
				target.NewSendMessageToTargetArgs(string(data)).
					SetSessionID(s.ID))
		},
		detach: func() {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			// TODO(maf): Use async invocation and ignore error.
			go tc.Target.DetachFromTarget(ctx,
				target.NewDetachFromTargetArgs().SetSessionID(s.ID))
		},
	}

	s.conn, err = rpcc.Dial("", nilConn, codec(s))
	if err != nil {
		return nil, err
	}

	return s, nil
}
