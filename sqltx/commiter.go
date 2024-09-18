package sqltx

import (
	"context"
	"database/sql"
)

// Cover nil trx context, use this while your hybrid implementor
func Commiter[T any](ctx context.Context, trx T, err *error) {
	_, errCtx := GetTransaction[T](ctx)
	if errCtx != nil {
		tx, ok := any(trx).(*sql.Tx)
		if ok {
			if *err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		}
	}
}
