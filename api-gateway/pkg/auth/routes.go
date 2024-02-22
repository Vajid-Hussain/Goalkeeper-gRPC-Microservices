package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/config"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/handler"
)

func RegisterRouter(engin *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	router := engin.Group("/auth")
	router.POST("/register", svc.Register)
	router.POST("/login", svc.Login)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	handler.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	handler.Login(ctx, svc.Client)
}
