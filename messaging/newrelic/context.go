package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type contextKeyType struct{}

var (
	segmentContextKey = contextKeyType(struct{}{})
)

func SegmentFromContext(ctx context.Context) *newrelic.Segment {
	segment, ok := ctx.Value(segmentContextKey).(*newrelic.Segment)
	if !ok {
		return nil
	}
	return segment
}

func SetSegmentToContext(ctx context.Context, segment *newrelic.Segment) context.Context {
	return context.WithValue(ctx, segmentContextKey, segment)
}
