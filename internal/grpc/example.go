package grpc

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
	grpccore "github.com/cynx-io/cynx-core/src/grpc"
)

func (s *Server) Health(ctx context.Context, req *core.GenericRequest) (*core.GenericResponse, error) {
	var resp core.GenericResponse
	return grpccore.HandleGrpc(ctx, req, &resp, s.ExampleService.HealthCheck)
}
