package usecaseInterface

import resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"

type IMailUseCase interface {
	TodayTask(string) ([]*resposemodel.TodayTask, error)
}
