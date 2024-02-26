package main

import (
	"log"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/clind"
	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("env is not loadded, ", err)
	}

	c := clind.InitReminderServiceClient(config.VaultSvcPort)
	c.DailyTask()
}
