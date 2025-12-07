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

// legacySession represents a legacy (non-flattened) session connection to a
// target. Communication is done via Target.SendMessageToTarget.
type legacySession struct {
	id    target.SessionID
	recvC chan []byte
	send  func([]byte) error

	init chan struct{} // Protect conn from early read.
	conn *rpcc.Conn
}

// Ensure that legacySession implements rpcc.Codec.
var _ rpcc.Codec = (*legacySession)(nil)

// WriteRequest implements rpcc.Codec.
func (s *legacySession) WriteRequest(r *rpcc.Request) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return s.send(data)
}

// ReadResponse implements rpcc.Codec.
func (s *legacySession) ReadResponse(r *rpcc.Response) error {
	<-s.init

	select {
	case m := <-s.recvC:
		return json.Unmarshal(m, r)
	case <-s.conn.Context().Done():
		return s.conn.Context().Err()
	}
}

// Conn returns the underlying *rpcc.Conn that uses legacySession as codec.
func (s *legacySession) Conn() *rpcc.Conn { return s.conn }

// Write forwards a target message to the session connection.
// When write returns an error, the session is closed.
func (s *legacySession) Write(data []byte) error {
	select {
	case s.recvC <- data:
		return nil
	case <-s.conn.Context().Done():
		return s.conn.Context().Err()
	}
}

// Close closes the underlying *rpcc.Conn.
func (s *legacySession) Close() error {
	return s.conn.Close()
}

var (
	// We only handle Close on conn to detach the session. The codec
	// handles the actual transport (Read / Write) in this case.
	legacySessionDetachConn = func(detach func() error) rpcc.DialOption {
		return rpcc.WithDialer(
			func(_ context.Context, _ string) (io.ReadWriteCloser, error) {
				return &closeConn{close: detach}, nil
			},
		)
	}
	legacySessionCodec = func(s *legacySession) rpcc.DialOption {
		return rpcc.WithCodec(func(_ io.ReadWriter) rpcc.Codec {
			return s
		})
	}
)

// newLegacySession creates a lightweight RPC connection to the target.
// Communication is done via the underlying *rpcc.Conn for the provided
// *cdp.Client.
func newLegacySession(ctx context.Context, id target.SessionID, tc *cdp.Client, detachTimeout time.Duration) (s *legacySession, err error) {
	s = &legacySession{
		id:    id,
		recvC: make(chan []byte, 1),
		init:  make(chan struct{}),
		send: func(data []byte) error {
			<-s.init
			// TODO(mafredri): Use async invocation.
			return tc.Target.SendMessageToTarget(s.conn.Context(),
				target.NewSendMessageToTargetArgs(string(data)).
					SetSessionID(s.id))
		},
	}

	detach := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), detachTimeout)
		defer cancel()

		err := tc.Target.DetachFromTarget(ctx,
			target.NewDetachFromTargetArgs().SetSessionID(s.id))
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("session: detach timed out for session %s", s.id)
		}
		return errors.Wrapf(err, "session: detach failed for session %s", s.id)
	}

	s.conn, err = rpcc.DialContext(ctx, "", legacySessionDetachConn(detach), legacySessionCodec(s))
	if err != nil {
		return nil, err
	}
	close(s.init)

	return s, nil
}
