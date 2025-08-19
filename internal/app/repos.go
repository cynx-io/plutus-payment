package app

import (
	"github.com/cynx-io/micro-name/internal/repository/database"
)

type Repos struct {
	ExampleRepo *database.ExampleRepo
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		ExampleRepo: database.NewExampleRepo(dependencies.DatabaseClient.Db),
	}
}
