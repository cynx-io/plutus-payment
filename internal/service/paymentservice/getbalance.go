package paymentservice

import (
	"context"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/response"
)

func (s *Service) GetBalance(ctx context.Context, req *core.GenericRequest, resp *pb.BalanceResponse) error {

	if req.Base.UserId == nil || *req.Base.UserId == 0 {
		response.ErrorUnauthorized(resp)
		return errors.New("user id is required")
	}

	balance, err := s.TblBalance.GetBalanceByUserId(ctx, *req.Base.UserId)
	if err != nil {
		response.ErrorDatabaseBalance(resp)
		return err
	}

	response.Success(resp)
	resp.Balance = balance.Response()
	return nil
}
