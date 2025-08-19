package app

import (
	"github.com/cynx-io/micro-name/internal/service/exampleservice"
)

type Services struct {
	ExampleService *exampleservice.Service
}

func NewServices(repos *Repos, dependencies *Dependencies) *Services {
	return &Services{
		ExampleService: exampleservice.New(repos.ExampleRepo),
	}
}
