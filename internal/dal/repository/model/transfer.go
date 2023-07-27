package model

import "time"

type Transfer struct {
	Id            int32     `json:"id"`
	FromAccountId int32     `json:"fromAccountId"`
	ToAccountId   int32     `json:"toAccountId"`
	Amount        int32     `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
}
