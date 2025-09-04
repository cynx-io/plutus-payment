package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
)

type TblBalance struct {
	entity.EssentialEntity
	UserId       int32   `gorm:"not null;uniqueIndex" json:"user_id"`
	TokenBalance float32 `gorm:"not null" json:"token_balance"`
}

func (b TblBalance) Response() *proto.Balance {
	return &proto.Balance{
		UserId:       b.UserId,
		TokenBalance: b.TokenBalance,
	}
}
