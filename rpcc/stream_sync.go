package rpcc

import (
	"errors"
	"fmt"
	"sync"
)

// syncMessageStore writes one message into a
// messageWriter and waits for message.next
// to be called before loading the next.
type syncMessageStore struct {
	mu      sync.Mutex
	writers map[string]streamWriter
	backlog []*message
	pending bool
}

func newSyncMessageStore() *syncMessageStore {
	return &syncMessageStore{
		writers: make(map[string]streamWriter),
	}
}

func (s *syncMessageStore) register(method string, w streamWriter) {
	s.mu.Lock()
	s.writers[method] = w
	s.mu.Unlock()
}

// write implements messageWriter, the message is stored
// in a messageWriter if there are no pending messages,
// otherwise appended to backlog.
func (s *syncMessageStore) write(m message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m.next = s.load
	if s.pending {
		s.backlog = append(s.backlog, &m)
		return
	}

	s.pending = true
	w, ok := s.writers[m.method]
	if !ok {
		panic("store: bad mojo " + m.method)
	}
	w.write(m)
}

// load writes the next message into a messageWriter,
// resets pending status if backlog is empty.
func (s *syncMessageStore) load() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.backlog) == 0 {
		s.pending = false
		return
	}

	m := s.backlog[0]
	w, ok := s.writers[m.method]
	if !ok {
		panic("load: bad mojo" + m.method)
	}
	w.write(*m)
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
