package paymentservice

import (
	"context"
	"errors"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) GetProductPriceListById(ctx context.Context, req *pb.GetProductPriceListByIdRequest, resp *pb.ProductPriceListResponse) error {

	productPriceList, err := s.TblProductPriceList.GetProductPriceListById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseProductPriceList(resp)
		return err
	}

	response.Success(resp)
	resp.ProductPriceList = productPriceList.Response()
	return nil
}
