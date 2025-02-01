package sqltx

import (
	"context"
)

// Cover nil trx context, use this while your hybrid implementor
func Commiter(ctx context.Context, trx Tx, err *error) {
	_, errCtx := GetTransaction(ctx)
	if errCtx != nil {
		tx, ok := any(trx).(Tx)
		if ok {
			if *err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		}
	}
}
