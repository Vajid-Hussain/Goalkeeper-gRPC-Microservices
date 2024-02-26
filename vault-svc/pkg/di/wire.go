package di

import (
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/config"
	cronjob "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/cron-job"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/db"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/service"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase"
)

func InitializeService(config *config.Config) (*service.Service, *service.MailService) {

	connections := db.DBConnection(config)

	mailRepository := repository.NewMailRepository(connections)
	mailUseCase := usecase.NewMailUseCase(mailRepository)
	mailService := service.NewMailService(mailUseCase)

	vaultRepository := repository.NewVaultRepository(connections)
	vaultUsecase := usecase.NewVaultUseCase(vaultRepository)
	service := service.NewVaultService(vaultUsecase)

	cronjob.CronJob(vaultRepository)

	return service, mailService
}
