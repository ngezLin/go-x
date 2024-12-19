package ctxval

import "context"

func GetCorrelationId(ctx context.Context) string {
	if v, ok := ctx.Value(correlationIdKey{}).(string); ok {
		return v
	}
	return ""
}

func GetTraceParent(ctx context.Context) string {
	if v, ok := ctx.Value(traceParentKey{}).(string); ok {
		return v
	}
	return ""
}
func GetSpanID(ctx context.Context) string {
	if v, ok := ctx.Value(spanIDKey{}).(string); ok {
		return v
	}
	return ""
}

func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.Value(traceIDKey{}).(string); ok {
		return v
	}
	return ""
}

func GetTraceSampled(ctx context.Context) bool {
	if v, ok := ctx.Value(traceSampledKey{}).(bool); ok {
		return v
	}
	return false
}

func GetUserAgent(ctx context.Context) string {
	if v, ok := ctx.Value(userAgentKey{}).(string); ok {
		return v
	}
	return ""
}

func GetHost(ctx context.Context) string {
	if v, ok := ctx.Value(hostKey{}).(string); ok {
		return v
	}
	return ""
}

func GetIP(ctx context.Context) string {
	if v, ok := ctx.Value(ipKey{}).(string); ok {
		return v
	}
	return ""
}

func GetForwardedFor(ctx context.Context) string {
	if v, ok := ctx.Value(forwardedForKey{}).(string); ok {
		return v
	}
	return ""
}

func GetPid(ctx context.Context) string {
	if v, ok := ctx.Value(pidKey{}).(string); ok {
		return v
	}
	return ""
}
