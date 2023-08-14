package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"log"
)

type DbConnection struct {
	conn *sqlx.DB
}

func NewConnection() DbConnection {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln("error creating db connection: ", err)
	}

	return DbConnection{
		conn: db,
	}
}

func (db DbConnection) BeginTx(ctx context.Context) (*sql.Tx, error) {
	trx, err := db.conn.BeginTx(ctx, nil)
	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		return nil, err
	}
	return trx, nil
}
