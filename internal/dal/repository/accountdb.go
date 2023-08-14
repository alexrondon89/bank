package repository

import (
	"bank/internal/dal/repository/model"
	"context"
	_ "github.com/lib/pq"
)

func (db DbConnection) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := db.conn.ExecContext(ctx, query, args)
	if err != nil {
		return -1, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return rowAffected, err
	}

	return rowAffected, nil
}

func (db DbConnection) ExecAccountQuery(ctx context.Context, query string, args ...interface{}) (model.AccountResult, error) {
	row := db.conn.QueryRowContext(ctx, query, args)
	var newAccount model.AccountResult
	err := row.Scan(&newAccount)
	if err != nil {
		return newAccount, err
	}

	return newAccount, nil
}

func (db DbConnection) SelectAccountQuery(ctx context.Context, query string, args ...interface{}) ([]model.AccountResult, error) {
	rows, err := db.conn.QueryxContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	var accounts []model.AccountResult
	defer rows.Close()
	for rows.Next() {
		var account model.AccountResult
		err = rows.Scan(&account)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
