package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
)

type TblCustomer struct {
	Email       *string `gorm:"size:255" json:"email"`
	FullName    *string `gorm:"size:255" json:"full_name"`
	PhoneNumber *string `gorm:"size:50" json:"phone_number"`
	entity.EssentialEntity
	ExternalId string `gorm:"size:255;not null" json:"external_id"`
	Provider   string `gorm:"size:50;not null" json:"provider"`
	UserId     int32  `gorm:"not null" json:"user_id"`
}
