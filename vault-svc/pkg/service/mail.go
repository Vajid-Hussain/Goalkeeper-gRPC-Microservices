package service

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/pb"
	usecaseInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase/interface"
)

type MailService struct {
	useCase usecaseInterface.IMailUseCase
	pb.UnimplementedRemainderServiceServer
}

func NewMailService(usecase usecaseInterface.IMailUseCase) *MailService {
	return &MailService{useCase: usecase}
}

func (u *MailService) TodayTask(ctx context.Context, req *pb.TodayTaskRequest) (*pb.TodayTaskResponse, error) {
	fmt.Println("today task work well::", req.UserID)

	result, err := u.useCase.TodayTask(req.UserID)
	if err != nil {
		return nil, err
	}

	var finalResult = new(pb.TodayTaskResponse)
	for _, val := range result {
		fmt.Println("---", val)
		finalResult.Data = append(finalResult.Data, val.Data)
	}
	fmt.Println("---", finalResult, result)
	return finalResult, nil
}
