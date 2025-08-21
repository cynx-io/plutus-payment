package micro

import (
	pbhermes "github.com/cynx-io/plutus-payment/api/proto/gen/hermes"
	"github.com/cynx-io/plutus-payment/internal/dependencies/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHermesUserClient() pbhermes.HermesUserServiceClient {

	conn, err := grpc.NewClient(config.Config.Hermes.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Hermes gRPC server: " + err.Error())
	}

	userClient := pbhermes.NewHermesUserServiceClient(conn)
	return userClient
}
