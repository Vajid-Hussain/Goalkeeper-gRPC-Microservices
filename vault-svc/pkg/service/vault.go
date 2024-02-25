package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
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

func (s *Service) GetCategories(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {

	result, err := s.vaultUsecase.GetCategories(req.UserID)
	if err != nil {
		return &pb.CategoryResponse{}, err
	}

	var categoryList []*pb.CategoryList
	for _, collection := range result {
		list := pb.CategoryList{
			ID:           collection.ID,
			CategoryName: collection.CategoryName,
		}
		categoryList = append(categoryList, &list)
	}

	return &pb.CategoryResponse{Category: categoryList}, nil
}

func (s *Service) GetDatas(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {

	var details requestmodel.GetDataRequest
	details.CategoryID = req.CategoryID
	details.UserID = req.UserID

	result, err := s.vaultUsecase.GetDatas(details)
	if err != nil {
		return &pb.GetDataResponse{}, err
	}

	var dataList []*pb.DatasList
	for _, val := range *result {
		data := pb.DatasList{
			Data:     val.Data,
			Reminder: &timestamp.Timestamp{Seconds: val.Reminder.Unix(), Nanos: int32(val.Reminder.Nanosecond())},
		}
		dataList = append(dataList, &data)
	}

	return &pb.GetDataResponse{Datas: dataList}, nil
}
