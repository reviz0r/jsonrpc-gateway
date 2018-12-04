package jsonrpc

// Error object
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

const (
	parseError     = "Parse error"
	invalidRequest = "Invalid Request"
	methodNotFound = "Method not found"
	invalidParams  = "Invalid params"
	internalError  = "Internal error"
)

// ErrParseError Invalid JSON was received by the server.
// An error occurred on the server while parsing the JSON text.
func ErrParseError(data interface{}) *Error {
	return &Error{
		Code:    -32700,
		Message: parseError,
		Data:    data,
	}
}

// ErrInvalidRequest The JSON sent is not a valid Request object.
func ErrInvalidRequest(data interface{}) *Error {
	return &Error{
		Code:    -32600,
		Message: invalidRequest,
		Data:    data,
	}
}

// ErrMethodNotFound The method does not exist / is not available.
func ErrMethodNotFound(data interface{}) *Error {
	return &Error{
		Code:    -32601,
		Message: methodNotFound,
		Data:    data,
	}
}

// ErrInvalidParams Invalid method parameter(s).
func ErrInvalidParams(data interface{}) *Error {
	return &Error{
		Code:    -32602,
		Message: invalidParams,
		Data:    data,
	}
}

// ErrInternalError Internal JSON-RPC error.
func ErrInternalError(data interface{}) *Error {
	return &Error{
		Code:    -32603,
		Message: internalError,
		Data:    data,
	}
}
