package grpc

import (
	"context"
	"github.com/cynx-io/cynx-core/src/grpc"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

func (s *Server) CreatePaymentInvoice(ctx context.Context, req *proto.CreatePaymentInvoiceRequest) (*proto.PaymentInvoiceResponse, error) {
	var resp proto.PaymentInvoiceResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.CreatePaymentInvoice)
}
