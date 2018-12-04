package jsonrpc

import (
	"encoding/json"
	"errors"
)

// request Запрос
type request struct {
	ID      json.RawMessage `json:"id,omitempty"`
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// isNotification Уведомление
func (r *request) isNotification() bool {
	return r.ID == nil
}

// isValidVersion Правильная ли версия
func (r *request) isValidVersion() bool {
	return r.Jsonrpc == jsonrpcVersion
}

// isMethodEmpty Пустой ли метод
func (r *request) isMethodEmpty() bool {
	return len(r.Method) == 0
}

// validate Корректный ли запрос
func (r *request) validate() error {
	if !r.isValidVersion() {
		return errors.New("invalid json-rpc version")
	}

	if r.isMethodEmpty() {
		return errors.New("method is empty")
	}

	return nil
}
