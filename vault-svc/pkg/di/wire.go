package di

import (
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/db"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/service"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase"
)

func InitializeService(config *config.Config) *service.Service {

	connections := db.DBConnection(config)

	vaultRepository := repository.NewVaultRepository(connections)
	vaultUsecase := usecase.NewVaultUseCase(vaultRepository)
	service := service.NewVaultService(vaultUsecase)

	return service
}
