package service

import (
	"context"
	"fmt"
	"net/http"

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
	fmt.Println("----", req.Email, req.Password)
	var user = reqestmodel.User{Email: req.Email, Password: req.Password}

	userData, err := s.authUseCase.UserCreate(user)
	fmt.Println("===========", userData, err)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("------", req.Email, req.Password)

	return &pb.LoginResponse{
		StatusCode: http.StatusOK,
		Error:      "all perfect",
	}, nil
}
