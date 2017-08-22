package rpcc

import (
	"errors"
	"fmt"
	"sync"
)

// syncMessageStore loads one message into a
// messageStorer and waits for message.next
// to be called before loading the next.
type syncMessageStore struct {
	mu      sync.Mutex
	buf     map[string]streamSender
	backlog []*streamMsg
	pending bool
}

func newSyncMessageStore() *syncMessageStore {
	return &syncMessageStore{
		buf: make(map[string]streamSender),
	}
}

func (s *syncMessageStore) register(method string, ms streamSender) {
	s.mu.Lock()
	s.buf[method] = ms
	s.mu.Unlock()
}

// send implements messageStorer, the message is stored
// in a messageStorer if there are no pending messages,
// otherwise appended to backlog.
func (s *syncMessageStore) send(m streamMsg) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m.next = s.load
	if s.pending {
		s.backlog = append(s.backlog, &m)
		return
	}

	s.pending = true
	ms, ok := s.buf[m.method]
	if !ok {
		panic("store: bad mojo " + m.method)
	}
	ms.send(m)
}

// load stores the next message into a messageStorer,
// resets pending status if backlog is empty.
func (s *syncMessageStore) load() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.backlog) == 0 {
		s.pending = false
		return
	}

	m := s.backlog[0]
	ms, ok := s.buf[m.method]
	if !ok {
		panic("load: bad mojo" + m.method)
	}
	ms.send(*m)
	s.backlog[0] = nil // Remove reference from underlying array.
	s.backlog = s.backlog[1:]
}

// Sync synchronizes two or more Streams.
func Sync(s ...Stream) error {
	if len(s) == 0 {
		return nil
	}

	store := newSyncMessageStore()

	// Validate that the Streams can be synced, they must belong to
	// the same Conn and the same type of stream cannot be synced
	// more than once.
	var conn *Conn
	prev := make(map[string]struct{})
	for _, ss := range s {
		sc, ok := ss.(*streamClient)
		if !ok {
			return fmt.Errorf("rpcc: cannot sync Stream of type %T", ss)
		}
		if conn == nil {
			conn = sc.conn
		}
		if sc.conn != conn {
			return errors.New("rpcc: cannot sync Stream with different Conn")
		}
		if _, ok := prev[sc.method]; ok {
			return fmt.Errorf("rpcc: cannot sync Stream with method %q twice", sc.method)
		}
		prev[sc.method] = struct{}{}

		// Grab a lock on remove.
		sc.mu.Lock()
		defer sc.mu.Unlock()

		if sc.remove == nil {
			return errors.New("rpcc: cannot sync closed Stream")
		}

		// Unsubscribe Stream from Conn listen.
		sc.remove()

		// Register stream client with store.
		store.register(sc.method, sc)

		// Resubscribe with store as intermediary.
		remove, err := conn.listen(sc.method, store)
		if err != nil {
			return err
		}
		sc.remove = remove
	}

	return nil
}
