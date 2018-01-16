package session

import (
	"context"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// Client establishes session connections to targets.
type Client struct {
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
	sc.sC <- s
	return s.Conn(), nil
}

// Close closes the Client and all active sessions. All rpcc.Conn
// created by Dial will be closed.
func (sc *Client) Close() error {
	// TODO(maf): Make Close safe to be called multiple times.
	close(sc.sC)
	return nil
}

func (sc *Client) watch(ev *sessionEvents, sessionCreated <-chan *session) {
	defer ev.Close()

	isClosing := func(err error) bool {
		if cdp.ErrorCause(err) == rpcc.ErrConnClosing {
			return true
		}
		return false
	}

	sessions := make(map[target.SessionID]*session)
	for {
		select {
		case s, ok := <-sessionCreated:
			// Check if Client was closed.
			if !ok {
				for _, ss := range sessions {
					ss.Close()
				}
				return
			}
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
				// TODO(maf): Remove panic.
				panic(err)
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
				// TODO(maf): Remove panic.
				panic(err)
			}

			if s, ok := sessions[m.SessionID]; ok {
				// We rely on the implementation of *rpcc.Conn
				// to read this message in a reasonably short
				// amount of time. Blocking here can potentially
				// delay other session messages but that should
				// not happen.
				s.Write([]byte(m.Message))
			}
		}
	}
}

// NewClient creates a new session client. The context is not inherited
// by Client.
//
// The cdp.Client will be used to listen to events and invoke commands
// on the Target domain. It will also be used by all rpcc.Conn created
// by Dial.
func NewClient(ctx context.Context, c *cdp.Client) (*Client, error) {
	sc := &Client{c: c, sC: make(chan *session, 1)}

	ev, err := newSessionEvents(ctx, c)
	if err != nil {
		return nil, err
	}

	go sc.watch(ev, sc.sC)
	return sc, nil
}

type sessionEvents struct {
	detached target.DetachedFromTargetClient
	message  target.ReceivedMessageFromTargetClient
}

func newSessionEvents(ctx context.Context, c *cdp.Client) (ev *sessionEvents, err error) {
	ev = new(sessionEvents)
	defer func() {
		if err != nil {
			ev.Close()
		}
	}()

	ev.detached, err = c.Target.DetachedFromTarget(nil)
	if err != nil {
		return nil, err
	}
	ev.message, err = c.Target.ReceivedMessageFromTarget(nil)
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
	return nil
}
