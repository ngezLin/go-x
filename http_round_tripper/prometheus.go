package httproundtripper

import (
	"net/http"
	"time"
)

type (
	prometheusRoundTripper struct {
		next       http.RoundTripper
		collectors *metric.HttpClientCollector
		opt        *Option
	}

	PrometheusRoundTripper interface {
		http.RoundTripper
	}
)

func NewPrometheusCollector(next http.RoundTripper, collectors *metric.HttpClientCollector, option *Option) *prometheusRoundTripper {
	_, _, n := trace.CallerTraceWithPosition(4)
	return &prometheusRoundTripper{
		next:       next,
		collectors: collectors,
		opt: &Option{
			service: n,
		},
	}
}

func (s *prometheusRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	startTime := time.Now()

	s.collectors.RecordGauge(s.opt.service, req.URL.String())

	trace := &promhttp.InstrumentTrace{
		DNSStart: func(t float64) {
			s.collectors.DnsLatencyVec.WithLabelValues("dns_start").Observe(t)
		},
		DNSDone: func(t float64) {
			s.collectors.DnsLatencyVec.WithLabelValues("dns_done").Observe(t)
		},
		TLSHandshakeStart: func(t float64) {
			s.collectors.TlsLatencyVec.WithLabelValues("tls_handshake_start").Observe(t)
		},
		TLSHandshakeDone: func(t float64) {
			s.collectors.TlsLatencyVec.WithLabelValues("tls_handshake_done").Observe(t)
		},
	}

	// Wrap the default RoundTripper with middleware.
	res, err = promhttp.InstrumentRoundTripperTrace(trace, s.next).RoundTrip(req)
	if err != nil {
		return
	}

	s.collectors.Record(time.Since(startTime), s.opt.service, req.Method, req.URL.String(), res.StatusCode)

	return
}
