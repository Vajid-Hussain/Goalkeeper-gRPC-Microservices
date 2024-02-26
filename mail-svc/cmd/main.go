package main

import (
	"log"
	"net"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/clind"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/pb"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("env is not loadded, ", err)
	}

	c := clind.InitReminderServiceClient(config.VaultSvcPort)

	lis, err := net.Listen("tcp", config.MailSvcPort)
	if err != nil {
		log.Fatal("tcp connection no build in system leavel")
	}
	
	
	grpcService := grpc.NewServer()
	pb.RegisterMailServiceServer(grpcService, server.MailServerInstance(c, config))

	if err:=grpcService.Serve(lis); err!=nil{
		log.Fatal("connection is not established")
	}
}
