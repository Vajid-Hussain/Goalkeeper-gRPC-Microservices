package clind

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReminderServiceClindStruct struct {
	Clind pb.RemainderServiceClient
}

func InitReminderServiceClient(url string) ReminderServiceClindStruct {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("connection dial or simply connection is not established")
	}

	c := ReminderServiceClindStruct{
		Clind: pb.NewRemainderServiceClient(cc),
	}
	return c
}

func (r *ReminderServiceClindStruct) DailyTask() {
	fmt.Println("----dailyTask")
	result, err := r.Clind.TodayTask(context.Background(), &pb.TodayTaskRequest{
		UserID: "15",
	})
	if err != nil {
		return
	}
	fmt.Println("===", result.Data[0],"9999", result)
}
