package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/db"
	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
	repositoryInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type vaultRepository struct {
	connectionString *db.CollectionConnections
}

func NewVaultRepository(db *db.CollectionConnections) repositoryInterface.IVaultRepository {
	return &vaultRepository{connectionString: db}
}

func (d *vaultRepository) CreateCollection(collectionData requestmodel.Colleciton) (string, error) {
	result, err := d.connectionString.Category.InsertOne(context.Background(), bson.M{"CategoryName": collectionData.CategoryName, "userID": collectionData.UserID})
	fmt.Println("result: ", result)
	if err != nil {
		return "", err
	}

	value := result.InsertedID.(primitive.ObjectID).Hex()

	return value, nil
}

func (d *vaultRepository) InsertData(data requestmodel.Data) (string, error) {
	validDateInt, _ := strconv.Atoi(data.Reminder)
	expire := time.Now().AddDate(0, 0, validDateInt)

	objID, err := d.connectionString.Data.InsertOne(context.Background(), bson.M{"userID": data.UserID, "categoryID": data.CategoryID, "data": data.Data, "expire": expire})
	if err != nil {
		return "", err
	}

	return objID.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *vaultRepository) GetCategories(userID string) ([]*resposemodel.Collections, error) {
	var category *resposemodel.Collections
	var categorySet []*resposemodel.Collections
	courser, _ := d.connectionString.Category.Find(context.Background(), bson.M{"userID": userID})

	defer courser.Close(context.Background())

	for courser.Next(context.Background()) {
		if err := courser.Decode(&category); err != nil {
			return nil, errors.New("error on unmarshel of category data in valt service repository")
		}
		categorySet = append(categorySet, category)
	}
	return categorySet, nil
}

func (d *vaultRepository) GetDatas(details requestmodel.GetDataRequest) (*[]resposemodel.GetDataResponse, error) {
	var dataList []resposemodel.GetDataResponse
	var singleData resposemodel.GetDataResponse
	currentData := time.Now().Truncate(24 * time.Hour)

	courser, _ := d.connectionString.Data.Find(context.Background(), bson.M{"categoryID": details.CategoryID, "userID": details.UserID, "expire": bson.M{"$gte": currentData}})
	defer courser.Close(context.Background())

	for courser.Next(context.Background()) {
		if err := courser.Decode(&singleData); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		dataList = append(dataList, singleData)
	}

	return &dataList, nil
}

func (d *vaultRepository) DeleteDataCronJob() {
	currentDate := time.Now().Truncate(24 * time.Hour)
	
	d.connectionString.Data.DeleteMany(context.Background(), bson.M{"expire": bson.M{"$lt": currentDate}})

	log.Println("--cron workded delete happened")
}	
