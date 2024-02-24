package service

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/pb"
)

type Service struct {
	pb.UnimplementedVaultServiceServer
}

func NewVaultService() *Service {
	return &Service{}
}

func (s *Service) CreateCollection(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	fmt.Println("----", req.CollectionName, req.UserID)

	return &pb.CreateCollectionResponse{
		CollecionName: req.CollectionName,
		UserID:        req.UserID,
	}, nil
}
