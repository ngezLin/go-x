package newrelic

import (
	"context"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func BeforePublish(ctx context.Context, name string) (context.Context, DistributedTracingAttribute) {
	dt := DistributedTracingAttribute{
		Header: make(http.Header),
	}

	tx := newrelic.FromContext(ctx)
	if tx == nil {
		return ctx, dt
	}

	tx.InsertDistributedTraceHeaders(dt.Header)

	segment := tx.StartSegment(name)

	return SetSegmentToContext(ctx, segment), dt
}

func AfterPublish(ctx context.Context) {
	if segment := SegmentFromContext(ctx); segment != nil {
		segment.End()
	}
}
