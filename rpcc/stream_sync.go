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
	conn    *Conn // Used as validation.
	writers map[string]streamWriter
	backlog []*message
	pending bool
	closers []func()
}

func newSyncMessageStore() *syncMessageStore {
	return &syncMessageStore{
		writers: make(map[string]streamWriter),
	}
}

func (s *syncMessageStore) subscribe(method string, w streamWriter, conn *Conn) (unsubscribe func(), err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.conn == nil {
		s.conn = conn
	} else if conn != s.conn {
		return nil, fmt.Errorf("rpcc: same Conn must be used")
	}

	if _, ok := s.writers[method]; ok {
		return nil, fmt.Errorf("%s already subscribed", method)
	}

	remove, err := conn.listen(method, s)
	if err != nil {
		return nil, err
	}

	s.writers[method] = w

	unsub := func() {
		remove()

		s.mu.Lock()
		delete(s.writers, method)
		if len(s.writers) == 0 {
			// Either close has been called
			// or all streams have closed.
			s.writers = nil
			s.backlog = nil
			s.closers = nil
		}
		s.mu.Unlock()
	}
	s.closers = append(s.closers, unsub)

	return unsub, nil
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

// Sync takes two or more streams and sets them into synchronous operation,
// relative to each other. This operation cannot be undone. If an error is
// returned this function is no-op and the streams will continue in asynchronous
// operation.
//
// All streams must belong to the same Conn and they must not be closed. Passing
// multiple streams of the same method to Sync is not supported and will return
// an error.
//
// A stream that is closed is removed and has no further affect on the streams
// that were synchronized.
//
// When two streams, A and B, are in sync they will both receive messages in the
// order that they arrived on Conn. If a message for both A and B arrives, in
// that order, it will not be possible to receive the message from B before the
// message from A has been received.
func Sync(s ...Stream) (err error) {
	if len(s) < 2 {
		return errors.New("rpcc: Sync: two or more streams must be provided")
	}

	store := newSyncMessageStore()
	var swap []func(bool) func()

	defer func() {
		// Perform swap, mutex lock (streamClient.mu) is still active.
		for _, s := range swap {
			defer s(err == nil)()
		}
		if err != nil {
			store.close()
		}
	}()

	for _, ss := range s {
		swapFn, err := ss.Sync(store)
		if err != nil {
			return err
		}
		swap = append(swap, swapFn)
	}

	return nil
}

func (s *streamClient) Sync(storer interface{}) (activate func(bool) func(), err error) {
	// The Stream lock must be held until the
	// swap has been done for all streams.
	s.mu.Lock()
	defer func() {
		if err != nil {
			s.mu.Unlock()
		}
	}()

	store, ok := storer.(*syncMessageStore)
	if !ok {
		return nil, fmt.Errorf("streamClient: Sync: bad store %T must be of type *syncMessageStore", storer)
	}

	if s.remove == nil {
		return nil, errors.New("rpcc: Sync: Stream is closed")
	}

	// Allow store to manage messages to streamClient.
	unsub, err := store.subscribe(s.method, s, s.conn)
	if err != nil {
		return nil, errors.New("rpcc: Sync: " + err.Error())
	}

	// Delay listener swap until all Streams have been
	// processed so that we can abort on error.
	return func(ok bool) func() {
		if ok {
			s.remove()       // Prevent direct events from Conn.
			s.remove = unsub // Remove from store on Close.

			// Clear stream messages to prevent sync issues.
			s.mbuf.clear()
		}
		return func() { s.mu.Unlock() }
	}, nil
}
