package middleware

import (
	"os"
	"regexp"
	"strconv"

	"github.com/google/uuid"
	"github.com/super-saga/go-x/ctxval"
	"github.com/super-saga/go-x/messaging"
	"github.com/super-saga/go-x/messaging/kafka"
)

var (
	traceParentRegex = regexp.MustCompile(`^([a-f0-9]{2})-` + // version
		`([a-f0-9]{32})-` + // traceId
		`([a-f0-9]{16})-` + // parentId
		`([a-f0-9]{2})(-.*)?$`) // flags
	w3cVersion  = "00"
	flagSampled = "01"
)

func Context(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
	return func(message messaging.Message) messaging.Response {
		var (
			pid           = strconv.Itoa(os.Getpid())
			userAgent     = "kafka-consumer-"
			correlationId string
			traceparent   string
			traceID       string
			spanID        string
			traceSampled  bool
		)

		if messageClaim, ok := message.GetMessageClaim().(kafka.MessageClaim); ok {
			for _, hdr := range messageClaim.Headers {
				key := string(hdr.Key)
				switch key {
				case "X-Correlation-Id":
					correlationId = string(hdr.Value)
				case "Traceparent":
					traceparent = string(hdr.Value)
				}
			}
		}

		if len(correlationId) == 0 {
			correlationId = uuid.NewString()
		}
		if len(traceparent) > 0 {
			traceID, spanID, traceSampled = processTraceParent(traceparent)
		}

		ctx := ctxval.Sets(
			message.Context(),
			ctxval.SetCorrelationId(correlationId),
			ctxval.SetTraceParent(traceparent),
			ctxval.SetTraceID(traceID),
			ctxval.SetSpanID(spanID),
			ctxval.SetTraceSampled(traceSampled),

			ctxval.SetUserAgent(userAgent),
			ctxval.SetPid(pid),
		)

		return next(message.WithContext(ctx))
	}
}

// https://www.w3.org/TR/trace-context/
//
// version-format   = trace-id "-" parent-id "-" trace-flags
// trace-id         = 32HEXDIGLC  ; 16 bytes array identifier. All zeroes forbidden
// parent-id        = 16HEXDIGLC  ; 8 bytes array identifier. All zeroes forbidden
// trace-flags      = 2HEXDIGLC   ; 8 bit flags. Currently, only one bit is used. See below for details
//
// Example: 00-0af7651916cd43dd8448eb211c80319c-00f067aa0ba902b7-01
func processTraceParent(traceParent string) (traceID, spanID string, traceSampled bool) {
	subMatches := traceParentRegex.FindStringSubmatch(traceParent)

	if subMatches == nil || len(subMatches) != 6 {
		return
	}
	if !validateVersionAndFlags(subMatches) {
		return
	}

	traceID = subMatches[2]
	spanID = subMatches[3]

	// TODO: better approach for this is to use binary bitwise operator AND
	// static final byte FLAG_SAMPLED = 1; // 00000001
	// boolean sampled = (traceFlags & FLAG_SAMPLED) == FLAG_SAMPLED;
	// will modify it later.
	traceSampled = subMatches[4] == flagSampled

	return
}

func validateVersionAndFlags(subMatches []string) bool {
	if subMatches[1] == w3cVersion {
		if subMatches[5] != "" {
			return false
		}
	}
	// Invalid version: https://w3c.github.io/trace-context/#version
	if subMatches[1] == "ff" {
		return false
	}
	return true
}
