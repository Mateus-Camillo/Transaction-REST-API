package usecase

import (
	"github.com/Mateus-Camillo/Transaction-REST-API/model"
	"github.com/Mateus-Camillo/Transaction-REST-API/repository"
)

type TransferUseCase struct {
	repository repository.TransferRepository
}

func NewTransferUseCase(repo repository.TransferRepository) TransferUseCase {
	return TransferUseCase{
		repository: repo,
	}
}

func (tu *TransferUseCase) TransferBalance(transfer model.Transfer) error {
	err := tu.repository.TransferBalance(transfer)

	if err != nil {
		return model.Transfer{}, err
	}

	return nil
}