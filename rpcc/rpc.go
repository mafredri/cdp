package rpcc

import (
	"encoding/json"
	"fmt"
)

type rpcRequest struct {
	ID     uint64      `json:"id"`               // ID chosen by client.
	Method string      `json:"method"`           // Method invoked on remote.
	Params interface{} `json:"params,omitempty"` // Method parameters, if any.
}

type rpcResponse struct {
	// Response to request.
	ID     uint64          `json:"id"`     // Echoes that of the rpcRequest.
	Result json.RawMessage `json:"result"` // Result from invokation, if any.
	Error  *rpcError       `json:"error"`  // Error, if any.

	// RPC notification from remote.
	Method string          `json:"method"` // Method invokation requested by remote.
	Params json.RawMessage `json:"params"` // Method parameters, if any.
}

func (r *rpcResponse) Reset() {
	r.ID = 0
	r.Result = nil
	r.Error = nil
	r.Method = ""
	r.Params = nil
}

func (r *rpcResponse) String() string {
	if r.Method != "" {
		return fmt.Sprintf("Method = %s, Params = %s", r.Method, r.Params)
	}
	if r.Error != nil {
		return fmt.Sprintf("ID = %d, Error = %s", r.ID, r.Error.Error())
	}
	return fmt.Sprintf("ID = %d, Result = %s", r.ID, r.Result)
}

type rpcError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *rpcError) Error() string {
	var data string
	if e.Data != "" {
		data = ", data = " + e.Data
	}
	return fmt.Sprintf("rpc error: %s (code = %d%s)", e.Message, e.Code, data)
}

var (
	_ error = (*rpcError)(nil)
)
