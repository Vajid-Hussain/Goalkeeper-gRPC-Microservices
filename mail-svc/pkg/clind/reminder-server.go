// package clind

// import (
// 	"context"
// 	"fmt"

// 	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type ReminderServiceClients struct {
// 	Client pb.ReminderServiceClient
// }

// func InitReminderServiceClient(port string) ReminderServiceClients {
// 	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		fmt.Println("could not connect")
// 	}

// 	c := ReminderServiceClients{
// 		Client: pb.NewReminderServiceClient(cc),
// 	}
// 	return c
// }

// func (c *ReminderServiceClients) TodayTask() {
// 	result, err := c.Client.TodayTask(context.Background(), &pb.TodayTaskRequst{
// 		UserID: "15",
// 	})
// 	fmt.Println("---", result, err)
// }

package clind

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReminderServiceClindStruct struct {
	Clind pb.ReminderServiceClient
}

func InitReminderServiceClient(url string) ReminderServiceClindStruct {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("connection dial or simply connection is not established")
	}

	c := ReminderServiceClindStruct{
		Clind: pb.NewReminderServiceClient(cc),
	}
	return c
}

func (r *ReminderServiceClindStruct) DailyTask() {
	fmt.Println("----dailyTask")
	result, err := r.Clind.TodayTask(context.Background(), &pb.TodayTaskRequst{
		UserID: "15",
	})

	fmt.Println("=====", result, err)
}
