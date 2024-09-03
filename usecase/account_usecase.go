package usecase

import (
	"github.com/Mateus-Camillo/Transaction-REST-API/model"
	"github.com/Mateus-Camillo/Transaction-REST-API/repository"
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
	if len(account.Credential) < 8 {
		return model.Account{}, model.ErrInvalidPassword
	}

	accountId, err := au.repository.CreateAccount(account)
	if err != nil {
		return model.Account{}, err
	}

	account.ID = accountId

	return account, nil
}
