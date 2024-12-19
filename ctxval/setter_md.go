package ctxval

import (
	"fmt"

	"google.golang.org/grpc/metadata"
)

func SetMDCorrelationId(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[correlationIdMDKey] = []string{v}
		return md
	}
}

func SetMDTraceParent(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceParentMDKey] = []string{v}
		return md
	}
}

func SetMDSpanID(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[spanIDMDKey] = []string{v}
		return md
	}
}

func SetMDTraceID(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceIDMDKey] = []string{v}
		return md
	}
}

func SetMDTraceSampled(v bool) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceSampledMDKey] = []string{fmt.Sprintf("%v", v)}
		return md
	}
}

func SetMDUserAgent(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[userAgentMDKey] = []string{v}
		return md
	}
}

func SetMDHost(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[hostMDKey] = []string{v}
		return md
	}
}

func SetMDIP(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[ipMDKey] = []string{v}
		return md
	}
}

func SetMDForwardedFor(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[forwardedForMDKey] = []string{v}
		return md
	}
}

func SetMDPid(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[pidMDKey] = []string{v}
		return md
	}
}
