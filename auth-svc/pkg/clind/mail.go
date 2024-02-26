package clind

import (
	"context"
	"fmt"
	"log"

	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/pb"
	usecaseinterface "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase/interface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MailServiceClindStruct struct {
	Clind   pb.MailServiceClient
	useCase usecaseinterface.IMailUsecase
}

func InitReminderServiceClind(url string, usecase usecaseinterface.IMailUsecase) *MailServiceClindStruct {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("can't make connection")
	}

	c := MailServiceClindStruct{
		Clind:   pb.NewMailServiceClient(cc),
		useCase: usecase,
	}
	return &c
}

func (m *MailServiceClindStruct) SendMail() {

	allUsers, err := m.useCase.GetUsers()
	if err != nil {
		fmt.Println("--error", err)
	}

	fmt.Println("---", allUsers)

	for _, val := range *allUsers {
		result, err := m.Clind.SendMail(context.Background(), &pb.SendMailRequest{
			UserID: val.UserID,
			Email:  val.Email,
		})
		log.Printf(result.String(), err)
	}

	fmt.Println("remainder mail send succesfully")
}
