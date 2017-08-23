package rpcc

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
)

var (
	// ErrStreamClosing indicates that the operation is illegal because
	// the stream is closing and there are no pending messages.
	ErrStreamClosing = errors.New("rpcc: the stream is closing")
)

// message contains the invoked method name, data and next func.
type message struct {
	method string
	data   []byte
	next   func()
}

// messageBuffer is an unbounded channel of message.
type messageBuffer struct {
	c       chan *message
	mu      sync.Mutex // Protects following.
	backlog []*message
}

func newMessageBuffer() *messageBuffer {
	return &messageBuffer{
		c: make(chan *message, 1),
	}
}

// store the message in ch, if empty, otherwise in queue.
func (b *messageBuffer) store(m *message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.backlog) == 0 {
		select {
		case b.c <- m:
			return
		default:
		}
	}
	b.backlog = append(b.backlog, m)
}

// load moves a message from the queue into ch.
func (b *messageBuffer) load() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.backlog) > 0 {
		select {
		case b.c <- b.backlog[0]:
			b.backlog[0] = nil // Remove reference from underlying array.
			b.backlog = b.backlog[1:]
		default:
		}
	}
}

// clear removes all messages from buffer.
func (b *messageBuffer) clear() {
	b.mu.Lock()
	backlog := b.backlog
	b.backlog = nil
	b.mu.Unlock()

	select {
	case m := <-b.c:
		m.next()
	default:
	}
	for _, m := range backlog {
		m.next()
	}
}

func (b *messageBuffer) get() <-chan *message {
	return b.c
}

// Stream represents a stream of notifications for a certain method.
type Stream interface {
	// Ready returns a channel that is closed when a message is
	// ready to be received via RecvMsg. Ready indicates that a call
	// to RecvMsg is non-blocking.
	//
	// Ready must not be called concurrently while relying on the
	// non-blocking behavior of RecvMsg. In this case both
	// goroutines will be competing for the same message and one
	// will block until the next message is available.
	//
	// Calling Close on the Stream will close the Ready channel
	// indefinitely, pending messages may still be received via
	// RecvMsg.
	//
	// Ready is provided for use in select statements.
	Ready() <-chan struct{}
	// RecvMsg unmarshals pending messages onto m. Blocks until the
	// next message is received, context is canceled or stream is
	// closed.
	//
	// When m is a *[]byte the message will not be decoded and the
	// raw bytes are copied into m.
	RecvMsg(m interface{}) error
	// Close closes the stream and no new messages will be received.
	// RecvMsg will return ErrStreamClosing once all pending messages
	// have been received.
	Close() error
}

// NewStream creates a new stream that listens to notifications from the
// RPC server. This function is called by generated code.
func NewStream(ctx context.Context, method string, conn *Conn) (Stream, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	return newStreamClient(ctx, method, conn)
}

func newStreamClient(ctx context.Context, method string, conn *Conn) (*streamClient, error) {
	s := &streamClient{
		conn:   conn,
		method: method,
		ctx:    ctx,
		mbuf:   newMessageBuffer(),
		ready:  make(chan struct{}),
		done:   make(chan struct{}),
	}

	remove, err := conn.listen(method, s)
	if err != nil {
		return nil, err
	}
	s.remove = remove

	go s.watch()

	return s, nil
}

type streamClient struct {
	// Used to sync streams.
	conn   *Conn
	method string

	// User provided context.
	ctx context.Context

	// mbuf stores all incoming messages
	// until they are ready to be received.
	mbuf *messageBuffer

	readyMu     sync.Mutex // Protects following.
	ready       chan struct{}
	seq         uint64
	readyClosed bool

	mu     sync.Mutex // Protects following.
	remove func()     // Unsubscribes from messages.
	done   chan struct{}
	err    error
}

func (s *streamClient) watch() {
	select {
	case <-s.ctx.Done():
		s.close(s.ctx.Err())
	case <-s.conn.ctx.Done():
		s.close(ErrConnClosing)
	case <-s.done:
	}
}

func (s *streamClient) Ready() <-chan struct{} {
	s.readyMu.Lock()
	defer s.readyMu.Unlock()
	return s.ready
}

func (s *streamClient) RecvMsg(m interface{}) (err error) {
	msg, err := s.recv()
	if err != nil {
		return err
	}

	if m, ok := m.(*[]byte); ok {
		*m = append(*m, msg.data...)
		return nil
	}

	return json.Unmarshal(msg.data, m)
}

func (s *streamClient) recv() (m *message, err error) {
	userCancelled := func() bool {
		select {
		case <-s.ctx.Done():
			return true
		default:
			return false
		}
	}

	select {
	case <-s.done:
		// Give precedence for user cancellation.
		if userCancelled() {
			return nil, s.ctx.Err()
		}

		// Send all messages before returning error.
		select {
		case m = <-s.mbuf.get():
		default:
			return nil, s.err
		}
	case m = <-s.mbuf.get():
		// Give precedence for user cancellation.
		if userCancelled() {
			return nil, s.ctx.Err()
		}
	}
	m.next()

	return m, nil
}

func (s *streamClient) write(m message) {
	s.readyMu.Lock()
	defer s.readyMu.Unlock()

	if s.seq == 0 && !s.readyClosed {
		// Close the ready channel
		// until the buffer is empty.
		close(s.ready)
	}

	s.seq++ // Keep track of pending messages.
	seq := s.seq

	next := m.next
	m.next = func() {
		s.readyMu.Lock()
		if s.seq == seq && !s.readyClosed {
			// This was the last message, open a blocking
			// ready-channel and reset pending status.
			s.ready = make(chan struct{})
			s.seq = 0
		}
		s.readyMu.Unlock()

		// Prime the next item (if any).
		s.mbuf.load()

		if next != nil {
			next() // Call the prior next func.
		}
	}

	s.mbuf.store(&m)
}

func (s *streamClient) close(err error) error {
	s.mu.Lock()
	remove := s.remove
	s.remove = nil
	s.mu.Unlock()

	if remove == nil {
		return errors.New("rpcc: the stream is already closed")
	}

	if err == nil {
		err = ErrStreamClosing
	}

	remove()    // Unsubscribe first to prevent incoming messages.
	s.err = err // Set err before cancel as reads are protected by context.
	close(s.done)

	// Unblock the ready channel.
	s.readyMu.Lock()
	s.readyClosed = true
	if s.seq == 0 {
		close(s.ready)
	}
	s.readyMu.Unlock()

	return nil
}

// Close closes the stream client.
func (s *streamClient) Close() error {
	return s.close(nil)
}

type streamWriter interface {
	write(message)
}

// streamClients handles multiple streams and allows the
// same message to be sent to one or more streamSender.
type streamClients struct {
	mu      sync.Mutex
	seq     uint64
	writers map[uint64]streamWriter
}

func newStreamClients() *streamClients {
	return &streamClients{
		writers: make(map[uint64]streamWriter),
	}
}

func (s *streamClients) add(w streamWriter) (seq uint64) {
	s.mu.Lock()
	seq = s.seq
	s.seq++
	s.writers[seq] = w
	s.mu.Unlock()
	return seq
}

func (s *streamClients) remove(seq uint64) {
	s.mu.Lock()
	delete(s.writers, seq)
	s.mu.Unlock()
}

func (s *streamClients) write(method string, args []byte) {
	m := message{method: method, data: args}

	s.mu.Lock()
	for _, w := range s.writers {
		w.write(m)
	}
	s.mu.Unlock()
}
