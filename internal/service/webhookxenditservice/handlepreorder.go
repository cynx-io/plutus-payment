package webhookxenditservice

import (
	"context"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pbananke "github.com/cynx-io/plutus-payment/api/proto/gen/ananke"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/helper"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"github.com/cynx-io/plutus-payment/internal/model/response"
)

func (s *Service) handlePreorder(ctx context.Context, req *pb.HandlePaymentInvoiceRequest, resp *core.GenericResponse, invoice *entity.TblPaymentInvoice) error {

	preorderResp, err := s.AnankePreorderClient.ChangePreorderStatusByInvoiceId(ctx, &pbananke.ChangePreorderStatusByInvoiceIdRequest{
		Base:              req.Base,
		InvoiceId:         invoice.Id,
		TransactionStatus: pbananke.TransactionStatus(helper.XenditInvoiceStatusToProto(req.Status)),
	})
	if err != nil {
		response.ErrorAnanke(resp)
		return err
	}

	if preorderResp.Base.Code != response.CodeSuccess.String() {
		response.ErrorAnanke(resp)
		return errors.New("failed to change preorder status: " + preorderResp.Base.Code + ": " + preorderResp.Base.Desc)
	}

	response.Success(resp)
	return nil
}
