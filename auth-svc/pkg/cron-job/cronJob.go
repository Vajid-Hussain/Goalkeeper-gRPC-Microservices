package cron

import (
	"fmt"
	"log"
	"os"
	"github.com/robfig/cron/v3"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/clind"
)

func CronJob(clindFuntions *clind.MailServiceClindStruct) {

	logFile, err := os.OpenFile("cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("log file creation error :", err)
	}
	log.SetOutput(logFile)

	c := cron.New()
	_, err = c.AddFunc("0 0 * * *", clindFuntions.SendMail)
	if err != nil {
		fmt.Println("error at attaching fucntion to cronejob")
	}

	c.Start()
	fmt.Println("--cron start")
}
