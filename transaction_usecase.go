package usecase

import (
	"example.com/m/dto"
	"example.com/m/entity"
	"example.com/m/repository"
)

type TransactionUsecase interface {
	FetchMachineUtil(assetReq dto.MachineTransactionRequest) (*[]entity.Machine, error)
}

type transactionUsecaseImpl struct {
	transactionRepo repository.TransactionRepo
}

type TransactionUsecaseConfig struct {
	TransactionRepo repository.TransactionRepo
}

func NewTransactionUsecase(t TransactionUsecaseConfig) TransactionUsecase {
	return &transactionUsecaseImpl{
		transactionRepo: t.TransactionRepo,
	}
}

func (t *transactionUsecaseImpl) FetchMachineUtil(assetReq dto.MachineTransactionRequest) (*[]entity.Machine, error) {

	machine, err := t.transactionRepo.FetchMachineUtil(assetReq)

	if err != nil {
		return nil, err
	}

	return machine, err
}
