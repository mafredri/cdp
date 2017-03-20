package rpcc

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
)

var (
	// ErrStreamClosing indicates that
	ErrStreamClosing = errors.New("rpcc: stream is closing")
)

type messageBuffer struct {
	ch    chan []byte
	mu    sync.Mutex
	queue [][]byte
}

func (b *messageBuffer) store(m []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.ch) == 0 {
		select {
		case b.ch <- m:
			return
		default:
		}
	}
	b.queue = append(b.queue, m)
}

func (b *messageBuffer) load() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.queue) > 0 {
		select {
		case b.ch <- b.queue[0]:
			// Pop from queue and ensure references are freed.
			copied := copy(b.queue, b.queue[1:])
			b.queue[copied] = nil
			b.queue = b.queue[:copied]
		default:
		}
	}
}

func (b *messageBuffer) get() <-chan []byte {
	return b.ch
}

// Stream represents a stream of notifications for a certain method.
type Stream interface {
	RecvMsg(m interface{}) error
	Close() error
}

// NewStream creates a new stream that listens to method notifications from the
// RPC server. This function is called by generated code.
func NewStream(ctx context.Context, method string, conn *Conn) (Stream, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	s := &streamClient{userCtx: ctx, done: make(chan struct{})}
	s.msgBuf.ch = make(chan []byte, 1)
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
	msgBuf messageBuffer

	mu     sync.Mutex // Protects following.
	remove func()     // Unsubscribes from messages.

	done chan struct{} // Protects err.
	err  error
}

func (s *streamClient) RecvMsg(m interface{}) (err error) {
	var data []byte

	select {
	case <-s.ctx.Done():
		// Give precedence for user cancellation.
		select {
		case <-s.userCtx.Done():
			return s.userCtx.Err()
		default:
		}

		// Send all messages before returning error.
		select {
		case data = <-s.msgBuf.get():
		default:
			<-s.done
			return s.err
		}
	case <-s.userCtx.Done():
		return s.userCtx.Err()
	case data = <-s.msgBuf.get():
		// Give precedence for user cancellation.
		select {
		case <-s.userCtx.Done():
			return s.userCtx.Err()
		default:
		}
	}

	// Preload the next message.
	s.msgBuf.load()

	return json.Unmarshal(data, m)
}

// Close closes the stream client.
func (s *streamClient) close(err error) error {
	s.mu.Lock()
	remove := s.remove
	s.remove = nil
	s.mu.Unlock()

	if remove == nil {
		return errors.New("rpcc: stream is already closed")
	}

	if err == nil {
		err = ErrStreamClosing
	}

	// Unsubscribe first to prevent incoming messages.
	remove()
	s.cancel()
	s.err = err
	close(s.done)

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

func (s *streamClients) send(args []byte) {
	s.mu.Lock()
	for _, client := range s.clients {
		client.msgBuf.store(args)
	}
	s.mu.Unlock()
}
