package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

func sendResponse(w http.ResponseWriter, batch bool, res ...*response) {
	encoder := json.NewEncoder(w)

	w.Header().Set(contentType, contentTypeJSON)

	if len(res) > 0 {
		if batch {
			encoder.Encode(res)
		} else {
			encoder.Encode(res[0])
		}
	}
}

func successResponse(marshaler *jsonpb.Marshaler, id json.RawMessage, result proto.Message) *response {
	buf := bytes.NewBuffer(make([]byte, 0))

	err := marshaler.Marshal(buf, result)
	if err != nil {
		return &response{
			ID:      id,
			Jsonprc: jsonrpcVersion,
			Error:   ErrInternalError(err.Error()),
		}
	}

	return &response{
		ID:      id,
		Jsonprc: jsonrpcVersion,
		Result:  buf.Bytes(),
	}
}

func errorResponse(id json.RawMessage, err error) *response {
	var structError *Error
	status, ok := status.FromError(err)
	if ok {
		structError = &Error{
			Code:    serverError.Int() - int(status.Code()),
			Message: serverError.String(),
			Data:    fmt.Sprintf("%s: %s", status.Code().String(), status.Message()),
		}
	}

	if structError == nil {
		structError, ok = err.(*Error)
		if !ok {
			structError = ErrInternalError(err.Error())
		}
	}

	return &response{
		ID:      id,
		Jsonprc: jsonrpcVersion,
		Error:   structError,
	}
}
