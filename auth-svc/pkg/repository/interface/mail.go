package repositoryIinterfaces

import resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"

type IMailRepository interface {
	FetchUserDetails() (*[]resposemodel.MailUserDetils, error)
}
