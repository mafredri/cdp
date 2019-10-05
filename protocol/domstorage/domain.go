// Code generated by cdpgen. DO NOT EDIT.

// Package domstorage implements the DOMStorage domain. Query and modify DOM
// storage.
package domstorage

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the DOMStorage domain. Query and modify DOM
// storage.
type domainClient struct {
	conn      *rpcc.Conn
	sessionID string
}

// NewClient returns a client for the DOMStorage domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// NewClient returns a client for the DOMStorage domain with the connection set to conn.
func NewSessionClient(conn *rpcc.Conn, sessionID string) *domainClient {
	return &domainClient{conn: conn, sessionID: sessionID}
}

// Clear invokes the DOMStorage method.
func (d *domainClient) Clear(ctx context.Context, args *ClearArgs) (err error) {
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.clear", d.sessionID, args, nil, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.clear", d.sessionID, nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "Clear", Err: err}
	}
	return
}

// Disable invokes the DOMStorage method. Disables storage tracking, prevents
// storage events from being sent to the client.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "DOMStorage.disable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the DOMStorage method. Enables storage tracking, storage
// events will now be delivered to the client.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "DOMStorage.enable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "Enable", Err: err}
	}
	return
}

// GetDOMStorageItems invokes the DOMStorage method.
func (d *domainClient) GetDOMStorageItems(ctx context.Context, args *GetDOMStorageItemsArgs) (reply *GetDOMStorageItemsReply, err error) {
	reply = new(GetDOMStorageItemsReply)
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.getDOMStorageItems", d.sessionID, args, reply, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.getDOMStorageItems", d.sessionID, nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "GetDOMStorageItems", Err: err}
	}
	return
}

// RemoveDOMStorageItem invokes the DOMStorage method.
func (d *domainClient) RemoveDOMStorageItem(ctx context.Context, args *RemoveDOMStorageItemArgs) (err error) {
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.removeDOMStorageItem", d.sessionID, args, nil, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.removeDOMStorageItem", d.sessionID, nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "RemoveDOMStorageItem", Err: err}
	}
	return
}

// SetDOMStorageItem invokes the DOMStorage method.
func (d *domainClient) SetDOMStorageItem(ctx context.Context, args *SetDOMStorageItemArgs) (err error) {
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.setDOMStorageItem", d.sessionID, args, nil, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMStorage.setDOMStorageItem", d.sessionID, nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMStorage", Op: "SetDOMStorageItem", Err: err}
	}
	return
}

func (d *domainClient) DOMStorageItemAdded(ctx context.Context) (ItemAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOMStorage.domStorageItemAdded", d.sessionID, d.conn)
	if err != nil {
		return nil, err
	}
	return &itemAddedClient{Stream: s}, nil
}

type itemAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *itemAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *itemAddedClient) Recv() (*ItemAddedReply, error) {
	event := new(ItemAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOMStorage", Op: "DOMStorageItemAdded Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DOMStorageItemRemoved(ctx context.Context) (ItemRemovedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOMStorage.domStorageItemRemoved", d.sessionID, d.conn)
	if err != nil {
		return nil, err
	}
	return &itemRemovedClient{Stream: s}, nil
}

type itemRemovedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *itemRemovedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *itemRemovedClient) Recv() (*ItemRemovedReply, error) {
	event := new(ItemRemovedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOMStorage", Op: "DOMStorageItemRemoved Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DOMStorageItemUpdated(ctx context.Context) (ItemUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOMStorage.domStorageItemUpdated", d.sessionID, d.conn)
	if err != nil {
		return nil, err
	}
	return &itemUpdatedClient{Stream: s}, nil
}

type itemUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *itemUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *itemUpdatedClient) Recv() (*ItemUpdatedReply, error) {
	event := new(ItemUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOMStorage", Op: "DOMStorageItemUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DOMStorageItemsCleared(ctx context.Context) (ItemsClearedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOMStorage.domStorageItemsCleared", d.sessionID, d.conn)
	if err != nil {
		return nil, err
	}
	return &itemsClearedClient{Stream: s}, nil
}

type itemsClearedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *itemsClearedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *itemsClearedClient) Recv() (*ItemsClearedReply, error) {
	event := new(ItemsClearedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOMStorage", Op: "DOMStorageItemsCleared Recv", Err: err}
	}
	return event, nil
}
