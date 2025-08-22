package micro

import (
	pbananke "github.com/cynx-io/plutus-payment/api/proto/gen/ananke"
	"github.com/cynx-io/plutus-payment/internal/dependencies/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAnankeUserClient() pbananke.PreorderServiceClient {

	conn, err := grpc.NewClient(config.Config.Ananke.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Ananke gRPC server: " + err.Error())
	}

	userClient := pbananke.NewPreorderServiceClient(conn)
	return userClient
}
