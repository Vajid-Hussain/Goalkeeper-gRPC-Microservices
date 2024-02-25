package vault

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/vault/handler"
)

func RegisterVault(ctx *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	authMiddlewire := auth.NewAuthServiceClient(authSvc)

	svc := ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := ctx.Group("/vault")
	{
		routes.Use(authMiddlewire.VerifyJwt)
		routes.POST("/createcollection", svc.CreateCollection)
		routes.POST("/inserdata", svc.InserData)
		routes.POST("/getcategory", svc.GetCategory)
		routes.POST("/getdatas", svc.GetDatas)
	}
}

func (s *ServiceClient) CreateCollection(ctx *gin.Context) {
	handler.CreateCollection(ctx, s.Client)
}

func (s *ServiceClient) InserData(ctx *gin.Context) {
	handler.InserData(ctx, s.Client)
}

func (s *ServiceClient) GetCategory(ctx *gin.Context) {
	handler.GetCategory(ctx, s.Client)
}

func (s *ServiceClient) GetDatas(ctx *gin.Context) {
	handler.GetDatas(ctx, s.Client)
}
