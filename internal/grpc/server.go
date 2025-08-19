package grpc

import (
	"context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/micro-name/internal/service/exampleservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	//pb.UnimplementedExampleServiceServer
	ExampleService *exampleservice.Service
}

func (s *Server) Start(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	//pb.RegisterExampleServiceServer(server, s)
	reflection.Register(server)

	logger.Info(ctx, "Starting gRPC server on ", address)
	return server.Serve(lis)
}
