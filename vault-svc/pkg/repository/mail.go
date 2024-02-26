package repository

import (
	"context"
	"time"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/db"
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
	repositoryInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
)

type mailRepository struct {
	connectionString *db.CollectionConnections
}

func NewMailRepository(db *db.CollectionConnections) repositoryInterface.IMailRepository {
	return &mailRepository{connectionString: db}
}

func (d *mailRepository) TodayTask(userID string) ([]*resposemodel.TodayTask, error) {
	var dataSet []*resposemodel.TodayTask
	now := time.Now().UTC().Truncate(24 * time.Hour)
	nextDay := now.Add(24 * time.Hour)

	courser, err := d.connectionString.Data.Find(context.Background(), bson.M{"userID": userID, "expire": bson.M{"$gte": now, "$lt": nextDay}})
	if err != nil {
		return nil, err
	}

	for courser.Next(context.Background()) {
		var data resposemodel.TodayTask
		if err := courser.Decode(&data); err != nil {
			return nil, err
		}
		dataSet = append(dataSet, &data)
	}
	return dataSet, nil
}
