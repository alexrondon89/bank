package model

import "time"

type TransferResult struct {
	Id            int32     `json:"id"`
	FromAccountId int32     `json:"fromAccountId"`
	ToAccountId   int32     `json:"toAccountId"`
	Amount        int32     `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"UpdatedAt"`
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    TransferResult `json:"transfer"`
	FromAccount AccountResult  `json:"from_account"`
	ToAccount   AccountResult  `json:"to_account"`
	FromEntry   EntryResult    `json:"from_entry"`
	ToEntry     EntryResult    `json:"to_entry"`
}
