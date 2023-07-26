package repository

import (
	"bank/bank-infrastructure/internal/repository/model"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (db *dbConnection) CreateAccount(ctx context.Context, input model.Account) (model.Account, error) {
	args := []interface{}{input.Owner, input.Balance, input.Currency}
	query := `
		INSERT INTO public.accounts (
		  owner,
		  balance,
		  currency
		) VALUES (
		  $1, $2, $3
		) RETURNING *;
	`

	row := db.conn.QueryRowxContext(ctx, query, args)
	var newAccount model.Account
	err := row.Scan(&newAccount)
	if err != nil {
		log.Fatalln("error scanning new account added: ", err)
	}
	return newAccount, err
}

func (db *dbConnection) GetAccounts(ctx context.Context, input map[string]interface{}) ([]model.Account, error) {
	query := `
			SELECT 
			    id,
				owner,
				balance,
				currency,
				createdAt
			FROM public.accounts 
	`

	if ids, ok := input["ids"]; ok {
		var strIds []string
		for _, value := range ids.([]int32) {
			strIds = append(strIds, strconv.Itoa(int(value)))
		}
		query += `
			WHERE ids IN (` + strings.Join(strIds, ", ") + `)
		`
	}

	limit := `1`
	if value, ok := input["limit"].(string); ok {
		limit = value
	}

	query += `LIMIT ` + limit

	if offset, ok := input["offset"].(string); ok {
		query += `
			OFFSET ` + offset
	}

	rows, err := db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var accounts []model.Account
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		err := rows.Scan(&account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (db *dbConnection) UpdateAccount(ctx context.Context, input model.Account) (model.Account, error) {
	if input.Balance == nil {
		return model.Account{}, errors.New("balance is a value needed")
	}

	var args []interface{}
	var set []string

	set = append(set, `balance = $1`)
	args = append(args, *input.Balance, *input.Id)
	query := fmt.Sprintf(`
		UPDATE	
			public.accounts
		SET
			%s
		WHERE
			id = $2
		RETURNING
			id,
			owner,
			balance,
			currency,
			createdAt
	`, strings.Join(set, ", "))

	row, err := db.conn.QueryContext(ctx, query, args)
	if err != nil {
		return model.Account{}, err
	}

	var accountUpdated model.Account
	err = row.Scan(&accountUpdated)
	if err != nil {
		log.Fatalln("error scanning new account added: ", err)
	}

	return accountUpdated, err
}

func (db *dbConnection) DeleteAccount(ctx context.Context, input model.Account) error {
	if input.Id == nil {
		return errors.New("id is a value needed")
	}

	var args []interface{}
	args = append(args, *input.Id)
	query := `
		DELETE FROM	
			public.accounts
		WHERE
			id = $1
		`
	_, err := db.conn.QueryContext(ctx, query, args)
	return err
}
