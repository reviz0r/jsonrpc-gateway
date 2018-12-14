package main

import (
	"context"
	"fmt"
	"time"

	"github.com/reviz0r/jsonrpc-gateway/example/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service .
type Service struct{}

// Now .
func (s *Service) Now(ctx context.Context, in *service.NowRequest) (*service.NowResponse, error) {
	const timestamp = "2006-01-02T15:04:05.000Z07:00"
	now := time.Now()

	if in.GetLocation() != "" {
		location, err := time.LoadLocation(in.GetLocation())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid location")
		}

		now = now.In(location)
		return &service.NowResponse{Now: now.Format(timestamp), Location: location.String()}, nil
	}

	return &service.NowResponse{Now: now.Format(timestamp)}, nil
}

// Sleep .
func (s *Service) Sleep(ctx context.Context, in *service.SleepRequest) (*service.SleepResponse, error) {
	duration, err := time.ParseDuration(in.GetDuration())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid duration")
	}

	time.Sleep(duration)

	return &service.SleepResponse{Result: fmt.Sprintf("Sleep duration: %s", duration.String())}, nil
}
