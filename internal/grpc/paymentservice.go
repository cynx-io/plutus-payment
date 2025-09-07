package grpc

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/grpc"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

func (s *Server) CreatePaymentInvoice(ctx context.Context, req *proto.CreatePaymentInvoiceRequest) (*proto.PaymentInvoiceResponse, error) {
	var resp proto.PaymentInvoiceResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.CreatePaymentInvoice)
}

func (s *Server) TopUpBalance(ctx context.Context, req *proto.TopUpBalanceRequest) (*proto.PaymentInvoiceResponse, error) {
	var resp proto.PaymentInvoiceResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.TopUpBalance)
}

func (s *Server) GetBalance(ctx context.Context, req *core.GenericRequest) (*proto.BalanceResponse, error) {
	var resp proto.BalanceResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.GetBalance)
}

func (s *Server) GetProductPriceListById(ctx context.Context, req *proto.GetProductPriceListByIdRequest) (*proto.ProductPriceListResponse, error) {
	var resp proto.ProductPriceListResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.GetProductPriceListById)
}

func (s *Server) GetTokenPriceListById(ctx context.Context, req *proto.GetTokenPriceListByIdRequest) (*proto.TokenPriceListResponse, error) {
	var resp proto.TokenPriceListResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.GetTokenPriceListById)
}

func (s *Server) ListTokenPriceList(ctx context.Context, req *core.GenericRequest) (*proto.TokenPriceListListResponse, error) {
	var resp proto.TokenPriceListListResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.ListTokenPriceList)
}

func (s *Server) PurchaseProduct(ctx context.Context, req *proto.PurchaseProductRequest) (*proto.TokenInvoiceResponse, error) {
	var resp proto.TokenInvoiceResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.PurchaseProduct)
}
