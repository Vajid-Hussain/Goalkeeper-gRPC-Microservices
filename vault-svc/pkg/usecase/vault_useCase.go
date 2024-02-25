package usecase

import (
	requestmodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/requestModel"
	resposemodel "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/models/resposeModel"
	repositoryInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/repository/interface"
	usecaseInterface "github.com/vajid-hussain/grpc-microservice-vault-svc/pkg/usecase/interface"
)

type vaultUseCase struct {
	vaultRepository repositoryInterface.IVaultRepository
}

func NewVaultUseCase(repo repositoryInterface.IVaultRepository) usecaseInterface.IVaultUsecase {
	return &vaultUseCase{vaultRepository: repo}
}

func (d *vaultUseCase) CreateCategory(data requestmodel.Colleciton) (string, error) {
	objID, err := d.vaultRepository.CreateCollection(data)
	if err != nil {
		return "", err
	}
	return objID, nil
}

func (d *vaultUseCase) InserDAta(data requestmodel.Data) (string, error) {
	return d.vaultRepository.InsertData(data)
}

func (d *vaultUseCase) GetCategories(userID string) ([]*resposemodel.Collections, error) {
	return d.vaultRepository.GetCategories(userID)
}

func (d *vaultUseCase) GetDatas(details requestmodel.GetDataRequest) (*[]resposemodel.GetDataResponse, error) {
	return d.vaultRepository.GetDatas(details)
}
