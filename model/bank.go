package model

type Account struct {
	ID         int     `json:"id_user"`
	Name       string  `json:"name"`
	Cpf        string  `json:"cpf"`
	Credential string  `json:"credential"`
	Balance    float64 `json:"balance"`
}

type Transfer struct {
	FromAccountID int     `json:"from_account_id"` 
	ToAccountID   int     `json:"to_account_id"`   
	Amount        float64 `json:"amount"`
}

type TransferResponse struct {
	Account string 
	Amount float64
}

type Response struct {
	Message string 
}