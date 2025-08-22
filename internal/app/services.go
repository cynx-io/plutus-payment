package app

import (
	"github.com/cynx-io/plutus-payment/internal/service/exampleservice"
	"github.com/cynx-io/plutus-payment/internal/service/paymentservice"
	"github.com/cynx-io/plutus-payment/internal/service/webhookxenditservice"
)

type Services struct {
	ExampleService       *exampleservice.Service
	PaymentService       *paymentservice.Service
	WebhookXenditService *webhookxenditservice.Service
}

func NewServices(repos *Repos) *Services {
	return &Services{
		ExampleService: exampleservice.New(repos.ExampleRepo),
		PaymentService: &paymentservice.Service{
			TblCustomer:       repos.TblCustomer,
			TblPaymentInvoice: repos.TblPaymentInvoice,

			HermesUserClient: repos.HermesUserClient,

			XenditClient: repos.XenditClient,
		},
		WebhookXenditService: &webhookxenditservice.Service{
			TblCustomer:       repos.TblCustomer,
			TblPaymentInvoice: repos.TblPaymentInvoice,

			HermesUserClient:     repos.HermesUserClient,
			AnankePreorderClient: repos.AnankePreorderClient,

			XenditClient: repos.XenditClient,
		},
	}
}
