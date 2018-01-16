package session

import (
	"context"
	"errors"
	"fmt"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// Client establishes session connections to targets.
type Client struct {
	ctx    context.Context
	cancel context.CancelFunc

	c  *cdp.Client
	sC chan *session
}

// Dial establishes a target session and creates a lightweight rpcc.Conn
// that uses SendMessageToTarget and ReceivedMessageFromTarget from the
// Target domain instead of a new websocket connection.
//
// Dial will invoke AttachToTarget. Close (rpcc.Conn) will invoke
// DetachFromTarget.
func (sc *Client) Dial(ctx context.Context, id target.ID) (*rpcc.Conn, error) {
	s, err := dial(ctx, id, sc.c)
	if err != nil {
		return nil, err
	}
	select {
	case sc.sC <- s:
	case <-sc.ctx.Done():
		s.Close()
		return nil, errors.New("Dial: Client is closed")
	}
	return s.Conn(), nil
}

// Close closes the Client and all active sessions. All rpcc.Conn
// created by Dial will be closed.
func (sc *Client) Close() error {
	sc.cancel()
	return nil
}

func (sc *Client) watch(ev *sessionEvents, sessionCreated <-chan *session) {
	defer ev.Close()

	isClosing := func(err error) bool {
		switch cdp.ErrorCause(err) {
		case rpcc.ErrConnClosing:
			// Cleanup, the underlying connection was closed
			// before the Client and the Client context does
			// not inherit from rpcc.Conn.
			sc.Close()
		case context.Canceled:
		default:
			return false
		}
		return true
	}

	sessions := make(map[target.SessionID]*session)
	defer func() {
		for _, ss := range sessions {
			ss.Close()
		}
	}()

	for {
		select {
		case s := <-sessionCreated:
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
	sc := &Client{c: c, sC: make(chan *session)}

	// TODO(mafredri): Inherit the context from rpcc.Conn in cdp.Client.
	// cdp.Client does not yet expose the context, nor rpcc.Conn.
	sc.ctx, sc.cancel = context.WithCancel(context.TODO())

	ev, err := newSessionEvents(sc.ctx, c)
	if err != nil {
		sc.Close()
		return nil, err
	}

	go sc.watch(ev, sc.sC)
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
