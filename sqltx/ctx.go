package sqltx

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
)

type Sql interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

const transaction = "sqltx.trx"
const trxDefiner = "sqltx.trx.definer"

func GetCaller() string {
	var fn = ""
	for i := 3; i < 97; i++ {
		_, _, fn := CallerTraceWithPosition(i)
		if !slices.Contains([]string{
			"SetTransaction",
			"TransactWithOptions",
			"Transact",
		}, fn) {
			break
		}
	}
	return fn
}
func SetTransaction[T any](ctx context.Context, tx T) context.Context {
	ctx = context.WithValue(ctx, transaction, tx)
	ctx = context.WithValue(ctx, trxDefiner, GetCaller())
	return ctx
}

func GetTransaction[T any](ctx context.Context) (T, error) {
	v, ok := ctx.Value(transaction).(T)
	if !ok {
		return v, fmt.Errorf("transaction not found from context!")
	}
	return v, nil
}

func GetTransactionDefiner(ctx context.Context) string {
	v, _ := ctx.Value(trxDefiner).(string)
	return v
}
