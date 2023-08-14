package dal

import (
	"bank/internal/dal/repository/model"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AccountDbInterface interface {
	Exec(ctx context.Context, query string, args ...interface{}) (int64, error)
	ExecAccountQuery(ctx context.Context, query string, args ...interface{}) (model.AccountResult, error)
	SelectAccountQuery(ctx context.Context, query string, args ...interface{}) ([]model.AccountResult, error)
}

type DalAccount struct {
	db AccountDbInterface
}

func NewAccountDb(db AccountDbInterface) DalAccount {
	return DalAccount{
		db: db,
	}
}

func (con *DalAccount) CreateAccount(ctx context.Context, input model.AccountParams) (model.AccountResult, error) {
	args := []interface{}{input.Owner, input.Balance, input.Currency}
	query := `
		INSERT INTO public.accounts (
		  owner,
		  balance,
		  currency
		) VALUES (
		  $1, $2, $3
		) RETURNING	
			id,
			owner,
			balance,
			currency,
			createdAt;
	`
	resp, err := con.db.ExecAccountQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (con *DalAccount) GetAccounts(ctx context.Context, input model.AccountParams) ([]model.AccountResult, error) {
	query := `
			SELECT 
			    id,
				owner,
				balance,
				currency,
				createdAt
			FROM public.accounts 
	`

	if input.Ids != nil {
		var strIds []string
		for _, value := range input.Ids {
			strIds = append(strIds, strconv.Itoa(int(value)))
		}
		query += `
			WHERE ids IN (` + strings.Join(strIds, ", ") + `)
		`
	}

	limit := 10
	if input.Limit != nil {
		limit = *input.Limit
	}

	query += `LIMIT ` + strconv.Itoa(limit)

	if input.Offset != nil {
		query += `
			OFFSET ` + strconv.Itoa(*input.Offset)
	}

	accounts, err := con.db.SelectAccountQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// todo this should update balance method
func (con *DalAccount) UpdateAccount(ctx context.Context, input model.AccountParams) (model.AccountResult, error) {
	if input.Balance == nil {
		return model.AccountResult{}, errors.New("balance is a value needed")
	}

	var args []interface{}
	var set []string

	set = append(set, `balance = $1`)
	args = append(args, *input.Balance, input.Ids[0])
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

	resp, err := con.db.ExecAccountQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (con *DalAccount) DeleteAccount(ctx context.Context, input model.AccountParams) error {
	if len(input.Ids) == 0 {
		return errors.New("id is a value needed")
	}

	var args []interface{}
	args = append(args, input.Ids[0])
	query := `
		DELETE FROM	
			public.accounts
		WHERE
			id = $1
		`

	_, err := con.db.Exec(ctx, query, args)
	return err
}
