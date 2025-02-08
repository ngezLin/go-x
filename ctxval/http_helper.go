package ctxval

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type ClientIP func(req *http.Request) string
type TrustedProxyCheck func(ip net.IP) bool

// SetContextFromHTTP is used to set audit related data to context from http request.
//
// clientIPs:
// clientIPs arg will execute the func in order.
// So it's possible to set preceding, first in first execute.
// If the first func return the value, it would not execute the next func.
// The next func is only executed if the previous return "".
func SetContextFromHTTP(ctx context.Context, req *http.Request, gcpProjectID string, cIPs ...ClientIP) context.Context {
	headerKeyXCorrelationId := req.Header.Get(headerKeyXCorrelationId)
	if headerKeyXCorrelationId == "" {
		headerKeyXCorrelationId = uuid.New().String()
	}

	headerKeyTraceparent := req.Header.Get(headerKeyTraceparent)
	headerKeyTrace := req.Header.Get(headerKeyTrace)
	traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerKeyTrace)
	traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

	userAgent := req.UserAgent()
	host := req.Host
	forwardedFor := req.Header.Get(headerKeyXForwardedFor)
	pid := strconv.Itoa(os.Getpid())

	var ip string
	cIPs = append(cIPs, defaultClientIP())
	for _, cIP := range cIPs {
		if cIP == nil {
			continue
		}
		ip = cIP(req)
		if len(ip) > 0 {
			break
		}
	}

	return Sets(ctx,
		SetCorrelationId(headerKeyXCorrelationId),
		SetTraceParent(headerKeyTraceparent),
		SetTraceID(traceID),
		SetSpanID(spanID),
		SetTraceSampled(traceSampled),

		SetUserAgent(userAgent),
		SetHost(host),
		SetIP(ip),
		SetForwardedFor(forwardedFor),
		SetPid(pid),
	)
}

// SetContextAndMetadataFromHTTP
// A little bit different from SetContextFromHTTP
// In this function we are going to return both context and metadata
// When we are parsing HTTP request to context in http server (grpc-gateway)
// we don't get the value in our grpc server
// So we need to pass the context through metadata
//
// clientIPs:
// clientIPs arg will execute the func in order.
// So it's possible to set preceding, first in first execute.
// If the first func return the value, it would not execute the next func.
// The next func is only executed if the previous return "".
func SetContextAndMetadataFromHTTP(ctx context.Context, req *http.Request, gcpProjectID string, cIPs ...ClientIP) (context.Context, metadata.MD) {
	headerKeyXCorrelationId := req.Header.Get(headerKeyXCorrelationId)
	if headerKeyXCorrelationId == "" {
		headerKeyXCorrelationId = uuid.New().String()
	}

	headerKeyTraceparent := req.Header.Get(headerKeyTraceparent)
	headerKeyTrace := req.Header.Get(headerKeyTrace)
	traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerKeyTrace)
	traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

	userAgent := req.UserAgent()
	host := req.Host
	forwardedFor := req.Header.Get(headerKeyXForwardedFor)
	pid := strconv.Itoa(os.Getpid())

	var ip string
	cIPs = append(cIPs, defaultClientIP())
	for _, cIP := range cIPs {
		if cIP == nil {
			continue
		}
		ip = cIP(req)
		if len(ip) > 0 {
			break
		}
	}

	return Sets(ctx,
			SetCorrelationId(headerKeyXCorrelationId),
			SetTraceParent(headerKeyTraceparent),
			SetTraceID(traceID),
			SetSpanID(spanID),
			SetTraceSampled(traceSampled),

			SetUserAgent(userAgent),
			SetHost(host),
			SetIP(ip),
			SetForwardedFor(forwardedFor),
			SetPid(pid),
		), SetsMD(metadata.New(nil),
			SetMDCorrelationId(headerKeyXCorrelationId),
			SetMDTraceParent(headerKeyTraceparent),
			SetMDTraceID(traceID),
			SetMDSpanID(spanID),
			SetMDTraceSampled(traceSampled),

			SetMDUserAgent(userAgent),
			SetMDHost(host),
			SetMDIP(ip),
			SetMDForwardedFor(forwardedFor),
			SetMDPid(pid),
		)
}

