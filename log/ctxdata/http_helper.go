package ctxdata

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
	headerXCorrelationId := req.Header.Get(headerXCorrelationId)
	if headerXCorrelationId == "" {
		headerXCorrelationId = uuid.New().String()
	}

	headerTraceparent := req.Header.Get(headerTraceparent)
	headerTrace := req.Header.Get(headerTrace)
	traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerTrace)
	traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

	userAgent := req.UserAgent()
	host := req.Host
	forwardedFor := req.Header.Get(headerXForwardedFor)
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
		SetCorrelationId(headerXCorrelationId),
		SetTraceParent(headerTraceparent),
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
	headerXCorrelationId := req.Header.Get(headerXCorrelationId)
	if headerXCorrelationId == "" {
		headerXCorrelationId = uuid.New().String()
	}

	headerTraceparent := req.Header.Get(headerTraceparent)
	headerTrace := req.Header.Get(headerTrace)
	traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerTrace)
	traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

	userAgent := req.UserAgent()
	host := req.Host
	forwardedFor := req.Header.Get(headerXForwardedFor)
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
			SetCorrelationId(headerXCorrelationId),
			SetTraceParent(headerTraceparent),
			SetTraceID(traceID),
			SetSpanID(spanID),
			SetTraceSampled(traceSampled),

			SetUserAgent(userAgent),
			SetHost(host),
			SetIP(ip),
			SetForwardedFor(forwardedFor),
			SetPid(pid),
		), SetsMD(metadata.New(nil),
			SetMDCorrelationId(headerXCorrelationId),
			SetMDTraceParent(headerTraceparent),
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
	if len(md.Get(headerTrace)) > 0 {
		headerTraceparent := md.Get(headerTraceparent)[0]
		headerTrace := md.Get(headerTrace)[0]
		traceID, spanID, traceSampled := deconstructXCloudTraceContext(headerTrace)
		traceID = fmt.Sprintf("projects/%s/traces/%s", gcpProjectID, traceID)

		return Sets(ctx,
			SetTraceParent(headerTraceparent),
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
		return req.Header.Get(headerXRealIP)
	}
}

// Client IP is searched from the rightmost X-Forwarded-For by that count minus one.
func WithXFFTrustedProxyCount(count int) ClientIP {
	return func(req *http.Request) string {
		xForwardedFor := req.Header.Get(headerXForwardedFor)
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
		xForwardedFor := req.Header.Get(headerXForwardedFor)
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

func SetCorrelationId(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, correlationIdKey{}, v)
	}
}

func SetMDCorrelationId(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[correlationIdMDKey] = []string{v}
		return md
	}
}

func GetCorrelationId(ctx context.Context) string {
	if v, ok := ctx.Value(correlationIdKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDCorrelationId(md metadata.MD) string {
	v := md.Get(correlationIdMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetTraceParent(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceParentKey{}, v)
	}
}

func SetMDTraceParent(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceParentMDKey] = []string{v}
		return md
	}
}

func GetTraceParent(ctx context.Context) string {
	if v, ok := ctx.Value(traceParentKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDTraceParent(md metadata.MD) string {
	v := md.Get(traceParentMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetSpanID(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, spanIDKey{}, v)
	}
}

func SetMDSpanID(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[spanIDMDKey] = []string{v}
		return md
	}
}

func GetSpanID(ctx context.Context) string {
	if v, ok := ctx.Value(spanIDKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDSpanID(md metadata.MD) string {
	v := md.Get(spanIDMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetTraceID(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceIDKey{}, v)
	}
}

func SetMDTraceID(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceIDMDKey] = []string{v}
		return md
	}
}

func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.Value(traceIDKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDTraceID(md metadata.MD) string {
	v := md.Get(traceIDMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetTraceSampled(v bool) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceSampledKey{}, v)
	}
}

func SetMDTraceSampled(v bool) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[traceSampledMDKey] = []string{fmt.Sprintf("%v", v)}
		return md
	}
}

func GetTraceSampled(ctx context.Context) bool {
	if v, ok := ctx.Value(traceSampledKey{}).(bool); ok {
		return v
	}
	return false
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

func SetUserAgent(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userAgentKey{}, v)
	}
}

func SetMDUserAgent(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[userAgentMDKey] = []string{v}
		return md
	}
}

func GetUserAgent(ctx context.Context) string {
	if v, ok := ctx.Value(userAgentKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDUserAgent(md metadata.MD) string {
	v := md.Get(userAgentMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetHost(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, hostKey{}, v)
	}
}

func SetMDHost(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[hostMDKey] = []string{v}
		return md
	}
}

func GetHost(ctx context.Context) string {
	if v, ok := ctx.Value(hostKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDHost(md metadata.MD) string {
	v := md.Get(hostMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetIP(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ipKey{}, v)
	}
}

func SetMDIP(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[ipMDKey] = []string{v}
		return md
	}
}

func GetIP(ctx context.Context) string {
	if v, ok := ctx.Value(ipKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDIP(md metadata.MD) string {
	v := md.Get(ipMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetForwardedFor(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, forwardedForKey{}, v)
	}
}

func SetMDForwardedFor(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[forwardedForMDKey] = []string{v}
		return md
	}
}

func GetForwardedFor(ctx context.Context) string {
	if v, ok := ctx.Value(forwardedForKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDForwardedFor(md metadata.MD) string {
	v := md.Get(forwardedForMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func SetPid(v string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pidKey{}, v)
	}
}

func SetMDPid(v string) SetMD {
	return func(md metadata.MD) metadata.MD {
		md[pidMDKey] = []string{v}
		return md
	}
}

func GetPid(ctx context.Context) string {
	if v, ok := ctx.Value(pidKey{}).(string); ok {
		return v
	}
	return ""
}

func GetMDPid(md metadata.MD) string {
	v := md.Get(pidMDKey)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
