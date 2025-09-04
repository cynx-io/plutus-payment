package paymentservice

import (
	"context"
	"errors"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) TopUpBalance(ctx context.Context, req *pb.TopUpBalanceRequest, resp *pb.PaymentInvoiceResponse) error {

	if req.Base.UserId == nil || *req.Base.UserId == 0 {
		response.ErrorUnauthorized(resp)
		return errors.New("user id is required")
	}

	tokenPriceList, err := s.TblTokenPriceList.GetTokenPriceListById(ctx, req.TokenPriceListId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseTokenPriceList(resp)
	}

	invoiceReq := &pb.CreatePaymentInvoiceRequest{
		Base:              req.Base,
		UserId:            *req.Base.UserId,
		Amount:            tokenPriceList.Price,
		Currency:          tokenPriceList.Currency,
		Description:       tokenPriceList.Description,
		SuccessReturnUrl:  req.SuccessReturnUrl,
		FailureReturnUrl:  req.FailureReturnUrl,
		DurationInSeconds: 24 * 60 * 60, // 24 hours
		PaymentFeature:    pb.PaymentFeature_TOKEN_TOP_UP,
	}

	err = s.CreatePaymentInvoice(ctx, invoiceReq, resp)
	if err != nil {
		return err
	}

	response.Success(resp)
	return nil
}
