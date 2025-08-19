package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
)

type TblExample struct {
	entity.EssentialEntity
}

func (u TblExample) Response() *TblExample {
	return &u
}
