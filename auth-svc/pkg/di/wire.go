package di

import (
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/clind"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/cron-job"

	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/db"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/service"
	"github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase"
)

func InitializeService(config *config.Config) *service.Service {
	DB := db.DBInit(config)

	authRepository := repository.NewAuthRepository(DB)
	authUseCase := usecase.NewAuthUsecase(authRepository, config)
	service := service.NewAuthServices(authUseCase)

	mailRepository := repository.NewMailRepository(DB)
	mailUseCase := usecase.NewMailUsecase(mailRepository)
	clind := clind.InitReminderServiceClind(config.MailSvcPort, mailUseCase)

	cron.CronJob(clind)

	return service
}
	