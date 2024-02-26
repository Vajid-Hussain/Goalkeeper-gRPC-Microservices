package usecase

import (
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
	repositoryInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository/interface"
	usecaseInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase/interface"
)

type mailUseCase struct {
	mailRepository repositoryInterface.IMailRepository
}

func NewMailUseCase(repo repositoryInterface.IMailRepository) usecaseInterface.IMailUseCase{
	return &mailUseCase{mailRepository: repo}
}

func (r *mailUseCase) TodayTask(userID string) ([]*resposemodel.TodayTask, error) {
	return r.mailRepository.TodayTask(userID)
}
