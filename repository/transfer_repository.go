package repository

import (
	"database/sql"
	"fmt"
	"github.com/Mateus-Camillo/Transaction-REST-API/model"
)

type TransferRepository struct {
	connection *sql.DB
}

func NewTransferRepository(connection *sql.DB) TransferRepository {
	return TransferRepository{
		connection: connection,
	}
}

func (ar *TransferRepository) TransferBalance(transfer model.Transfer) error {
	tx, err := ar.connection.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	senderQuery := "UPDATE bank SET balance = balance - $1 WHERE username = $2"
	_, err = tx.Exec(senderQuery, transfer.Amount, transfer.Amount)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	receiverQuery, err := ar.connection.Prepare("UPDATE bank SET balance = balance + $1 WHERE username = $2")
	_, err = tx.Exec(receiverQuery, transfer.Amount, transfer.Amount)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	if err = tx.Commit(); err != nil {
        fmt.Println(err)
		return err
    }
	
	return nil
}
