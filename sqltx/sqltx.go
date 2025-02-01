package sqltx

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	xlog "github.com/super-saga/go-x/log"
)

// DefaultRetriesOptions will be retryable for serializable isolation level error only
var DefaultRetriesOptions = RetryOptions{
	IsRetryable: SerizalizeError,
	MaxRetries:  3,
	BackOff: func() time.Duration {
		return time.Millisecond * 150
	},
}

// DefaultMaxRetries configures the default number of max retries attempted by TransactWithRetry.
var DefaultMaxRetries = 5

// DefaultIsRetryable configures the default function for determining whether the error returned from the operation is
// retryable. By default, all errors are retryable. A RollbackErr is never retryable..
var DefaultIsRetryable = alwaysRetryable

// Transact executes the operation inside a transaction, committing the transaction on completion. If the operation
// returns an error or panic, the transaction will be rolled back, returning the original error or propagating the
// original panic. If the rollback caused by an error also receives an error, a RollbackErr will be returned. If the
// rollback caused by a panic returns an error, the error message and original panic merged and propagated as a new
// panic.
func Transact(ctx context.Context, conn Database, operation func(ctx context.Context) error) (err error) {
	return TransactWithOptions(ctx, conn, nil, operation)
}

// TransactWithOptions executes the operation inside a transaction, committing the transaction on completion. If the
// operation returns an error or panic, the transaction will be rolled back, returning the original error or propagating
// the original panic. If the rollback caused by an error also receives an error, a RollbackErr will be returned. If the
// rollback caused by a panic returns an error, the error message and original panic merged and propagated as a new
// panic.
//
// The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is
// used that the driver doesn't support, an error will be returned.
func TransactWithOptions(ctx context.Context, conn Database, txOpts *sql.TxOptions, operation func(ctx context.Context) error) (err error) {
	var (
		child bool = true
		tx    Tx
	)
	_, err = GetTransaction(ctx)
	if err != nil {
		tx, err = conn.BeginTx(ctx, txOpts)
		if err != nil {
			return err
		}
		ctx = SetTransaction(ctx, tx)
		child = false
	}

	if child {
		xlog.Info(ctx, "[DATABASE.TRANSACTION.PUSH.CHILD]")
	} else {
		xlog.Info(ctx, "[DATABASE.TRANSACTION.BEGIN]")
	}

	defer func() {
		if !child {
			if p := recover(); p != nil {
				tx.Rollback()
				err = fmt.Errorf("panic happened because: " + fmt.Sprintf("%v", p))
				xlog.Panic(ctx, "[DATABASE.TRANSACTION.PANIC]", xlog.Err(err))
			} else if err != nil {
				tx.Rollback()
				xlog.Warn(ctx, "[DATABASE.TRANSACTION.ROLLBACK]", xlog.Err(err))
			} else {
				err = tx.Commit()
				xlog.Info(ctx, "[DATABASE.TRANSACTION.COMMIT]")
			}
		} else {
			xlog.Info(ctx, "[DATABASE.TRANSACTION.POP.CHILD]")
		}
	}()
	err = operation(ctx)
	return
}

// TransactWithRetry runs the operation using Transact, performing retries according to RetryOptions. If all retries
// fail, the error from the last attempt will be returned. If a rollback fails, no further attempts will be made and the
// RollbackErr will be returned.
//
// Since the transaction operation may be executed multiple times, it is important that any mutations it applies
// to application state (outside the database) be idempotent.
func TransactWithRetry(ctx context.Context, conn *sql.DB, retry *RetryOptions, operation func(ctx context.Context) error) error {
	return TransactWithRetryAndOptions(ctx, conn, nil, retry, operation)
}

// TransactWithRetryAndOptions runs the operation using Transact, performing retries according to RetryOptions. If all
// retries fail, the error from the last attempt will be returned. If a rollback fails, no further attempts will be made
// and the RollbackErr will be returned.
//
// Since the transaction operation may be executed multiple times, it is important that any mutations it applies to
// application state (outside the database) be idempotent.
//
// The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is
// used that the driver doesn't support, an error will be returned.
func TransactWithRetryAndOptions(ctx context.Context, conn *sql.DB, txOpts *sql.TxOptions, retry *RetryOptions, operation func(ctx context.Context) error) error {
	if retry == nil {
		retry = &DefaultRetriesOptions
	}

	if retry.MaxRetries == 0 {
		retry.MaxRetries = DefaultMaxRetries
	}
	if retry.MaxRetries < 0 {
		retry.MaxRetries = math.MaxInt32
	}
	if retry.BackOff == nil {
		retry.BackOff = newSimpleExponentialBackOff().NextBackOff
	}
	if retry.IsRetryable == nil {
		retry.IsRetryable = DefaultIsRetryable
	}
	if retry.After == nil {
		retry.After = time.After
	}

	var err error
	for i := 0; i < retry.MaxRetries; i++ {
		err = TransactWithOptions(ctx, conn, txOpts, operation)
		if err == nil {
			return nil
		}
		if !retry.IsRetryable(err) {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-retry.After(retry.BackOff()):
			continue
		}
	}
	return err
}

// RollbackErr is the error returned if the transaction operation returned an error, and the rollback automatically
// attempted also returns an error.
type RollbackErr struct {
	// The original error that the operation returned.
	OriginalErr error
	// The error returned by sql.Tx.Rollback()
	Err error
}

// Unwrap returns the OriginalErr.
func (r *RollbackErr) Unwrap() error {
	return r.OriginalErr
}

// Cause returns the OriginalErr.
func (r *RollbackErr) Cause() error {
	return r.Unwrap()
}

// Error returns a formatted error message containing both the OriginalErr and RollbackErr.
func (r *RollbackErr) Error() string {
	return fmt.Sprintf("error rolling back failed transaction: %v, original transaction error: %v", r.Err, r.OriginalErr)
}

// RetryOptions controls how TransactWithRetry behaves.
type RetryOptions struct {
	// MaxRetries configures how many attempts will be made to complete the operation when a retryable error is
	// encountered. The default is DefaultMaxRetries. If set to a negative number, math.MaxInt32 attempts will be made.
	MaxRetries int
	// BackOff is called on each retry, and should return a time.Duration indicating how long to wait before the next
	// attempt. The default is an exponential backoff based on the values of DefaultInitialBackOff, DefaultMaxBackOff,
	// and DefaultBackOffFactor. If a negative Duration is returned by NextBackOff(), retries will be aborted.
	//
	// Most backoff implementations are compatible, including github.com/cenkalti/backoff and
	// github.com/jpillora/backoff.
	BackOff func() time.Duration
	// IsRetryable determines whether the error from the operation should be retried. Return true to retry.
	IsRetryable func(err error) bool
	// After is an optional value to be used for mocking out time.After() for testing. If set, backoff wait
	// will use this function instead of time.After().
	After func(d time.Duration) <-chan time.Time
}

func alwaysRetryable(error) bool {
	return true
}

func SerizalizeError(err error) bool {
	return err.Error() == "ERROR: could not serialize access due to concurrent update (SQLSTATE 40001)"
}
