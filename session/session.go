package session

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/internal/errors"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// session represents a session connection to a target.
type session struct {
	ID       target.SessionID
	TargetID target.ID
	recvC    chan []byte
	send     func([]byte) error

	init chan struct{} // Protect conn from early read.
	conn *rpcc.Conn
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
	<-s.init

	select {
	case m := <-s.recvC:
		return json.Unmarshal(m, r)
	case <-s.conn.Context().Done():
		return s.conn.Context().Err()
	}
}

// Conn returns the underlying *rpcc.Conn that uses session as codec.
func (s *session) Conn() *rpcc.Conn { return s.conn }

// Write forwards a target message to the session connection.
// When write returns an error, the session is closed.
func (s *session) Write(data []byte) error {
	select {
	case s.recvC <- data:
		return nil
	case <-s.conn.Context().Done():
		return s.conn.Context().Err()
	}
}

// Close closes the underlying *rpcc.Conn.
func (s *session) Close() error {
	return s.conn.Close()
}

var (
	// We only handle Close on conn to detach the session. The codec
	// handles the actual transport (Read / Write) in this case.
	sessionDetachConn = func(detach func() error) rpcc.DialOption {
		return rpcc.WithDialer(
			func(_ context.Context, _ string) (io.ReadWriteCloser, error) {
				return &closeConn{close: detach}, nil
			},
		)
	}
	sessionCodec = func(s *session) rpcc.DialOption {
		return rpcc.WithCodec(func(_ io.ReadWriter) rpcc.Codec {
			return s
		})
	}
)

// dial attaches to the target via the provided *cdp.Client and creates
// a lightweight RPC connection to the target. Communication is done via
// the underlying *rpcc.Conn for the provided *cdp.Client.
func dial(ctx context.Context, id target.ID, tc *cdp.Client, detachTimeout time.Duration) (s *session, err error) {
	args := target.NewAttachToTargetArgs(id)
	reply, err := tc.Target.AttachToTarget(ctx, args)
	if err != nil {
		return nil, err
	}

	s = &session{
		TargetID: id,
		ID:       reply.SessionID,
		recvC:    make(chan []byte, 1),
		init:     make(chan struct{}),
		send: func(data []byte) error {
			<-s.init
			// TODO(maf): Use async invocation.
			return tc.Target.SendMessageToTarget(s.conn.Context(),
				target.NewSendMessageToTargetArgs(string(data)).
					SetSessionID(s.ID))
		},
	}

	detach := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), detachTimeout)
		defer cancel()

		err := tc.Target.DetachFromTarget(ctx,
			target.NewDetachFromTargetArgs().SetSessionID(s.ID))
		if err := cdp.ErrorCause(err); err == context.DeadlineExceeded {
			return fmt.Errorf("session: detach timed out for session %s", s.ID)
		}
		return errors.Wrapf(err, "session: detach failed for session %s", s.ID)
	}

	s.conn, err = rpcc.DialContext(ctx, "", sessionDetachConn(detach), sessionCodec(s))
	if err != nil {
		return nil, err
	}
	close(s.init)

	return s, nil
}
