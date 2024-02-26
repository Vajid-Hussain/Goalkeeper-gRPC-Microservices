package server

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/clind"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/model"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/pb"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/utils"
)

type MailService struct {
	RemainderClind *clind.ReminderServiceClindStruct
	pb.UnimplementedMailServiceServer
	EmailDetails *config.Config
}

func MailServerInstance(reminder *clind.ReminderServiceClindStruct, credential *config.Config) *MailService {
	return &MailService{RemainderClind: reminder, EmailDetails: credential}
}

func (s *MailService) SendMail(ctx context.Context, req *pb.SendMailRequest) (*pb.SendMailResponse, error) {

	result, err := s.RemainderClind.Clind.TodayTask(context.Background(), &pb.TodayTaskRequest{
		UserID: req.UserID,
	})
	// fmt.Println("result,", result, err, "lenght", len(result.Data))
	if len(result.Data) == 0 || err != nil {
		return nil, err
	}

	SendMailPreRequirement := model.SendRemainder{
		From:       s.EmailDetails.FromEmail,
		To:         req.Email,
		Data:       result.Data,
		AppPasskey: s.EmailDetails.AppPasskey,
	}
	fmt.Println("=============", SendMailPreRequirement)
	utils.RemainderMain(&SendMailPreRequirement)

	return &pb.SendMailResponse{}, nil
}
	