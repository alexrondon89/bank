package model

import "time"

type AccountParams struct {
	Ids      []int32 `json:"id"`
	Owner    *string `json:"owner"`
	Balance  *int32  `json:"balance"`
	Currency *string `json:"currency"`
	Limit    *int    `json:"limit"`
	Offset   *int    `json:"offset"`
}

type AccountResult struct {
	Id        *int32     `json:"id"`
	Owner     *string    `json:"owner"`
	Balance   *int32     `json:"balance"`
	Currency  *string    `json:"currency"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
