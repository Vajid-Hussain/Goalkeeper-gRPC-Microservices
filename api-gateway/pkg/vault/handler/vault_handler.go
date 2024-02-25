package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/vault/pb"
	requestmodel "github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/vault/requestModel"
)

func CreateCollection(ctx *gin.Context, c pb.VaultServiceClient) {
	var collectionDetail requestmodel.CollecionDetails

	err := ctx.BindJSON(&collectionDetail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	collectionDetail.UserID = ctx.MustGet("userID").(string)

	// fmt.Println("@@@@", collectionDetail)
	response, err := c.CreateCollection(context.Background(), &pb.CreateCollectionRequest{
		CollectionName: collectionDetail.CollectionName,
		UserID:         collectionDetail.UserID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func InserData(ctx *gin.Context, c pb.VaultServiceClient) {
	fmt.Print("workded")
	data := requestmodel.Data{}

	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := c.InserData(context.Background(), &pb.DataRequest{
		UserID:     ctx.MustGet("userID").(string),
		CategoryID: data.CategoryID,
		Data:       data.Data,
		Reminder:   data.Reminder,
	})
	fmt.Println("000", data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetCategory(ctx *gin.Context, c pb.VaultServiceClient) {
	result, err := c.GetCategories(context.Background(), &pb.CategoryRequest{
		UserID: ctx.MustGet("userID").(string),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func GetDatas(ctx *gin.Context, c pb.VaultServiceClient) {
	var dataRequest requestmodel.GetDataRequest

	ctx.BindJSON(&dataRequest)
	result, err := c.GetDatas(context.Background(), &pb.GetDataRequest{
		UserID:     ctx.MustGet("userID").(string),
		CategoryID: dataRequest.CategoryID,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var finalResult []requestmodel.GetDataResponse
	for _, val := range result.Datas {
		finalResult = append(finalResult, requestmodel.GetDataResponse{Data: val.Data, Remainder: time.Unix((*val).Reminder.Seconds, int64((*val).Reminder.Nanos))})
	}

	ctx.JSON(http.StatusOK, finalResult)
}
