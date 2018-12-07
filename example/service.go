package main

import (
	"context"
	"time"

	"github.com/reviz0r/jsonrpc-gateway/example/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service .
type Service struct{}

// Sleep .
func (s *Service) Sleep(ctx context.Context, in *service.Request) (*service.Response, error) {
	duration, err := time.ParseDuration(in.Duration)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "f**k")
	}

	time.Sleep(duration)

	return new(service.Response), nil
}
