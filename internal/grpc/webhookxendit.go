package grpc

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/grpc"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

func (s *Server) HandlePaymentInvoice(ctx context.Context, req *proto.HandlePaymentInvoiceRequest) (*core.GenericResponse, error) {
	var resp core.GenericResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.WebhookXenditService.HandlePaymentInvoice)
}
