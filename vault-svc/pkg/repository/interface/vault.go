package repositoryInterface

import (
	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
)

type IVaultRepository interface {
	CreateCollection(requestmodel.Colleciton) (string, error)
	InsertData(requestmodel.Data) (string, error)
	GetCategories(string) ([]*resposemodel.Collections, error)
	GetDatas( requestmodel.GetDataRequest) (*[]resposemodel.GetDataResponse, error)
	DeleteDataCronJob()
}
