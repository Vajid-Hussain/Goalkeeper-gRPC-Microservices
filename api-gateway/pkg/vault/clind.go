package vault

import (
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/vault/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.VaultServiceClient
}

func InitServiceClient(c *config.Config) pb.VaultServiceClient {
	cc, err := grpc.Dial(c.Vault_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect: ", err)
	}

	return pb.NewVaultServiceClient(cc)
}
