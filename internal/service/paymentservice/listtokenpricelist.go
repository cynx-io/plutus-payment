package paymentservice

import (
	"context"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) ListTokenPriceList(ctx context.Context, req *core.GenericRequest, resp *pb.TokenPriceListListResponse) error {

	tokenPriceList, err := s.TblTokenPriceList.ListTokenPriceList(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseTokenPriceList(resp)
		return err
	}

	listResponse := make([]*pb.TokenPriceList, len(tokenPriceList))
	for i, v := range tokenPriceList {
		listResponse[i] = v.Response()
	}

	response.Success(resp)
	resp.TokenPriceLists = listResponse
	return nil
}
