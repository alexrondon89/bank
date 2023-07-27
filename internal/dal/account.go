package dal

import (
	"bank/internal/dal/repository"
	"bank/internal/dal/repository/model"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type accountConn interface {
	Exec(ctx context.Context, query string, args ...interface{}) (int64, error)
	ExecAccountQuery(ctx context.Context, query string, args ...interface{}) (model.Account, error)
	SelectAccountQuery(ctx context.Context, query string, args ...interface{}) ([]model.Account, error)
}

type DalAccount struct {
	db accountConn
}

func NewAccountDb(db repository.PgConnection) DalAccount {
	return DalAccount{
		db: db,
	}
}

func (con *DalAccount) CreateAccount(ctx context.Context, input model.Account) (model.Account, error) {
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

func (con *DalAccount) GetAccounts(ctx context.Context, input map[string]interface{}) ([]model.Account, error) {
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

	accounts, err := con.db.SelectAccountQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// todo this should be update balance method
func (con *DalAccount) UpdateAccount(ctx context.Context, input model.Account) (model.Account, error) {
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

	resp, err := con.db.ExecAccountQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (con *DalAccount) DeleteAccount(ctx context.Context, input model.Account) error {
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

	_, err := con.db.Exec(ctx, query, args)
	return err
}
