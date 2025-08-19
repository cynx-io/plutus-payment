package exampleservice

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/micro-name/internal/model/response"
)

func (s *Service) HealthCheck(ctx context.Context, req *core.GenericRequest, resp *core.GenericResponse) error {
	response.Success(resp)
	return nil
}
