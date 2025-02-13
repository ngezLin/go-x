package ctxval

import "context"

func SetCorrelationId(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, correlationIdKey{}, v)
	}
}

func SetTraceParent(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceParentKey{}, v)
	}
}

func SetSpanID(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, spanIDKey{}, v)
	}
}

func SetTraceID(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceIDKey{}, v)
	}
}

func SetTraceSampled(v bool) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceSampledKey{}, v)
	}
}

func SetUserAgent(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userAgentKey{}, v)
	}
}

func SetHost(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, hostKey{}, v)
	}
}

func SetIP(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ipKey{}, v)
	}
}

func SetForwardedFor(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, forwardedForKey{}, v)
	}
}

func SetPid(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pidKey{}, v)
	}
}

func SetIdempotency(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, idempotencyKey{}, v)
	}
}

func SetUserDevice(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userDeviceKey{}, v)
	}
}
