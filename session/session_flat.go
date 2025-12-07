package session

import (
	"context"
	"sync"

	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

// flatSession wraps an rpcc.Conn and tracks whether Chrome already
// detached it, so we don't call DetachFromTarget twice.
type flatSession struct {
	id   target.SessionID
	conn *rpcc.Conn

	mu       sync.Mutex // Protects following.
	detached bool
}

func (s *flatSession) markDetached() {
	s.mu.Lock()
	s.detached = true
	s.mu.Unlock()
}

func (s *flatSession) isDetached() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.detached
}

// newFlatSession creates a flattened session connection to the target.
func newFlatSession(conn *rpcc.Conn, sessionID target.SessionID, detachFn func(context.Context) error) (*flatSession, error) {
	fs := &flatSession{id: sessionID}

	closer := func(ctx context.Context) error {
		if fs.isDetached() {
			return nil
		}
		return detachFn(ctx)
	}

	sconn, err := rpcc.NewSession(conn, string(sessionID), rpcc.WithSessionClose(closer))
	if err != nil {
		return nil, err
	}

	fs.conn = sconn
	return fs, nil
}
