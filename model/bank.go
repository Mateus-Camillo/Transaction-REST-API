package model

type Account struct {
	ID         int     `json:"id_user"`
	Name       string  `json:"name"`
	Cpf        string  `json:"cpf"`
	Credential string  `json:"credential"`
	Balance    float64 `json:"balance"`
}

type Transfer struct {
	FromAccountID int     `json:"from_account_id"` // ID da conta de origem
	ToAccountID   int     `json:"to_account_id"`   // ID da conta de destino
	Amount        float64 `json:"amount"`
}
