package session

import (
	"context"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/internal/errors"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// Manager establishes session connections to targets.
type Manager struct {
	ctx    context.Context
	cancel context.CancelFunc

	c    *cdp.Client
	sC   chan *session
	done chan error
	errC chan error
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
func (m *Manager) Dial(ctx context.Context, id target.ID) (*rpcc.Conn, error) {
	s, err := dial(ctx, id, m.c, defaultDetachTimeout)
	if err != nil {
		return nil, err
	}
	select {
	case m.sC <- s:
	case <-m.ctx.Done():
		s.Close()
		return nil, errors.New("session.Manager: Dial failed: Manager is closed")
	}
	return s.Conn(), nil
}

// Close closes the Manager and all active sessions. All rpcc.Conn
// created by Dial will be closed.
func (m *Manager) Close() error {
	m.cancel()
	if m.done != nil {
		errors.Wrapf(<-m.done, "session.Manager: close failed")
	}
	return nil
}

func (m *Manager) watch(ev *sessionEvents, created <-chan *session, done, errC chan<- error) {
	defer ev.Close()

	isClosing := func(err error) bool {
		switch cdp.ErrorCause(err) {
		case rpcc.ErrConnClosing:
			// Cleanup, the underlying connection was closed
			// before the Manager and its context does not
			// inherit from rpcc.Conn.
			m.cancel()
		case context.Canceled:
		default:
			return false
		}
		return true
	}

	sendErr := func(err error) {
		select {
		case errC <- err:
		default:
		}
	}

	sessions := make(map[target.SessionID]*session)
	defer func() {
		var err []error
		for _, ss := range sessions {
			// TODO(mafredri): Speed up by closing sessions concurrently.
			err = append(err, ss.Close())
		}
		done <- errors.Merge(err...)
		close(done)
		close(errC)
	}()

	for {
		select {
		case s := <-created:
			sessions[s.ID] = s

		// Checking detached should be sufficient for monitoring the
		// session. A DetachedFromTarget event is always sent before
		// TargetDestroyed.
		case <-ev.detached.Ready():
			ev, err := ev.detached.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				sendErr(errors.Wrapf(err, "Manager.watch: error receiving detached event"))
				continue
			}

			if s, ok := sessions[ev.SessionID]; ok {
				delete(sessions, s.ID)
				s.Close()
			}

		case <-ev.message.Ready():
			ev, err := ev.message.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				sendErr(errors.Wrapf(err, "Manager.watch: error receiving message event"))
				continue
			}

			if s, ok := sessions[ev.SessionID]; ok {
				// We rely on the implementation of *rpcc.Conn
				// to read this message in a reasonably short
				// amount of time. Blocking here can potentially
				// delay other session messages but that should
				// not happen.
				err = s.Write([]byte(ev.Message))
				if err != nil {
					delete(sessions, s.ID)
				}
			}
		}
	}
}

// Err is a channel that blocks until the Manager encounters an error.
// The channel is closed if Manager is closed.
//
// Errors could happen if the debug target sends events that cannot be
// decoded from JSON.
func (m *Manager) Err() <-chan error {
	return m.errC
}

// NewManager creates a new session Manager.
//
// The cdp.Client will be used to listen to events and invoke commands
// on the Target domain. It will also be used by all rpcc.Conn created
// by Dial.
func NewManager(c *cdp.Client) (*Manager, error) {
	m := &Manager{
		c:    c,
		sC:   make(chan *session),
		errC: make(chan error, 1),
	}

	// TODO(mafredri): Inherit the context from rpcc.Conn in cdp.Client.
	// cdp.Client does not yet expose the context, nor rpcc.Conn.
	m.ctx, m.cancel = context.WithCancel(context.TODO())

	ev, err := newSessionEvents(m.ctx, c)
	if err != nil {
		close(m.errC)
		m.Close()
		return nil, err
	}

	m.done = make(chan error, 1)
	go m.watch(ev, m.sC, m.done, m.errC)
	return m, nil
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
