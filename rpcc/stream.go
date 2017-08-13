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

type streamMsg struct {
	method string
	data   []byte
}

type messageBuffer struct {
	ch    chan *streamMsg
	mu    sync.Mutex
	queue []*streamMsg
	ready func()
}

func newMessageBuffer(ready func()) *messageBuffer {
	if ready == nil {
		ready = func() {}
	}
	return &messageBuffer{
		ch:    make(chan *streamMsg, 1),
		ready: ready,
	}
}

// store the message in ch, if empty, otherwise in queue.
func (b *messageBuffer) store(m *streamMsg) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.queue) == 0 {
		select {
		case b.ch <- m:
			b.ready()
			return
		default:
		}
	}
	b.queue = append(b.queue, m)
}

// load moves a message from the queue into ch.
func (b *messageBuffer) load() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.queue) > 0 {
		select {
		case b.ch <- b.queue[0]:
			b.ready()
			b.queue[0] = nil // Remove reference from underlying array.
			b.queue = b.queue[1:]
		default:
		}
	}
}

func (b *messageBuffer) get() <-chan *streamMsg {
	return b.ch
}

// Stream represents a stream of notifications for a certain method.
type Stream interface {
	// Ready returns a channel that sends a single empty struct when
	// a message is ready to be received via RecvMsg. Ready has no
	// backlog and will at most send once per incoming message.
	//
	// Ready is closed when the stream is closed, pending messages
	// may still be read via RecvMsg.
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

	s := &streamClient{
		userCtx: ctx,
		ready:   make(chan struct{}, 1),
		done:    make(chan struct{}),
	}
	s.msgBuf = newMessageBuffer(s.markReady)
	s.ctx, s.cancel = context.WithCancel(context.Background())

	var err error
	s.remove, err = conn.listen(method, s)
	if err != nil {
		return nil, err
	}

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
	ctx     context.Context
	cancel  context.CancelFunc

	// msgBuf stores all incoming messages
	// until they are ready to be received.
	msgBuf *messageBuffer
	ready  chan struct{}

	mu     sync.Mutex // Protects following.
	remove func()     // Unsubscribes from messages.

	done chan struct{} // Protects err.
	err  error
}

func (s *streamClient) Ready() <-chan struct{} {
	return s.ready
}

func (s *streamClient) markReady() {
	s.mu.Lock()
	if s.remove != nil { // Stream is active.
		select {
		case s.ready <- struct{}{}:
		default:
		}
	}
	s.mu.Unlock()
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

func (s *streamClient) recv() (m *streamMsg, err error) {
	userCancelled := func() bool {
		select {
		case <-s.userCtx.Done():
			return true
		default:
			return false
		}
	}

	// Check cancellation once here to avoid race in select.
	if userCancelled() {
		return m, s.userCtx.Err()
	}

	select {
	case <-s.userCtx.Done():
		return m, s.userCtx.Err()
	case <-s.ctx.Done():
		// Give precedence for user cancellation.
		if userCancelled() {
			return m, s.userCtx.Err()
		}

		// Send all messages before returning error.
		select {
		case m = <-s.msgBuf.get():
		default:
			<-s.done
			return m, s.err
		}
	case m = <-s.msgBuf.get():
		// We could check for userCancelled here,
		// but this message would be lost.
	}

	// Preload the next message.
	s.msgBuf.load()

	return m, nil
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

	// Unsubscribe first to prevent incoming messages.
	remove()
	s.cancel()
	s.err = err
	close(s.done)
	close(s.ready)

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
