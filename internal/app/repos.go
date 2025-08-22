package app

import (
	pbananke "github.com/cynx-io/plutus-payment/api/proto/gen/ananke"
	pbhermes "github.com/cynx-io/plutus-payment/api/proto/gen/hermes"
	"github.com/cynx-io/plutus-payment/internal/repository/database"
	xendit2 "github.com/cynx-io/plutus-payment/internal/repository/externalapi/xendit"
	"github.com/cynx-io/plutus-payment/internal/repository/micro"
	"github.com/xendit/xendit-go/v7"
)

type Repos struct {
	ExampleRepo *database.ExampleRepo

	TblCustomer       *database.TblCustomer
	TblPaymentInvoice *database.TblPaymentInvoice

	HermesUserClient     pbhermes.HermesUserServiceClient
	AnankePreorderClient pbananke.PreorderServiceClient
	XenditClient         *xendit.APIClient
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		ExampleRepo: database.NewExampleRepo(dependencies.DatabaseClient.Db),

		TblCustomer:       database.NewCustomerRepo(dependencies.DatabaseClient.Db),
		TblPaymentInvoice: database.NewPaymentInvoiceRepo(dependencies.DatabaseClient.Db),

		HermesUserClient:     micro.NewHermesUserClient(),
		AnankePreorderClient: micro.NewAnankeUserClient(),

		XenditClient: xendit2.New(),
	}
}
