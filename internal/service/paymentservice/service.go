package paymentservice

import (
	pbhermes "github.com/cynx-io/plutus-payment/api/proto/gen/hermes"
	"github.com/cynx-io/plutus-payment/internal/repository/database"
	"github.com/xendit/xendit-go/v7"
)

type Service struct {
	TblCustomer       *database.TblCustomer
	TblPaymentInvoice *database.TblPaymentInvoice

	HermesUserClient pbhermes.HermesUserServiceClient

	XenditClient *xendit.APIClient
}
