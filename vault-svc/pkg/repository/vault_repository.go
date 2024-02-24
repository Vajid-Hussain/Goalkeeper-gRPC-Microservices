package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/db"
	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
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

	datas, _ := d.connectionString.Category.Find(context.Background(), bson.M{"userID": data.UserID})
	fmt.Println("------------------", datas)

	return objID.InsertedID.(primitive.ObjectID).Hex(), nil
}

// func (d *vaultRepository) GetCollection(userID string) {
// 	data, _ := d.connectionString.Category.Find(context.Background(), bson.M{"userID": userID})
// }
