package exampleservice

import (
	"github.com/cynx-io/micro-name/internal/repository/database"
)

type Service struct {
	ExampleRepo *database.ExampleRepo
}

func New(exampleRepo *database.ExampleRepo) *Service {
	return &Service{
		ExampleRepo: exampleRepo,
	}
}
