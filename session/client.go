package session

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// Client establishes session connections to targets.
type Client struct {
	ctx    context.Context
	cancel context.CancelFunc

	c    *cdp.Client
	sC   chan *session
	done chan error
}

const (
	// This timeout is used for invoking DetachFromTarget when a
	// session connection is closed. Meaning we will wait until
	// timeout for a confirmation response. If a response is not
	// received in this time, it's still possible that the session
	// will eventually close.
	//
	// The default timeout is a compromise between not blocking for
	// extended periods of time and allowing slow connections to
	// deliver the message.
	//
	// TODO(mafredri): Should we allow configuring the timeout?
	defaultDetachTimeout = 5 * time.Second
)

// Dial establishes a target session and creates a lightweight rpcc.Conn
// that uses SendMessageToTarget and ReceivedMessageFromTarget from the
// Target domain instead of a new websocket connection.
//
// Dial will invoke AttachToTarget. Close (rpcc.Conn) will invoke
// DetachFromTarget.
func (sc *Client) Dial(ctx context.Context, id target.ID) (*rpcc.Conn, error) {
	s, err := dial(ctx, id, sc.c, defaultDetachTimeout)
	if err != nil {
		return nil, err
	}
	select {
	case sc.sC <- s:
	case <-sc.ctx.Done():
		s.Close()
		return nil, errors.New("session.Client: Dial failed: Client is closed")
	}
	return s.Conn(), nil
}

// Close closes the Client and all active sessions. All rpcc.Conn
// created by Dial will be closed.
func (sc *Client) Close() error {
	sc.cancel()
	if sc.done != nil {
		err := <-sc.done
		if err != nil {
			return wrapf(err, "session.Client: close failed")
		}
	}
	return nil
}

func (sc *Client) watch(ev *sessionEvents, created <-chan *session, done chan<- error) {
	defer ev.Close()

	isClosing := func(err error) bool {
		switch cdp.ErrorCause(err) {
		case rpcc.ErrConnClosing:
			// Cleanup, the underlying connection was closed
			// before the Client and the Client context does
			// not inherit from rpcc.Conn.
			sc.cancel()
		case context.Canceled:
		default:
			return false
		}
		return true
	}

	sessions := make(map[target.SessionID]*session)
	defer func() {
		var errs []error
		for _, ss := range sessions {
			// TODO(mafredri): Speed up by closing sessions concurrently.
			err := ss.Close()
			if err != nil {
				errs = append(errs, err)
			}
		}
		done <- multiError(errs)
		close(done)
	}()

	for {
		select {
		case s := <-created:
			sessions[s.ID] = s

		// Checking detached should be sufficient for monitoring the
		// session. A DetachedFromTarget event is always sent before
		// TargetDestroyed.
		case <-ev.detached.Ready():
			m, err := ev.detached.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				// TODO(maf): Remove logging.
				fmt.Printf("Client.watch: %v\n", err)
			}

			if s, ok := sessions[m.SessionID]; ok {
				delete(sessions, s.ID)
				s.Close()
			}

		case <-ev.message.Ready():
			m, err := ev.message.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				// TODO(maf): Remove logging.
				fmt.Printf("Client.watch: %v\n", err)
			}

			if s, ok := sessions[m.SessionID]; ok {
				// We rely on the implementation of *rpcc.Conn
				// to read this message in a reasonably short
				// amount of time. Blocking here can potentially
				// delay other session messages but that should
				// not happen.
				err = s.Write([]byte(m.Message))
				if err != nil {
					delete(sessions, s.ID)
				}
			}
		}
	}
}

// NewClient creates a new session client.
//
// The cdp.Client will be used to listen to events and invoke commands
// on the Target domain. It will also be used by all rpcc.Conn created
// by Dial.
func NewClient(c *cdp.Client) (*Client, error) {
	sc := &Client{
		c:  c,
		sC: make(chan *session),
	}

	// TODO(mafredri): Inherit the context from rpcc.Conn in cdp.Client.
	// cdp.Client does not yet expose the context, nor rpcc.Conn.
	sc.ctx, sc.cancel = context.WithCancel(context.TODO())

	ev, err := newSessionEvents(sc.ctx, c)
	if err != nil {
		sc.Close()
		return nil, err
	}

	sc.done = make(chan error, 1)
	go sc.watch(ev, sc.sC, sc.done)
	return sc, nil
}

type sessionEvents struct {
	detached target.DetachedFromTargetClient
	message  target.ReceivedMessageFromTargetClient
}

func newSessionEvents(ctx context.Context, c *cdp.Client) (events *sessionEvents, err error) {
	ev := new(sessionEvents)
	defer func() {
		if err != nil {
			ev.Close()
		}
	}()

	ev.detached, err = c.Target.DetachedFromTarget(ctx)
	if err != nil {
		return nil, err
	}
	ev.message, err = c.Target.ReceivedMessageFromTarget(ctx)
	if err != nil {
		return nil, err
	}

	err = cdp.Sync(ev.detached, ev.message)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (ev *sessionEvents) Close() (err error) {
	for _, c := range []interface {
		Close() error
	}{
		ev.detached,
		ev.message,
	} {
		if c != nil {
			e := c.Close()
			if err == nil {
				err = e
			}
		}
	}
	return err
}
