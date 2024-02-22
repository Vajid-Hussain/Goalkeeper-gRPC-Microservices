package auth

import (
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/config"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	fmt.Println("+++++++++++", c)
	cc, err := grpc.Dial(c.Auth_port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("cound not connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
