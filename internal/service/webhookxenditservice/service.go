package webhookxenditservice

import (
	pbananke "github.com/cynx-io/plutus-payment/api/proto/gen/ananke"
	pbhermes "github.com/cynx-io/plutus-payment/api/proto/gen/hermes"
	"github.com/cynx-io/plutus-payment/internal/repository/database"
	"github.com/xendit/xendit-go/v7"
)

type Service struct {
	TblCustomer       *database.TblCustomer
	TblPaymentInvoice *database.TblPaymentInvoice

	HermesUserClient     pbhermes.HermesUserServiceClient
	AnankePreorderClient pbananke.PreorderServiceClient

	XenditClient *xendit.APIClient
}
