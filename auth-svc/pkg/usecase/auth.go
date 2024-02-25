package usecase

import (
	"errors"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/config"
	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
	repositoryIinterfaces "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository/interface"
	usecaseinterface "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase/interface"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/utils"
)

type authUsecase struct {
	authRepository repositoryIinterfaces.IAuthRepository
	config         *config.Config
}

func NewAuthUsecase(repository repositoryIinterfaces.IAuthRepository, c *config.Config) usecaseinterface.IAuthUsecase {
	return &authUsecase{authRepository: repository,
		config: c}
}

func (r *authUsecase) UserCreate(userData reqestmodel.User) (*resposemodel.UserData, error) {
	var emailCredential = reqestmodel.Email{From: r.config.FromEmail, To: userData.Email, AppPasskey: r.config.AppPasskey}

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

	result.Jwt = utils.CreateJwt(r.config.JwtSecret, result.ID)

	utils.GreetingMain(emailCredential)

	return result, nil
}

func (r *authUsecase) Login(userData reqestmodel.User) (*resposemodel.UserData, error) {
	userDetails, err := r.authRepository.GetUserDetails(userData.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println("&&&", userDetails)
	result := utils.CompareHast(userData.Password, userDetails.Password)
	if !result {
		return nil, errors.New("password is mismatch")
	}

	userDetails.Jwt = utils.CreateJwt(r.config.JwtSecret, userDetails.ID)
	return userDetails, nil
}

func (r *authUsecase) VerifyJwtToken(token string) (string, error) {
	return utils.VerifyJwt(token, r.config.JwtSecret)
}
