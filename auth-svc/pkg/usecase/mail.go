package usecase

import (
	resposemodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/resposeModel"
	repositoryIinterfaces "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/repository/interface"
	usecaseinterface "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/usecase/interface"
)

type mailUsecase struct {
	mailRepository repositoryIinterfaces.IMailRepository
}

func NewMailUsecase(repository repositoryIinterfaces.IMailRepository) usecaseinterface.IMailUsecase{
	return &mailUsecase{mailRepository: repository}
}

func (r *mailUsecase) GetUsers() (*[]resposemodel.MailUserDetils, error) {
	return r.mailRepository.FetchUserDetails()
}