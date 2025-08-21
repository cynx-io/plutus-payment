package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TblPaymentInvoice struct {
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	entity.EssentialEntity
	Id             string      `gorm:"primaryKey" json:"id"`
	ExternalId     string      `gorm:"size:255;not null" json:"external_id"`
	RequestId      string      `gorm:"size:255;not null" json:"request_id"`
	Status         string      `gorm:"size:50;not null" json:"status"`
	PaymentLinkUrl string      `gorm:"size:255;not null" json:"payment_link_url"`
	Currency       string      `gorm:"size:10;not null" json:"currency"`
	Description    string      `gorm:"size:255;not null" json:"description"`
	Provider       string      `gorm:"size:50;not null" json:"provider"`
	Customer       TblCustomer `gorm:"foreignkey:CustomerId" json:"customer"`
	Amount         int64       `gorm:"not null" json:"amount"`
	CustomerId     int32       `gorm:"not null" json:"customer_id"`
}

func (p TblPaymentInvoice) Response() *proto.PaymentInvoice {

	protoStatus := helper.ToProtoInvoiceStatus(p.Status)
	return &proto.PaymentInvoice{
		Id:             p.Id,
		ExternalId:     p.ExternalId,
		RequestId:      p.RequestId,
		Status:         protoStatus,
		PaymentLinkUrl: p.PaymentLinkUrl,
		Currency:       p.Currency,
		Description:    p.Description,
		Provider:       p.Provider,
		CustomerId:     p.CustomerId,
		Amount:         p.Amount,
		ExpiresAt:      timestamppb.New(p.ExpiresAt),
	}

}