func SetContextFromGRPC(ctx context.Context, gcpProjectID string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	// Handle forwarded http header, but not handling forwarded context from
	// the function above
	if len(md.Get(headerKeyTrace)) > 0 {
		headerKeyTraceparent := md.Get(headerKeyTraceparent)[0]
		headerKeyTrace := md.Get(headerKeyTrace)[0]
		traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerKeyTrace)
		traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

		return Sets(ctx,
			SetTraceParent(headerKeyTraceparent),
			SetTraceID(traceID),
			SetSpanID(spanID),
			SetTraceSampled(traceSampled),
		)
	}

	return Sets(ctx,
		SetCorrelationId(GetMDCorrelationId(md)),
		SetTraceParent(GetMDTraceParent(md)),
		SetTraceID(GetMDTraceID(md)),
		SetSpanID(GetMDSpanID(md)),
		SetTraceSampled(GetMDTraceSampled(md)),

		SetUserAgent(GetMDUserAgent(md)),
		SetHost(GetMDHost(md)),
		SetIP(GetMDIP(md)),
		SetForwardedFor(GetMDForwardedFor(md)),
		SetPid(GetMDPid(md)),
	)
}

func defaultClientIP() ClientIP {
	return WithRemoteAddrClientIP()
}

// Client IP is searched from the req.RemoteAddr.
func WithRemoteAddrClientIP() ClientIP {
	return func(req *http.Request) string {
		return req.RemoteAddr
	}
}

// Client IP is searched from the X-Real-IP header.
func WithXRIClientIP() ClientIP {
	return func(req *http.Request) string {
		return req.Header.Get(headerKeyXRealIP)
	}
}

// Client IP is searched from the rightmost X-Forwarded-For by that count minus one.
func WithXFFTrustedProxyCount(count int) ClientIP {
	return func(req *http.Request) string {
		xForwardedFor := req.Header.Get(headerKeyXForwardedFor)
		if len(xForwardedFor) == 0 {
			return ""
		}
		xForwardedForArr := strings.Split(xForwardedFor, ",")
		pos := len(xForwardedForArr) - count
		if pos > 0 {
			return strings.TrimSpace(xForwardedForArr[pos-1])
		}
		return ""
	}
}

// Client IP is searched from the rightmost X-Forwarded-For, skipping all addresses that are on the trusted proxy list.
func WithXFFTrustedProxyChecker(trusted TrustedProxyCheck) ClientIP {
	return func(req *http.Request) string {
		xForwardedFor := req.Header.Get(headerKeyXForwardedFor)
		if len(xForwardedFor) == 0 {
			return ""
		}
		xForwardedForArr := strings.Split(xForwardedFor, ",")
		for i := len(xForwardedForArr) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(xForwardedForArr[i])
			if !trusted(net.ParseIP(ip)) {
				return ip
			}
		}
		return ""
	}
}

// taken from https://github.com/googleapis/google-cloud-go/blob/master/logging/logging.go#L774
var reCloudTraceContext = regexp.MustCompile(
	// Matches on "TRACE_ID"
	`([a-f\d]+)?` +
		// Matches on "/SPAN_ID"
		`(?:/([a-f\d]+))?` +
		// Matches on ";0=TRACE_TRUE"
		`(?:;o=(\d))?`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string, traceSampled bool) {
	// As per the format described at https://cloud.google.com/trace/docs/setup#force-trace
	//    "X-Cloud-Trace-Context: TRACE_ID/SPAN_ID;o=TRACE_TRUE"
	// for example:
	//    "X-Cloud-Trace-Context: 105445aa7843bc8bf206b120001000/1;o=1"
	//
	// We expect:
	//   * traceID (optional):          "105445aa7843bc8bf206b120001000"
	//   * spanID (optional):           "1"
	//   * traceSampled (optional):     true
	matches := reCloudTraceContext.FindStringSubmatch(s)
	traceID, spanID, traceSampled = matches[1], matches[2], matches[3] == "1"
	if spanID == "0" {
		spanID = ""
	}
	return
}
