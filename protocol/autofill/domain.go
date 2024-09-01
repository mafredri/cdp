// Code generated by cdpgen. DO NOT EDIT.

// Package autofill implements the Autofill domain. Defines commands and
// events for Autofill.
package autofill

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Autofill domain. Defines commands and
// events for Autofill.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Autofill domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Trigger invokes the Autofill method. Trigger autofill on a form identified
// by the fieldId. If the field and related form cannot be autofilled, returns
// an error.
func (d *domainClient) Trigger(ctx context.Context, args *TriggerArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Autofill.trigger", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Autofill.trigger", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Autofill", Op: "Trigger", Err: err}
	}
	return
}

// SetAddresses invokes the Autofill method. Set addresses so that developers
// can verify their forms implementation.
func (d *domainClient) SetAddresses(ctx context.Context, args *SetAddressesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Autofill.setAddresses", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Autofill.setAddresses", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Autofill", Op: "SetAddresses", Err: err}
	}
	return
}

// Disable invokes the Autofill method. Disables autofill domain
// notifications.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Autofill.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Autofill", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Autofill method. Enables autofill domain notifications.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Autofill.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Autofill", Op: "Enable", Err: err}
	}
	return
}

func (d *domainClient) AddressFormFilled(ctx context.Context) (AddressFormFilledClient, error) {
	s, err := rpcc.NewStream(ctx, "Autofill.addressFormFilled", d.conn)
	if err != nil {
		return nil, err
	}
	return &addressFormFilledClient{Stream: s}, nil
}

type addressFormFilledClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *addressFormFilledClient) GetStream() rpcc.Stream { return c.Stream }

func (c *addressFormFilledClient) Recv() (*AddressFormFilledReply, error) {
	event := new(AddressFormFilledReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Autofill", Op: "AddressFormFilled Recv", Err: err}
	}
	return event, nil
}