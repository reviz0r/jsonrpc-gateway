package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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
	structError, ok := err.(*Error)
	if !ok {
		structError = ErrInternalError(nil)
	}

	return &response{
		ID:      id,
		Jsonprc: jsonrpcVersion,
		Error:   structError,
	}
}
