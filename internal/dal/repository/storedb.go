package repository

import (
	"database/sql"
	"fmt"
)

func (db DbConnection) ExecTrx(trx *sql.Tx, query string, args ...interface{}) (int64, error) {
	result, err := trx.Exec(query, args...)
	if err != nil {
		if rbErr := trx.Rollback(); rbErr != nil {
			return -1, fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return -1, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowAffected, nil
}
