package usecase

import (
	"errors"

	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
	repositoryIinterfaces "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository/interface"
	usecaseinterface "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase/interface"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/utils"
)

type authUsecase struct {
	authRepository repositoryIinterfaces.IAuthRepository
}

func NewAuthUsecase(repository repositoryIinterfaces.IAuthRepository) usecaseinterface.IAuthUsecase {
	return &authUsecase{authRepository: repository}
}

func (r *authUsecase) UserCreate(userData reqestmodel.User) (*resposemodel.UserData, error) {
	if count := r.authRepository.UserExist(userData.Email); count > 0 {
		return nil, errors.New("user alrady exist try with another mail")
	}

	var err error
	userData.Password, err = utils.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	result, err := r.authRepository.CreateUser(&userData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *authUsecase) Login(userData reqestmodel.User) (*resposemodel.UserData, error) {
	userDetails, err := r.authRepository.GetUserDetails(userData.Email)
	if err != nil {
		return nil, err
	}

	result := utils.CompareHast(userData.Password, userData.Password)
	if !result {
		return nil, errors.New("password is mismatch")
	}
	return userDetails, nil
}