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

type bufItem struct {
	m    *streamMsg
	next func()
}

func (bi *bufItem) message() *streamMsg {
	bi.next()
	return bi.m
}

type messageBuffer struct {
	c      chan *bufItem
	mu     sync.Mutex // Protects following.
	seq    uint
	queue  []*bufItem
	rc     chan struct{}
	closed bool
}

func newMessageBuffer() *messageBuffer {
	return &messageBuffer{
		c:  make(chan *bufItem, 1),
		rc: make(chan struct{}),
	}
}

// store the message in ch, if empty, otherwise in queue.
func (b *messageBuffer) store(m *streamMsg) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.seq == 0 && !b.closed {
		// Close the ready channel
		// until the buffer is empty.
		close(b.rc)
	}

	b.seq++ // Keep track of pending messages.
	seq := b.seq

	// nextItem will be called when this
	// message is fetched from the buffer.
	nextItem := func() {
		b.mu.Lock()
		if b.seq == seq && !b.closed {
			// This was the last message, open a blocking
			// ready-channel and reset pending status.
			b.rc = make(chan struct{})
			b.seq = 0
		}
		b.mu.Unlock()

		// Prime the next item (if any).
		b.load()
	}
	bi := &bufItem{m: m, next: nextItem}

	if len(b.queue) == 0 {
		select {
		case b.c <- bi:
			return
		default:
		}
	}
	b.queue = append(b.queue, bi)
}

// load moves a message from the queue into ch.
func (b *messageBuffer) load() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.queue) > 0 {
		select {
		case b.c <- b.queue[0]:
			b.queue[0] = nil // Remove reference from underlying array.
			b.queue = b.queue[1:]
		default:
		}
	}
}

// ready returns a channel that is closed when the buffer is non-empty.
func (b *messageBuffer) ready() <-chan struct{} {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.rc
}

func (b *messageBuffer) get() <-chan *bufItem {
	return b.c
}

// close ensures the ready channel is closed.
func (b *messageBuffer) close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.closed = true
	if b.seq == 0 {
		close(b.rc)
	}
}

// streamMsg contains the invoked method name and data.
type streamMsg struct {
	method string
	data   []byte
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

	s := new(streamClient)
	s.userCtx = ctx
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.msgBuf = newMessageBuffer()

	remove, err := conn.listen(method, s)
	if err != nil {
		return nil, err
	}
	s.remove = remove

	go func() {
		select {
		case <-s.ctx.Done():
		case <-conn.ctx.Done():
			s.close(ErrConnClosing)
		case <-ctx.Done():
			s.close(ctx.Err())
		}
	}()

	return s, nil
}

type streamClient struct {
	userCtx context.Context
	cancel  context.CancelFunc
	ctx     context.Context // Protects following.
	err     error

	// msgBuf stores all incoming messages
	// until they are ready to be received.
	msgBuf *messageBuffer

	mu     sync.Mutex // Protects following.
	remove func()     // Unsubscribes from messages.
}

func (s *streamClient) Ready() <-chan struct{} {
	return s.msgBuf.ready()
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

func (s *streamClient) recv() (*streamMsg, error) {
	userCancelled := func() bool {
		select {
		case <-s.userCtx.Done():
			return true
		default:
			return false
		}
	}

	var bi *bufItem
	select {
	case <-s.ctx.Done():
		// Give precedence for user cancellation.
		if userCancelled() {
			return nil, s.userCtx.Err()
		}

		// Send all messages before returning error.
		select {
		case bi = <-s.msgBuf.get():
		default:
			return nil, s.err
		}
	case bi = <-s.msgBuf.get():
		// Give precedence for user cancellation.
		if userCancelled() {
			return nil, s.userCtx.Err()
		}
	}

	return bi.message(), nil
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
	s.cancel()

	// Unblock the ready channel.
	s.msgBuf.close()

	return nil
}

// Close closes the stream client.
func (s *streamClient) Close() error {
	return s.close(nil)
}

// streamClients handles multiple instances of streamClient and
// enables sending of the same message to multiple clients.
type streamClients struct {
	mu      sync.Mutex
	seq     uint64
	clients map[uint64]*streamClient
}

func newStreamService() *streamClients {
	return &streamClients{
		clients: make(map[uint64]*streamClient),
	}
}

func (s *streamClients) add(client *streamClient) (seq uint64) {
	s.mu.Lock()
	seq = s.seq
	s.seq++
	s.clients[seq] = client
	s.mu.Unlock()
	return seq
}

func (s *streamClients) remove(seq uint64) {
	s.mu.Lock()
	delete(s.clients, seq)
	s.mu.Unlock()
}

func (s *streamClients) send(method string, args []byte) {
	m := &streamMsg{method: method, data: args}

	s.mu.Lock()
	for _, client := range s.clients {
		client.msgBuf.store(m)
	}
	s.mu.Unlock()
}
