package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

type TblTokenInvoice struct {
	entity.EssentialEntity
	UserId               int32               `gorm:"not null" json:"user_id"`
	ProductPriceListId   *string             `gorm:"size:36;index;default:null" json:"product_price_list_id"`
	IsTopUp              bool                `gorm:"not null;default:false" json:"is_top_up"`
	PreviousTokenBalance float32             `gorm:"not null" json:"previous_token_balance"`
	NewTokenBalance      float32             `gorm:"not null" json:"new_token_balance"`
	TokenUsed            float32             `gorm:"not null" json:"token_used"`
	ProductPriceList     TblProductPriceList `gorm:"foreignKey:ProductPriceListId;references:Id" json:"product_price_list"`
}

func (b TblTokenInvoice) Response() *proto.TokenInvoice {
	return &proto.TokenInvoice{
		Id:                   b.Id,
		UserId:               b.UserId,
		ProductPriceListId:   b.ProductPriceListId,
		IsTopUp:              b.IsTopUp,
		PreviousTokenBalance: b.PreviousTokenBalance,
		NewTokenBalance:      b.NewTokenBalance,
		TokenUsed:            b.TokenUsed,
	}
}
