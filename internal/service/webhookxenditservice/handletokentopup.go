package webhookxenditservice

import (
	"context"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) handleTokenTopUp(ctx context.Context, req *pb.HandlePaymentInvoiceRequest, resp *core.GenericResponse, invoice *entity.TblPaymentInvoice) error {

	tokenInvoice, err := s.TblTokenInvoice.GetTokenInvoiceByPaymentInvoiceId(ctx, invoice.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseTokenPriceList(resp)
	}

	if !tokenInvoice.IsTopUp {
		response.ErrorNotAllowed(resp)
		return errors.New("not a top up invoice")
	}

	if tokenInvoice.Status != int32(pb.PaymentInvoiceStatus_PENDING) {
		response.ErrorNotAllowed(resp)
		return errors.New("invoice already processed")
	}

	err = s.TblBalance.IncrementBalance(ctx, tokenInvoice.UserId, -tokenInvoice.TokenUsed)
	if err != nil {
		response.ErrorDatabaseBalance(resp)
		return err
	}

	err = s.TblTokenInvoice.UpdateTokenInvoiceStatus(ctx, tokenInvoice.Id, int32(pb.PaymentInvoiceStatus_COMPLETED))
	if err != nil {
		response.ErrorDatabaseTokenInvoice(resp)
		return err
	}

	response.Success(resp)
	return nil
}
