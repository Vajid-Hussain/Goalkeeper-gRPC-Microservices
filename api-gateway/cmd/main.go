package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config err", err)
	}

	engin := gin.Default()

	auth.RegisterRouter(engin, config)

	engin.Run(config.Port)
}
