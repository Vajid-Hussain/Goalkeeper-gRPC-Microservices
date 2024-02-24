package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/vault"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config err", err)
	}

	engin := gin.Default()

	authSVC:=auth.RegisterRouter(engin, config)
	vault.RegisterVault(engin, config, authSVC)

	engin.Run(config.Port)
}
