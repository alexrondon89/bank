package dal

import (
	"bank/internal/dal/repository/model"
	"context"
	"database/sql"
)

type StoreDbInterface interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	ExecTrx(trx *sql.Tx, query string, args ...interface{}) (int64, error)
}

type DalStore struct {
	db StoreDbInterface
}

func NewStore(db StoreDbInterface) DalStore {
	return DalStore{
		db: db,
	}
}

func (ds DalStore) TransferTx(ctx context.Context, input model.TransferTxParams) (interface{}, error) {
	trx, err := ds.db.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	args := []interface{}{input.FromAccountID, input.ToAccountID, input.Amount}
	const createTransfer = `
		INSERT INTO transfers (
		  from_account_id,
		  to_account_id,
		  amount
		) VALUES (
		  $1, $2, $3
		) RETURNING id, from_account_id, to_account_id, amount, created_at
		`
	_, err = ds.db.ExecTrx(trx, createTransfer, args)
	if err != nil {
		return nil, err
	}

	argsEntryFrom := []interface{}{input.FromAccountID, input.Amount}
	const createEntry = `
		INSERT INTO entries (
		  account_id,
		  amount
		) VALUES (
		  $1, $2
		) RETURNING id, account_id, amount, created_at
		`
	_, err = ds.db.ExecTrx(trx, createEntry, argsEntryFrom)
	if err != nil {
		return nil, err
	}

	argsEntryTo := []interface{}{input.FromAccountID, input.Amount}
	_, err = ds.db.ExecTrx(trx, createEntry, argsEntryTo)
	if err != nil {
		return nil, err
	}

	err = trx.Commit()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
