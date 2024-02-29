package auth

import (
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/pb"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// kuberesolver.RegisterInCluster()
	// client, _ := kuberesolver.NewInClusterK8sClient()
	// resolver.Register(kuberesolver.NewBuilder(client))

	// fmt.Println("+++++++++++", c)

	cc, err := grpc.Dial(c.Auth_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// fmt.Println(c.Auth_port, "\n", cc)
	if err != nil {
		fmt.Println("cound not connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
