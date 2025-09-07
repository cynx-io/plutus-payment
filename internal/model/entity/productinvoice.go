package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

type TblTokenInvoice struct {
	ProductPriceListId *string `gorm:"size:36;index;default:null" json:"product_price_list_id"`
	TokenPriceListId   *int32  `gorm:"default:null" json:"token_price_list_id"`
	entity.EssentialEntity
	ProductPriceList     TblProductPriceList `gorm:"foreignKey:ProductPriceListId;references:Id" json:"product_price_list"`
	TokenPriceList       TblTokenPriceList   `gorm:"foreignKey:TokenPriceListId;references:Id" json:"token_price_list"`
	UserId               int32               `gorm:"not null" json:"user_id"`
	PreviousTokenBalance float32             `gorm:"not null" json:"previous_token_balance"`
	NewTokenBalance      float32             `gorm:"not null" json:"new_token_balance"`
	TokenUsed            float32             `gorm:"not null" json:"token_used"`
	IsTopUp              bool                `gorm:"not null;default:false" json:"is_top_up"`
	PaymentInvoiceId     *string             `gorm:"size:255;index;default:null" json:"payment_invoice_id"`
	PaymentInvoice       TblPaymentInvoice   `gorm:"foreignKey:PaymentInvoiceId;references:Id" json:"payment_invoice"`
	Status               int32               `gorm:"not null" json:"status"`
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
