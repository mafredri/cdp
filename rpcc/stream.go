package rpcc

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
)

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

	s := &streamClient{
		ch: make(chan []byte, 1),
	}
	var err error
	s.close, err = conn.listen(method, s.ch)
	if err != nil {
		return nil, err
	}

	go func() {
		var err error
		select {
		case <-conn.ctx.Done():
			err = ErrConnClosing
		case <-ctx.Done():
			err = ctx.Err()
		}
		s.setErr(err)
		s.Close()
	}()

	return s, nil
}

type streamClient struct {
	ch chan []byte

	mu     sync.Mutex // Protects following.
	close  func()
	closed bool
	err    error
}

func (s *streamClient) RecvMsg(m interface{}) error {
	s.mu.Lock()
	err := s.err
	s.mu.Unlock()
	if err != nil {
		return err
	}

	select {
	case data, ok := <-s.ch:
		// Stream errors are fatal, no messages will be
		// processed once an error is encountered.
		s.mu.Lock()
		err = s.err
		s.mu.Unlock()
		if err != nil {
			return err
		}

		if !ok {
			return errors.New("rpcc: empty response on stream channel")
		}
		return json.Unmarshal(data, m)
	}
}

// Close closes the stream client.
func (s *streamClient) Close() error {
	s.mu.Lock()
	closed := s.closed
	s.closed = true
	s.mu.Unlock()
	if closed {
		return errors.New("rpcc: stream already closed")
	}
	s.setErr(errors.New("rpcc: stream is closed"))
	s.close()

	// At this point the channel has been removed from
	// the stream service, and is safe to close.
	close(s.ch)

	return nil
}

// setErr only sets the stream error once.
func (s *streamClient) setErr(err error) {
	s.mu.Lock()
	if s.err == nil {
		s.err = err
	}
	s.mu.Unlock()
}

// streamService manages stream subscribers, the service enables
// registering multiple subscribers to a CDP event.
type streamService struct {
	mu    sync.RWMutex
	seq   uint64
	chans map[uint64]chan<- []byte
}

func newStreamService() *streamService {
	return &streamService{chans: make(map[uint64]chan<- []byte)}
}

func (s *streamService) add(ch chan<- []byte) (seq uint64) {
	s.mu.Lock()
	seq = s.seq
	s.seq++
	s.chans[seq] = ch
	s.mu.Unlock()
	return seq
}

func (s *streamService) remove(seq uint64) {
	s.mu.Lock()
	delete(s.chans, seq)
	s.mu.Unlock()
}

// send is called in jsonrpc2 and transmits
// on all active channels for the stream.
func (s *streamService) send(args []byte) {
	s.mu.RLock()
	for _, ch := range s.chans {
		// We cannot block here since the user decides
		// when stream messages are received.
		go func(ch chan<- []byte) { ch <- args }(ch)
	}
	s.mu.RUnlock()
}
