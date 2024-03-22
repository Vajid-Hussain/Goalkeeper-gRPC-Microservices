package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/di"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error on viper: ", err)
	}

	service := di.InitializeService(&configs)

	lis, err := net.Listen("tcp", configs.Port)
	if err != nil {
		log.Fatalln("cannot make a tcp socket: ", err)
	}

	// tlsCredential, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal("cannot load TLS credentials: ", err)
	// }

	// // grpcTlsService:= grpc.NewServer(
	// // 	grpc.Creds(tlsCredential),
	// // 	grpc.UnaryInterceptor(interceptor.Unary()),

	// // )

	creds, err := credentials.NewServerTLSFromFile("./service.pem", "./service.key")
	// creds, err := credentials.NewServerTLSFromFile("./../service.pem", "./../service.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}
	fmt.Println("--")

	s := grpc.NewServer(grpc.Creds(creds))

	// grpcService := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, service)

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
