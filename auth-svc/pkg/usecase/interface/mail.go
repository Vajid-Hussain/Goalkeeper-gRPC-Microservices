package usecaseinterface

import resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"

type IMailUsecase interface {
	GetUsers() (*[]resposemodel.MailUserDetils, error)
}
