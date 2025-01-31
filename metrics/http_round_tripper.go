package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/super-saga/go-x/trace"
)

type (
	promHttpClientRoundTripper struct {
		next       http.RoundTripper
		collectors *HttpClientCollector
		opts       *promHttpClientRoundTripperOptions
	}
	promHttpClientRoundTripperOptions struct {
		service string
	}
	promHttpClientRoundTripperOption func(*promHttpClientRoundTripperOptions)

	PromHttpClientRoundTripper interface {
		http.RoundTripper
	}
)

func NewPromHttpClientRoundTripper(next http.RoundTripper, collectors *HttpClientCollector, options ...promHttpClientRoundTripperOption) *promHttpClientRoundTripper {
	opts := NewOpts(options...)
	if opts.service == "" {
		_, _, n := trace.CallerTrace()
		opts.service = n
	}
	return &promHttpClientRoundTripper{
		next:       next,
		collectors: collectors,
		opts:       opts,
	}
}

func NewOpts(opts ...promHttpClientRoundTripperOption) *promHttpClientRoundTripperOptions {
	_, _, n := trace.CallerTrace()
	opt := &promHttpClientRoundTripperOptions{
		service: n,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func WithOptionService(service string) promHttpClientRoundTripperOption {
	return func(o *promHttpClientRoundTripperOptions) {
		o.service = service
	}
}

func (s *promHttpClientRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	startTime := time.Now()
	endpoint := fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path)

	s.collectors.RecordGauge(s.opts.service, endpoint)

	// Wrap the default RoundTripper with middleware.
	res, err = s.next.RoundTrip(req)
	if err != nil {
		return
	}

	s.collectors.Record(time.Since(startTime), s.opts.service, req.Method, endpoint, res.StatusCode)

	return
}
