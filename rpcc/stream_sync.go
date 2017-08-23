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
	closers []func()
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

func (s *syncMessageStore) register(method string, w streamWriter, conn *Conn) (unregister func(), err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.writers[method]; ok {
		return nil, fmt.Errorf("%s already registered", method)
	}

	remove, err := conn.listen(method, s)
	if err != nil {
		return nil, err
	}

	s.writers[method] = w

	unreg := func() {
		remove()

		s.mu.Lock()
		delete(s.writers, method)
		s.mu.Unlock()
	}
	s.closers = append(s.closers, unreg)

	return unreg, nil
}

func (s *syncMessageStore) close() {
	s.mu.Lock()
	closers := s.closers
	s.closers = nil
	s.mu.Unlock()

	for _, c := range closers {
		c()
	}
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
	w := s.writers[m.method]
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
	w := s.writers[m.method]
	w.write(*m)
	s.backlog[0] = nil // Remove reference from underlying array.
	s.backlog = s.backlog[1:]
}

// Sync synchronizes two or more Streams. On error there will be no
// changes done to the Streams.
func Sync(s ...Stream) (err error) {
	store := newSyncMessageStore()
	defer func() {
		if err != nil {
			store.close()
		}
	}()

	var conn *Conn
	var swap []func()

	for _, ss := range s {
		sc, ok := ss.(*streamClient)
		if !ok {
			return fmt.Errorf("rpcc: Sync: bad Stream type: %T", ss)
		}
		if conn == nil {
			conn = sc.conn
		}
		if sc.conn != conn {
			return errors.New("rpcc: Sync: all Streams must share same Conn")
		}

		// Grab a lock on remove and keep it
		// for the duration of the sync.
		sc.mu.Lock()
		defer sc.mu.Unlock()

		if sc.remove == nil {
			return errors.New("rpcc: Sync: Stream is closed")
		}

		// Register stream client with store.
		unregister, err := store.register(sc.method, sc, sc.conn)
		if err != nil {
			return errors.New("rpcc: Sync: " + err.Error())
		}

		// Delay listener swap until all Streams
		// have been processed successfully.
		swap = append(swap, func() {
			// Unsubscribe Stream from Conn listen.
			sc.remove()
			sc.remove = unregister

			// Clear stream messages to prevent sync issues.
			sc.mbuf.clear()
		})
	}

	// Perform swap, mutex lock (streamClient.mu) is still active.
	for _, s := range swap {
		s()
	}

	return nil
}
