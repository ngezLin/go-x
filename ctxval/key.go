package ctxval

type (
	correlationIdKey struct{}
	traceParentKey   struct{}
	traceIDKey       struct{}
	spanIDKey        struct{}
	traceSampledKey  struct{}

	userAgentKey    struct{}
	hostKey         struct{}
	ipKey           struct{}
	forwardedForKey struct{}
	pidKey          struct{}
)
