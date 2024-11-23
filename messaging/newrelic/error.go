package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NoticeError(ctx context.Context, err error) context.Context {
	tx := newrelic.FromContext(ctx)
	if tx == nil {
		return ctx
	}
	tx.NoticeError(err)

	return newrelic.NewContext(ctx, tx)
}

func NoticeExpectedError(ctx context.Context, err error) context.Context {
	tx := newrelic.FromContext(ctx)
	if tx == nil {
		return ctx
	}
	tx.NoticeExpectedError(err)

	return newrelic.NewContext(ctx, tx)
}
