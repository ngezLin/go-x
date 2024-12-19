package ctxval

import (
	"strconv"

	"google.golang.org/grpc/metadata"
)

func GetMDCorrelationId(md metadata.MD) string {
	v := md.Get(correlationIdMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDTraceParent(md metadata.MD) string {
	v := md.Get(traceParentMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDSpanID(md metadata.MD) string {
	v := md.Get(spanIDMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDTraceID(md metadata.MD) string {
	v := md.Get(traceIDMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDTraceSampled(md metadata.MD) bool {
	v := md.Get(traceSampledMDKey)
	if len(v) == 0 {
		return false
	}

	b, err := strconv.ParseBool(v[0])
	if err != nil {
		return false
	}

	return b
}

func GetMDUserAgent(md metadata.MD) string {
	v := md.Get(userAgentMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDHost(md metadata.MD) string {
	v := md.Get(hostMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDIP(md metadata.MD) string {
	v := md.Get(ipMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDForwardedFor(md metadata.MD) string {
	v := md.Get(forwardedForMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func GetMDPid(md metadata.MD) string {
	v := md.Get(pidMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
