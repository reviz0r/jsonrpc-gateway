package jsonrpc

import (
	"context"
	"io"

	"github.com/golang/protobuf/proto"
)

// Service Сервис
type Service interface {
	Methods() []string
	Exec(ctx context.Context, method string, in io.Reader) (proto.Message, error)
	JsonrpcService()
}

// Method Метод
type Method func(ctx context.Context, in io.Reader) (proto.Message, error)
