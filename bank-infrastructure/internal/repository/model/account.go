package model

import "time"

type Account struct {
	Id        *int32     `json:"id"`
	Owner     *string    `json:"owner"`
	Balance   *int32     `json:"balance"`
	Currency  *string    `json:"currency"`
	CreatedAt *time.Time `json:"createdAt"`
}
