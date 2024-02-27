package db

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionConnections struct {
	Category *mongo.Collection
	Data     *mongo.Collection
}

func DBConnection(cofig *config.Config) *CollectionConnections {
	var connections CollectionConnections
	var ctx = context.TODO()

	clindOption := options.Client().ApplyURI(cofig.MongoUrl)
	client, err := mongo.Connect(ctx, clindOption)
	if err != nil {
		fmt.Println("connection refuse when make a connection to mongod")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("ping")
	} else {
		fmt.Println("connected to mongodb")
	}

	categoryCollection := client.Database("goalkeeper").Collection("vault")
	dataCollection := client.Database("goalkeeper").Collection("datas")

	connections.Category = categoryCollection
	connections.Data = dataCollection
	return &connections
}
