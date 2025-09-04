package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

type TblProductPriceList struct {
	entity.EssentialEntity
	Id          string  `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Description string  `gorm:"size:255;not null" json:"description"`
	Token       float32 `gorm:"not null" json:"token"`
	Denominator int32   `gorm:"not null" json:"denominator"`
}

func (b TblProductPriceList) Response() *proto.ProductPriceList {
	return &proto.ProductPriceList{
		Id:          b.Id,
		Name:        b.Name,
		Description: b.Description,
		Token:       b.Token,
		Denominator: proto.Denominator(b.Denominator),
	}
}
