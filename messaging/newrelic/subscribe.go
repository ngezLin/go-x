package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func BeforeSubscribe(ctx context.Context, nr *newrelic.Application, name string, dt DistributedTracingAttribute) context.Context {
	if nr == nil {
		return ctx
	}
	tx := nr.StartTransaction(name)
	tx.AcceptDistributedTraceHeaders(newrelic.TransportQueue, dt.Header)

	return newrelic.NewContext(ctx, tx)
}

func AfterSubscribe(ctx context.Context) {
	if tx := newrelic.FromContext(ctx); tx != nil {
		tx.End()
	}
}
