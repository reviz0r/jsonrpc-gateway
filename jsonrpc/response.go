package jsonrpc

import "encoding/json"

// response Ответ
type response struct {
	ID      json.RawMessage `json:"id"`
	Jsonprc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}
