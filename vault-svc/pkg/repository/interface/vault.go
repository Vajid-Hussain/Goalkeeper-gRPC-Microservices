package repositoryInterface

import requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"

type IVaultRepository interface {
	CreateCollection(requestmodel.Colleciton) (string, error)
	InsertData( requestmodel.Data) (string, error)
}
