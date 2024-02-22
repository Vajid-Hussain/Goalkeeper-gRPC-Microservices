package db

import (
	"log"

	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit(c *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("database is not connected", err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
