package repository

import (
	"database/sql"
	"fmt"

	"github.com/Mateus-Camillo/Transaction-REST-API/model"
)

type AccountRepository struct {
	connection *sql.DB
}

func NewAccountRepository(connection *sql.DB) AccountRepository {
	return AccountRepository{
		connection: connection,
	}
}

func (ar *AccountRepository) CreateAccount(account model.Account) (int, error) {
	var id int

	query, err := ar.connection.Prepare("INSERT INTO bank" +
		"(username, cpf, credential, balance)" +
		" VALUES ($1, $2, $3, $4) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(account.Name, account.Cpf, account.Credential, account.Balance).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
