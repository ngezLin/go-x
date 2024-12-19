package ctxval

const (
	headerKeyTraceparent    = "Traceparent"
	headerKeyTrace          = "X-Cloud-Trace-Context"
	headerKeyXForwardedFor  = "X-Forwarded-For"
	headerKeyXCorrelationId = "X-Correlation-Id"
	headerKeyXRealIP        = "X-Real-IP"
)

const (
	correlationIdMDKey = "correlation-id-md-key"
	traceParentMDKey   = "trace-parent-md-key"
	traceIDMDKey       = "trace-id-md-key"
	spanIDMDKey        = "span-id-md-key"
	traceSampledMDKey  = "trace-sampled-md-key"
	userAgentMDKey     = "user-agent-md-key"
	hostMDKey          = "host-md-key"
	ipMDKey            = "ip-md-key"
	forwardedForMDKey  = "forwarded-for-md-key"
	pidMDKey           = "pid-md-key"
)
