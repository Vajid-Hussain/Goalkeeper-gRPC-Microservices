package repositoryInterface

import resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"

type IMailRepository interface {
	TodayTask(string) ([]*resposemodel.TodayTask, error)
}
