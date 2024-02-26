package cronjob

import (
	"fmt"
	"log"
	"os"

	"github.com/robfig/cron/v3"
	repositoryInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository/interface"
)

func CronJob(vaultRepository repositoryInterface.IVaultRepository) {

	logFile, err := os.OpenFile("cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("log file creation error :", err)
	}
	log.SetOutput(logFile)

	c := cron.New()
	_, err = c.AddFunc("0 0 * * *", vaultRepository.DeleteDataCronJob)
	if err != nil {
		fmt.Println("error at attaching fucntion to cronejob")
	}
	c.Start()
	fmt.Println("--cron start")
}
