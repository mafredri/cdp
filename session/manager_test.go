package session

import (
	"testing"

	"github.com/mafredri/cdp/internal/errors"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/rpcc"
)

type testEventClient struct {
	w    chan struct{}
	done chan struct{}
}

func (c *testEventClient) next()      { c.done = make(chan struct{}); close(c.w) }
func (c *testEventClient) markReady() { close(c.done) }
func (c *testEventClient) Ready() <-chan struct{} {
	<-c.w
	c.w = make(chan struct{})
	return c.done
}
func (c *testEventClient) Close() error                { return nil }
func (c *testEventClient) RecvMsg(m interface{}) error { panic("not implemented") }

func newTestEventClient() *testEventClient {
	return &testEventClient{w: make(chan struct{})}
}

var _ rpcc.Stream = (*testEventClient)(nil)

type testDetacher struct {
	*testEventClient
	err error
}

func (ev *testDetacher) Recv() (*target.DetachedFromTargetReply, error) {
	return nil, ev.err
}

type testMessenger struct {
	*testEventClient
	err error
}

func (ev *testMessenger) Recv() (*target.ReceivedMessageFromTargetReply, error) {
	return nil, ev.err
}

func TestManager_ErrorsAreSentOnErrChan(t *testing.T) {
	detached := &testDetacher{testEventClient: newTestEventClient()}
	message := &testMessenger{testEventClient: newTestEventClient()}
	ev := &sessionEvents{
		detached: detached,
		message:  message,
	}

	m := Manager{
		cancel:         func() {},
		errC:           make(chan error, 1),
		flatSessionC:   make(chan *flatSession),
		legacySessionC: make(chan *legacySession),
		done:           make(chan error, 1),
	}
	go m.watch(ev)

	message.next()

	detached.next()
	detached.err = errors.New("detach nope")
	detached.markReady()
	err := <-m.Err()
	if !errors.Is(err, detached.err) {
		t.Errorf("got error: %v; want: %v", err, detached.err)
	}
	detached.next()

	message.next()
	message.err = errors.New("message nope")
	message.markReady()
	err = <-m.Err()
	if !errors.Is(err, message.err) {
		t.Errorf("got error: %v; want: %v", err, message.err)
	}
	message.next()

	// Close the watcher goroutine.
	detached.err = rpcc.ErrConnClosing
	detached.markReady()
}
