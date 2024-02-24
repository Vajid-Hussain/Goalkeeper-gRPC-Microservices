package repository

import (
	"errors"
	"fmt"

	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
	repositoryIinterfaces "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repositoryIinterfaces.IAuthRepository {
	return &authRepository{DB: db}
}

func (d *authRepository) CreateUser(userData *reqestmodel.User) (*resposemodel.UserData, error) {
	var createdUser resposemodel.UserData
	query := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING *"
	result := d.DB.Raw(query, userData.Email, userData.Password).Scan(&createdUser)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, fmt.Errorf("insert user data face some issue: %v", result.Error)
	}

	return &createdUser, nil
}

func (d *authRepository) UserExist(email string) (count int) {

	query := "SELECT count(*) FROM users WHERE email= $1"
	d.DB.Raw(query, email).Scan(&count)
	return count
}

func (d *authRepository) GetUserDetails(email string) (*resposemodel.UserData, error) {
	userDetails := resposemodel.UserData{}
	query := "SELECT * FROM users WHERE email= $1"
	result := d.DB.Raw(query, email).Scan(&userDetails)
	if result.RowsAffected == 0 {
		return nil, errors.New("no user exit with same email")
	}
	return &userDetails, nil
}
