package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/pb"
)

type AuthMiddleWire struct {
	svc *ServiceClient
}

func NewAuthServiceClient(service *ServiceClient) *AuthMiddleWire {
	return &AuthMiddleWire{svc: service}
}

func (c *AuthMiddleWire) VerifyJwt(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, errors.New("no authorization token exist"))
	}

	res, err := c.svc.Client.JwtValidate(context.Background(), &pb.JwtRequest{
		Jwt: token,
	})
	// fmt.Println("-------", res, err)
	if err != nil {
		fmt.Println("***", err)
		ctx.JSON(http.StatusUnauthorized, err.Error())
		ctx.Abort()
		return
	}
	ctx.Set("userID", res.UserID)
	ctx.Next()
}
