package session

import (
	"context"
	"fmt"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/internal/errors"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

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
	defaultDetachTimeout = 5 * time.Second
)

// Manager establishes session connections to targets. It handles both
// flattened sessions (Chrome 77+) and legacy sessions.
type Manager struct {
	ctx    context.Context
	cancel context.CancelFunc

	c       *cdp.Client
	conn    *rpcc.Conn
	flatten bool

	flatSessionC   chan *flatSession
	legacySessionC chan *legacySession
	done           chan error
	errC           chan error
}

// ManagerOption represents an option for NewManager.
type ManagerOption func(*managerOptions)

type managerOptions struct {
	flatten bool
}

// WithNoFlatten returns a ManagerOption that disables flattened session
// mode. Use this for Chrome versions older than 77.
func WithNoFlatten() ManagerOption {
	return func(o *managerOptions) {
		o.flatten = false
	}
}

// NewManager creates a new session Manager.
//
// The cdp.Client will be used to listen to events and invoke commands
// on the Target domain. It will also be used by all rpcc.Conn created
// by Dial.
func NewManager(c *cdp.Client, opts ...ManagerOption) (*Manager, error) {
	conn := c.Conn()
	if conn.SessionID() != "" {
		return nil, errors.New("session.Manager: cannot use session connection, must use parent connection")
	}

	opt := managerOptions{flatten: true}
	for _, o := range opts {
		o(&opt)
	}

	m := &Manager{
		c:              c,
		conn:           conn,
		flatten:        opt.flatten,
		flatSessionC:   make(chan *flatSession),
		legacySessionC: make(chan *legacySession),
		errC:           make(chan error, 1),
	}

	// Inherit context from the parent connection.
	m.ctx, m.cancel = context.WithCancel(m.conn.Context())

	ev, err := newSessionEvents(m.ctx, c)
	if err != nil {
		close(m.errC)
		m.Close()
		return nil, err
	}

	m.done = make(chan error, 1)
	go m.watch(ev)
	return m, nil
}

// Dial establishes a target session and creates a lightweight rpcc.Conn.
//
// Dial invokes Target.AttachToTarget, closing the returned rpcc.Conn
// invokes Target.DetachFromTarget.
func (m *Manager) Dial(ctx context.Context, id target.ID) (*rpcc.Conn, error) {
	args := target.NewAttachToTargetArgs(id).SetFlatten(m.flatten)
	reply, err := m.c.Target.AttachToTarget(ctx, args)
	if err != nil {
		if m.flatten {
			return nil, fmt.Errorf("session.Manager: attach to target failed (flatten=true requires Chrome 77+, use WithNoFlatten for older versions): %w", err)
		}
		return nil, err
	}

	if m.flatten {
		return m.dialFlat(reply.SessionID)
	}
	return m.dialLegacy(ctx, reply.SessionID)
}

// Close closes the Manager and all active sessions. All rpcc.Conn
// created by Dial will be closed.
func (m *Manager) Close() error {
	m.cancel()

	// Wait for watch to close all sessions and finish.
	if m.done != nil {
		return errors.Wrapf(<-m.done, "session.Manager: close failed")
	}
	return nil
}

// Err is a channel that blocks until the Manager encounters an error.
// The channel is closed if Manager is closed.
//
// Errors could happen if the debug target sends events that cannot be
// decoded from JSON.
func (m *Manager) Err() <-chan error {
	return m.errC
}

func (m *Manager) dialFlat(sessionID target.SessionID) (*rpcc.Conn, error) {
	detach := func(ctx context.Context) error {
		err := m.c.Target.DetachFromTarget(ctx, target.NewDetachFromTargetArgs().SetSessionID(sessionID))
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("session.Manager: detach timed out for session %s", sessionID)
		}
		return err
	}

	fs, err := newFlatSession(m.conn, sessionID, detach)
	if err != nil {
		return nil, err
	}

	select {
	case m.flatSessionC <- fs:
	case <-m.ctx.Done():
		fs.conn.Close()
		return nil, errors.New("session.Manager: Dial failed: Manager is closed")
	}

	return fs.conn, nil
}

func (m *Manager) dialLegacy(ctx context.Context, sessionID target.SessionID) (*rpcc.Conn, error) {
	s, err := newLegacySession(ctx, sessionID, m.c, defaultDetachTimeout)
	if err != nil {
		return nil, err
	}
	select {
	case m.legacySessionC <- s:
	case <-m.ctx.Done():
		s.Close()
		return nil, errors.New("session.Manager: Dial failed: Manager is closed")
	}
	return s.Conn(), nil
}

func (m *Manager) watch(ev *sessionEvents) {
	defer ev.Close()

	isClosing := func(err error) bool {
		// Test if this is an rpcc.closeError.
		var e interface{ Closed() bool }
		if ok := errors.As(err, &e); ok && e.Closed() {
			// Cleanup, the underlying connection was closed
			// before the Manager and its context does not
			// inherit from rpcc.Conn.
			m.cancel()
			return true
		}

		if errors.Is(err, context.Canceled) {
			// Manager was closed.
			return true
		}
		return false
	}

	flatSessions := make(map[target.SessionID]*flatSession)
	legacySessions := make(map[target.SessionID]*legacySession)
	defer func() {
		var errs []error
		for _, fs := range flatSessions {
			errs = append(errs, fs.conn.Close())
		}
		for _, ls := range legacySessions {
			errs = append(errs, ls.Close())
		}
		m.done <- errors.Merge(errs...)
		close(m.done)
		close(m.errC)
	}()

	for {
		select {
		case fs := <-m.flatSessionC:
			flatSessions[fs.id] = fs

		case ls := <-m.legacySessionC:
			legacySessions[ls.id] = ls

		// Checking detached should be sufficient for monitoring the
		// session. A DetachedFromTarget event is always sent before
		// TargetDestroyed.
		case <-ev.detached.Ready():
			ev, err := ev.detached.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				err = errors.Wrapf(err, "Manager.watch: error receiving detached event")
				sendOrDiscardErr(m.errC, err)
				continue
			}

			// Handle flattened session detachment.
			if fs, ok := flatSessions[ev.SessionID]; ok {
				delete(flatSessions, ev.SessionID)
				fs.markDetached()
				fs.conn.Close()
			}

			// Handle legacy session detachment.
			if ls, ok := legacySessions[ev.SessionID]; ok {
				delete(legacySessions, ls.id)
				ls.Close()
			}

		case <-ev.message.Ready():
			ev, err := ev.message.Recv()
			if err != nil {
				if isClosing(err) {
					return
				}
				err = errors.Wrapf(err, "Manager.watch: error receiving message event")
				sendOrDiscardErr(m.errC, err)
				continue
			}

			// Forward messages to legacy sessions only. Flattened
			// sessions receive messages directly via rpcc.Conn.
			if ls, ok := legacySessions[ev.SessionID]; ok {
				// We rely on the implementation of *rpcc.Conn
				// to read this message in a reasonably short
				// amount of time. Blocking here can potentially
				// delay other session messages but that should
				// not happen.
				err = ls.Write([]byte(ev.Message))
				if err != nil {
					delete(legacySessions, ls.id)
				}
			}
		}
	}
}

func sendOrDiscardErr(errC chan<- error, err error) {
	select {
	case errC <- err:
	default:
	}
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
