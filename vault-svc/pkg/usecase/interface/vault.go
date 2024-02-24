package usecaseInterface

import requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"

type IVaultUsecase interface {
	CreateCategory(requestmodel.Colleciton)  (string, error)
	InserDAta( requestmodel.Data) (string, error) 
}
