package usecaseinterface

import (
	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
)

type IAuthUsecase interface {
	UserCreate(reqestmodel.User) (*resposemodel.UserData, error)
	Login(reqestmodel.User) (*resposemodel.UserData, error)
	VerifyJwtToken(string) (string, error)
}
