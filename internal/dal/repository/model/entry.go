package model

import "time"

type EntryResult struct {
	Id        int32     `json:"id"`
	AccountId int32     `json:"accountId"`
	Amount    int32     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type EntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}
