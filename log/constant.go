package log

const (
	LogKeyTraceparent  = "traceparent"
	LogKeyTraceID      = "logging.googleapis.com/trace"
	LogKeySpanID       = "logging.googleapis.com/spanId"
	LogKeyTraceSampled = "logging.googleapis.com/trace_sampled"

	LogKeyCorrelationId = "correlation-id"
	LogKeyUserAgent     = "user-agent"
	LogKeyHost          = "host"
	LogKeyIP            = "ip"
	LogKeyForwardedFor  = "x-forwarded-for"
	LogKeyPid           = "pid"
)
