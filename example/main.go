package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/reviz0r/jsonrpc-gateway/jsonrpc"

	"google.golang.org/grpc"

	"github.com/reviz0r/jsonrpc-gateway/example/service"
)

var (
	grpcPort = "50051"
	jrpcPort = "8080"
)

func main() {
	if port := os.Getenv("GRPC_PORT"); port != "" {
		grpcPort = port
	}

	if port := os.Getenv("JRPC_PORT"); port != "" {
		jrpcPort = port
	}

	ctx, wg := grace()

	address := fmt.Sprintf(":%s", grpcPort)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen port %s: %s", grpcPort, err)
	}

	server := grpc.NewServer()
	service.RegisterTimeServer(server, new(Service))

	repo := jsonrpc.New()

	repo.RegisterService(
		&service.TimeService{
			Address: fmt.Sprintf("localhost:%s", grpcPort),
			Opts:    []grpc.DialOption{grpc.WithInsecure()},
		},
	)

	mux := http.NewServeMux()
	mux.Handle("/rpc", repo)

	jsonrpcServer := http.Server{
		Addr:    fmt.Sprintf(":%s", jrpcPort),
		Handler: mux,
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := server.Serve(listen); err != nil {
			log.Fatalf("failed to serve grpc %s", err)
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := jsonrpcServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve jsonrpc %s", err)
		}
	}(wg)

	log.Println("starting")
	<-ctx.Done()
	server.Stop()
	jsonrpcServer.Shutdown(context.Background())

	wg.Wait()
	log.Println("shutdown")
}

func grace() (context.Context, *sync.WaitGroup) {
	var wg = new(sync.WaitGroup)
	var stop = make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(group *sync.WaitGroup, ch chan os.Signal) {
		defer group.Done()
		<-ch
		cancel()
	}(wg, stop)

	return ctx, wg
}
