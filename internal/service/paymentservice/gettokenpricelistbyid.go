package paymentservice

import (
	"context"
	"errors"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) GetTokenPriceListById(ctx context.Context, req *pb.GetTokenPriceListByIdRequest, resp *pb.TokenPriceListResponse) error {

	tokenPriceList, err := s.TblTokenPriceList.GetTokenPriceListById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseTokenPriceList(resp)
		return err
	}

	response.Success(resp)
	resp.TokenPriceList = tokenPriceList.Response()
	return nil
}
