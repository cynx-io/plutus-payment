package xendit

import (
	"github.com/cynx-io/plutus-payment/internal/dependencies/config"
	exxendit "github.com/xendit/xendit-go/v7"
)

func New() *exxendit.APIClient {
	client := exxendit.NewClient(config.Config.Xendit.ApiKey)
	if client == nil {
		panic("Failed to create Xendit client: client is nil")
	}
	return client
}
