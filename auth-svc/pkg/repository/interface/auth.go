package repositoryIinterfaces

import (
	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
)

type IAuthRepository interface {
	CreateUser(*reqestmodel.User) (*resposemodel.UserData, error)
	UserExist(string) int
	GetUserDetails(string) (*resposemodel.UserData, error)
}
