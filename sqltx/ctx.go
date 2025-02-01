package sqltx

import (
	"context"
	"database/sql"
	"fmt"
)

type Sql interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	Stmt(stmt *sql.Stmt) *sql.Stmt
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Tx interface {
	Commit() error
	Rollback() error
	Sql
}

const transaction = "sqltx.trx"

func SetTransaction(ctx context.Context, tx Tx) context.Context {
	ctx = context.WithValue(ctx, transaction, tx)
	return ctx
}

func GetTransaction(ctx context.Context) (Tx, error) {
	v, ok := any(ctx.Value(transaction)).(Tx)
	if !ok {
		return v, fmt.Errorf("transaction not found from context!")
	}
	return v, nil
}
