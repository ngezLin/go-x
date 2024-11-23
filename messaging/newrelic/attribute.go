package newrelic

import (
	"context"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type SegmentAttribute func(segment *newrelic.Segment)
type DistributedTracingAttribute struct {
	http.Header
}

func AddSegmentAttrs(ctx context.Context, attrs ...SegmentAttribute) context.Context {
	segment := SegmentFromContext(ctx)
	if segment == nil {
		return ctx
	}
	for _, attr := range attrs {
		attr(segment)
	}
	return SetSegmentToContext(ctx, segment)
}

func WithSegmentAttr(k string, v any) SegmentAttribute {
	return func(segment *newrelic.Segment) {
		segment.AddAttribute(k, v)
	}
}
