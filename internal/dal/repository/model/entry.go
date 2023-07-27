package model

import "time"

type Entry struct {
	Id        int32     `json:"id"`
	AccountId int32     `json:"accountId"`
	Amount    int32     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}
