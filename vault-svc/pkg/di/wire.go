package di

import (
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/config"
	"github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/service"
)

func InitializeService(config *config.Config) *service.Service {
	service := service.NewVaultService()
	return service
}
