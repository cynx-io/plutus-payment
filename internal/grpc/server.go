package grpc

import (
	"context"
	"github.com/cynx-io/cynx-core/src/logger"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/app"
	"github.com/cynx-io/plutus-payment/internal/dependencies/config"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type Server struct {
	proto.UnimplementedPaymentServiceServer
	proto.UnimplementedWebhookXenditServiceServer

	Services *app.Services
}

func (s *Server) Start(ctx context.Context) error {

	var g errgroup.Group
	address := config.Config.App.Address + ":" + strconv.Itoa(config.Config.App.Port)

	g.Go(func() error {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		proto.RegisterPaymentServiceServer(server, s)
		proto.RegisterWebhookXenditServiceServer(server, s)
		reflection.Register(server)

		logger.Info(ctx, "Starting gRPC server on ", address)
		return server.Serve(lis)
	})
	return g.Wait()
}
