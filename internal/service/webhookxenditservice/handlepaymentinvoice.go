package webhookxenditservice

import (
	"context"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/dependencies/config"
	"github.com/cynx-io/plutus-payment/internal/helper"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) HandlePaymentInvoice(ctx context.Context, req *pb.HandlePaymentInvoiceRequest, resp *core.GenericResponse) error {

	if req.WebhookKey != config.Config.Xendit.WebhookKey {
		response.ErrorUnauthorized(resp)
		return errors.New("invalid webhook key")
	}

	invoice, err := s.TblPaymentInvoice.GetPaymentInvoiceById(ctx, req.ExternalId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return errors.New("payment invoice not found")
		}
		response.ErrorDatabaseInvoice(resp)
		return err
	}

	status := helper.XenditInvoiceStatusToProto(req.Status)

	if status == pb.PaymentInvoiceStatus_PENDING {
		response.Success(resp)
		resp.Base.Desc = "Payment invoice is still pending, current status: " + pb.PaymentInvoiceStatus_name[invoice.Status]
		return nil
	}

	if invoice.Status != int32(pb.PaymentInvoiceStatus_PENDING) {
		response.Success(resp)
		resp.Base.Desc = "Payment invoice is not pending: " + pb.PaymentInvoiceStatus_name[invoice.Status]
		return nil
	}

	err = s.TblPaymentInvoice.UpdatePaymentInvoiceStatus(ctx, invoice.Id, int32(status))
	if err != nil {
		response.ErrorDatabaseInvoice(resp)
		return err
	}

	if invoice.PaymentFeature == int32(pb.PaymentFeature_PREORDER) {
		return s.handlePreorder(ctx, req, resp, invoice)
	}

	if invoice.PaymentFeature == int32(pb.PaymentFeature_TOKEN_TOP_UP) {
		return s.handleTokenTopUp(ctx, req, resp, invoice)
	}

	response.Success(resp)
	return nil
}
