package paymentservice

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) PurchaseProduct(ctx context.Context, req *pb.PurchaseProductRequest, resp *pb.TokenInvoiceResponse) error {

	if req.Base.UserId == nil || *req.Base.UserId == 0 {
		response.ErrorUnauthorized(resp)
		return errors.New("user id is required")
	}

	productPriceList, err := s.TblProductPriceList.GetProductPriceListById(ctx, req.ProductPriceListId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabaseProductPriceList(resp)
		return err
	}

	balance, err := s.TblBalance.GetBalanceByUserId(ctx, *req.Base.UserId)
	if err != nil {
		response.ErrorDatabaseBalance(resp)
		return err
	}

	if balance.TokenBalance < productPriceList.Token {
		response.ErrorInsufficientBalance(resp)
		return fmt.Errorf("insufficient balance : balance %f, required %f", balance.TokenBalance, productPriceList.Token)
	}

	err = s.TblBalance.DecrementBalance(ctx, *req.Base.UserId, productPriceList.Token)
	if err != nil {
		response.ErrorDatabaseBalance(resp)
		return errors.New("balance decrement failed: " + err.Error())
	}

	invoice := &entity.TblTokenInvoice{
		UserId:               *req.Base.UserId,
		ProductPriceListId:   &req.ProductPriceListId,
		IsTopUp:              false,
		NewTokenBalance:      balance.TokenBalance - productPriceList.Token,
		PreviousTokenBalance: balance.TokenBalance,
		TokenUsed:            productPriceList.Token,
	}

	err = s.TblTokenInvoice.CreateTokenInvoice(ctx, invoice)
	if err != nil {
		response.ErrorDatabaseTokenPriceList(resp)
		return errors.New("create token invoice failed: " + err.Error())
	}

	response.Success(resp)
	resp.TokenInvoice = invoice.Response()
	return nil
}
