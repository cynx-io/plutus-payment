package app

import (
	"context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/plutus-payment/internal/dependencies"
)

type Dependencies struct {
	DatabaseClient *dependencies.DatabaseClient
}

func NewDependencies(ctx context.Context) *Dependencies {

	logger.Info(ctx, "Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient()
	if err != nil {
		logger.Fatal(ctx, "Failed to connect to database: ", err)
		panic("fail to connect to database " + err.Error())
	}

	logger.Info(ctx, "Dependencies initialized")
	return &Dependencies{
		DatabaseClient: databaseClient,
	}
}
