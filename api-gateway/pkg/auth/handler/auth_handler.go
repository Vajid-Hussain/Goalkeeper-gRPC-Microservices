package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	requestmodel "github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/models/requestModel"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/pb"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := requestmodel.RegisterReqestBody{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	// fmt.Println("----", body)

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	fmt.Println("###########", res, err)

	if err != nil {
		// fmt.Println("error at rquest in handler")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	loginData := requestmodel.RegisterReqestBody{}

	err := ctx.BindJSON(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	// fmt.Println("----", loginData)

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    loginData.Email,
		Password: loginData.Password,
	})
	// fmt.Println("###########", res)
	if err != nil {
		fmt.Println("===error at login :", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
