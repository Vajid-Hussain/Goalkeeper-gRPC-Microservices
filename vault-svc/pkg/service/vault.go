package service

import (
	"context"
	"fmt"

	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/pb"
	usecaseInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase/interface"
)

type Service struct {
	vaultUsecase usecaseInterface.IVaultUsecase
	pb.UnimplementedVaultServiceServer
}

func NewVaultService(usecase usecaseInterface.IVaultUsecase) *Service {
	return &Service{vaultUsecase: usecase}
}

func (s *Service) CreateCollection(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	fmt.Println("----", req.CollectionName, req.UserID)
	var data = requestmodel.Colleciton{CategoryName: req.CollectionName, UserID: req.UserID}

	objID, err := s.vaultUsecase.CreateCategory(data)
	if err != nil {
		return &pb.CreateCollectionResponse{}, err
	}

	return &pb.CreateCollectionResponse{
		CollectionID:  objID,
		CollecionName: req.CollectionName,
		UserID:        req.UserID,
	}, nil
}

func (s *Service) InserData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	data := requestmodel.Data{}
	data.CategoryID = req.CategoryID
	data.Data = req.Data
	data.Reminder = req.Reminder
	data.UserID = req.UserID

	result, err := s.vaultUsecase.InserDAta(data)
	if err != nil {
		return &pb.DataResponse{}, err
	}

	return &pb.DataResponse{
		DataID: result,
		Data:   &pb.DataRequest{UserID: data.UserID, CategoryID: data.CategoryID, Data: data.Data},
	}, nil
}

// func (s *Service) GetCollection(userID string) {

// }