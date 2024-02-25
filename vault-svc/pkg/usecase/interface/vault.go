package usecaseInterface

import (
	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
)

type IVaultUsecase interface {
	CreateCategory(requestmodel.Colleciton) (string, error)
	InserDAta(requestmodel.Data) (string, error)
	GetCategories(string) ([]*resposemodel.Collections, error)
	GetDatas(requestmodel.GetDataRequest) (*[]resposemodel.GetDataResponse, error)
}
