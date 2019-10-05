// Code generated by cdpgen. DO NOT EDIT.

// Package systeminfo implements the SystemInfo domain. The SystemInfo domain
// defines methods and events for querying low-level system information.
package systeminfo

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the SystemInfo domain. The SystemInfo domain
// defines methods and events for querying low-level system information.
type domainClient struct {
	conn      *rpcc.Conn
	sessionID string
}

// NewClient returns a client for the SystemInfo domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// NewClient returns a client for the SystemInfo domain with the connection set to conn.
func NewSessionClient(conn *rpcc.Conn, sessionID string) *domainClient {
	return &domainClient{conn: conn, sessionID: sessionID}
}

// GetInfo invokes the SystemInfo method. Returns information about the
// system.
func (d *domainClient) GetInfo(ctx context.Context) (reply *GetInfoReply, err error) {
	reply = new(GetInfoReply)
	err = rpcc.InvokeRPC(ctx, "SystemInfo.getInfo", d.sessionID, nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "SystemInfo", Op: "GetInfo", Err: err}
	}
	return
}

// GetProcessInfo invokes the SystemInfo method. Returns information about all
// running processes.
func (d *domainClient) GetProcessInfo(ctx context.Context) (reply *GetProcessInfoReply, err error) {
	reply = new(GetProcessInfoReply)
	err = rpcc.InvokeRPC(ctx, "SystemInfo.getProcessInfo", d.sessionID, nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "SystemInfo", Op: "GetProcessInfo", Err: err}
	}
	return
}
