package ctxval

type (
	correlationIdKey struct{}
	traceParentKey   struct{}
	traceIDKey       struct{}
	spanIDKey        struct{}
	traceSampledKey  struct{}
	idempotencyKey   struct{}

	userAgentKey    struct{}
	userDeviceKey   struct{}
	hostKey         struct{}
	ipKey           struct{}
	forwardedForKey struct{}
	pidKey          struct{}
)
