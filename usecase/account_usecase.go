package usecase

import (
	"github.com/Mateus-Camillo/Transaction-REST-API/model"
	"github.com/Mateus-Camillo/Transaction-REST-API/repository"
	"errors"
)

type AccountUseCase struct {
	repository repository.AccountRepository
}

func NewAccountUseCase(repo repository.AccountRepository) AccountUseCase {
	return AccountUseCase{
		repository: repo,
	}
}

func (au *AccountUseCase) CreateAccount(account model.Account) (model.Account, error) {
	accountId, err := au.repository.CreateAccount(account)

	if err != nil {
		return model.Account{}, err
	}

	if len(account.Credential) < 8 {
		return model.Account{}, ErrInvalidPassword
	}

	account.ID = accountId

	return account, nil
}
