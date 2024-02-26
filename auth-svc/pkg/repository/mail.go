package repository

import (
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
	repositoryIinterfaces "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type mailRepository struct {
	DB *gorm.DB
}

func NewMailRepository(db *gorm.DB) repositoryIinterfaces.IMailRepository {
	return &mailRepository{DB: db}
}

func (d *mailRepository) FetchUserDetails() (*[]resposemodel.MailUserDetils, error) {
	var users []resposemodel.MailUserDetils
	query := "SELECT * FROM users"
	result := d.DB.Raw(query).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
