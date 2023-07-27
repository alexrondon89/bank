package repository

import (
	"bank/internal/dal/repository/model"
	"context"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"log"
)

type PgConnection struct {
	conn *sqlx.DB
}

func NewConnection() PgConnection {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln("error creating db connection: ", err)
	}

	return PgConnection{
		conn: db,
	}
}

func (db PgConnection) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
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

func (db PgConnection) ExecAccountQuery(ctx context.Context, query string, args ...interface{}) (model.Account, error) {
	row := db.conn.QueryRowContext(ctx, query, args)
	var newAccount model.Account
	err := row.Scan(&newAccount)
	if err != nil {
		return newAccount, err
	}

	return newAccount, nil
}

func (db PgConnection) SelectAccountQuery(ctx context.Context, query string, args ...interface{}) ([]model.Account, error) {
	rows, err := db.conn.QueryxContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	var accounts []model.Account
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		err = rows.Scan(&account)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
