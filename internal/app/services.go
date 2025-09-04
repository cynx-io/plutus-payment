package app

import (
	"github.com/cynx-io/plutus-payment/internal/service/paymentservice"
	"github.com/cynx-io/plutus-payment/internal/service/webhookxenditservice"
)

type Services struct {
	PaymentService       *paymentservice.Service
	WebhookXenditService *webhookxenditservice.Service
}

func NewServices(repos *Repos) *Services {
	return &Services{
		PaymentService: &paymentservice.Service{
			TblCustomer:         repos.TblCustomer,
			TblPaymentInvoice:   repos.TblPaymentInvoice,
			TblProductPriceList: repos.TblProductPriceList,
			TblTokenInvoice:     repos.TblTokenInvoice,
			TblBalance:          repos.TblBalance,
			TblTokenPriceList:   repos.TblTokenPriceList,

			HermesUserClient: repos.HermesUserClient,

			XenditClient: repos.XenditClient,
		},
		WebhookXenditService: &webhookxenditservice.Service{
			TblCustomer:       repos.TblCustomer,
			TblPaymentInvoice: repos.TblPaymentInvoice,
			TblTokenInvoice:   repos.TblTokenInvoice,
			TblBalance:        repos.TblBalance,

			HermesUserClient:     repos.HermesUserClient,
			AnankePreorderClient: repos.AnankePreorderClient,

			XenditClient: repos.XenditClient,
		},
	}
}
