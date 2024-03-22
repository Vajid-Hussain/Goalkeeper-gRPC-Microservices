package service

import (
	"context"
	"errors"
	"regexp"

	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/pb"
	usecaseinterface "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase/interface"
)

type Service struct {
	authUseCase usecaseinterface.IAuthUsecase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServices(authUsecase usecaseinterface.IAuthUsecase) *Service {
	return &Service{authUseCase: authUsecase}
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	r, _ := regexp.Compile("p([a-zA-Z0-9!@#$%^&*]+)ch")
	if !r.MatchString(req.Password) {
		return nil, errors.New("password not meat the critiria")
	}

	// fmt.Println("----", req.Email, req.Password)
	var user = reqestmodel.User{Email: req.Email, Password: req.Password}

	userData, err := s.authUseCase.UserCreate(user)
	// fmt.Println("===========", userData, err)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{
		Jwt:    userData.Jwt,
		UserID: userData.ID,
		Email:  userData.Email,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// fmt.Println("------", req.Email, req.Password)
	var loginDetails = reqestmodel.User{Email: req.Email, Password: req.Password}
	result, err := s.authUseCase.Login(loginDetails)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	return &pb.LoginResponse{Login: &pb.RegisterResponse{Jwt: result.Jwt, Email: result.Email, UserID: result.ID}}, nil
}

func (s *Service) JwtValidate(ctx context.Context, req *pb.JwtRequest) (*pb.JwtResponse, error) {
	userID, err := s.authUseCase.VerifyJwtToken(req.Jwt)
	// fmt.Println("*********", req.Jwt, userID, err)
	if err != nil {
		return &pb.JwtResponse{}, err
	}
	return &pb.JwtResponse{UserID: userID}, nil
}
