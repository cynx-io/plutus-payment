package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

type TblTokenPriceList struct {
	entity.EssentialEntity
	Name        string  `gorm:"size:255;not null" json:"name"`
	Description string  `gorm:"size:255;not null" json:"description"`
	Currency    string  `gorm:"size:10;not null" json:"currency"`
	Token       float32 `gorm:"not null" json:"token"`
	Price       float32 `gorm:"not null" json:"price"`
}

func (b TblTokenPriceList) Response() *proto.TokenPriceList {
	return &proto.TokenPriceList{
		Id:          b.Id,
		Name:        b.Name,
		Description: b.Description,
		Token:       b.Token,
		Price:       b.Price,
		Currency:    b.Currency,
	}
}
