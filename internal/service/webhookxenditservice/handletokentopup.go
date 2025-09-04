package webhookxenditservice

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"github.com/cynx-io/plutus-payment/internal/model/response"
)

func (s *Service) handleTokenTopUp(ctx context.Context, req *pb.HandlePaymentInvoiceRequest, resp *core.GenericResponse, invoice *entity.TblPaymentInvoice) error {

	err := s.TblBalance.IncrementBalance(ctx, invoice.Customer.UserId, invoice.Amount)
	if err != nil {
		response.ErrorDatabaseBalance(resp)
	}

	response.Success(resp)
	return nil
}
