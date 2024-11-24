package http_round_tripper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	xlog "github.com/super-saga/go-x/log"
)

type (
	logRoundTripper struct {
		next http.RoundTripper
		opts *logRoundTripperOption
	}
	LogRoundTripper interface {
		http.RoundTripper
	}
)

type (
	logRoundTripperOption struct {
		operation string
		service   string
	}
	logRoundTripperOptions func(*logRoundTripperOption) *logRoundTripperOption
)

func (s *logRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	var (
		ctx                = req.Context()
		timestamp          = time.Now().Local()
		timestampRfc       = timestamp.Format(time.RFC3339)
		body, responseBody string
	)

	if req.Body != nil {
		var bodyBytes []byte
		bodyBytes, _ = io.ReadAll(req.Body)
		// write back to request body
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		//unpretty request json
		dst := &bytes.Buffer{}
		_ = json.Compact(dst, []byte(bodyBytes))
		if dst != nil {
			body = dst.String()
		}
	}

	headers := req.Header.Clone()
	headers.Del("Authorization")
	headers.Del("X-Secret-Key")
	headerStr, err := json.Marshal(headers)
	if err != nil {
		return
	}

	//logging start
	xlog.Info(ctx, fmt.Sprintf("start request %s to %s", s.opts.service, req.URL.String()),
		xlog.String("endpoint", req.URL.String()),
		xlog.String("method", req.Method),
		xlog.String("operation", s.opts.operation),
		xlog.String("service", s.opts.service),
		xlog.String("timestamp", timestampRfc),
		xlog.String("request", body),
		xlog.Any("header", string(headerStr)),
	)

	// Wrap the default RoundTripper with middleware.
	res, err = s.next.RoundTrip(req)
	if err != nil {
		return
	}

	if res.Body != nil {
		var bodyBytes []byte
		bodyBytes, _ = io.ReadAll(res.Body)
		// write back to request body
		res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		//unpretty request json
		dst := &bytes.Buffer{}
		json.Compact(dst, []byte(bodyBytes))
		responseBody = dst.String()
	}

	resHeaderStr, _ := json.Marshal(res.Header)

	//finish start
	xlog.Info(ctx, fmt.Sprintf("finish request %s to %s", s.opts.service, req.URL.String()),
		xlog.String("endpoint", req.URL.String()),
		xlog.String("operation", s.opts.operation),
		xlog.String("service", s.opts.service),
		xlog.String("method", req.Method),
		xlog.String("timestamp", timestampRfc),
		xlog.String("response", string(responseBody)),
		xlog.Any("header", string(resHeaderStr)),
		xlog.Int("status", res.StatusCode),
		xlog.String("latency", fmt.Sprintf("%dms", time.Since(timestamp).Microseconds())),
	)

	return
}

func WithLogRoundTripperOperationOption(operation string) func(*logRoundTripperOption) *logRoundTripperOption {
	return func(s *logRoundTripperOption) *logRoundTripperOption {
		s.operation = operation
		return s
	}
}

func WithLogRoundTripperServiceOption(service string) func(*logRoundTripperOption) *logRoundTripperOption {
	return func(s *logRoundTripperOption) *logRoundTripperOption {
		s.service = service
		return s
	}
}

func NewLogRoundTripper(next http.RoundTripper, opts ...logRoundTripperOptions) *logRoundTripper {
	opt := &logRoundTripperOption{}
	for _, v := range opts {
		opt = v(opt)
	}
	return &logRoundTripper{
		next: next,
		opts: opt,
	}
}
