package main

import (
	"log"
	"net"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/di"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("error on viper:", err)
	}

	service := di.InitializeService(config)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalln("cannot make a tcp socket:", err)
	}

	grpcService := grpc.NewServer()
	pb.RegisterVaultServiceServer(grpcService, service)

	if err := grpcService.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
