package metric

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	HttpClientCollector struct {
		reg                    prometheus.Registerer
		ApiRequestDurationHist *prometheus.HistogramVec
		CountRequest           *prometheus.CounterVec
		InFlightGaugeVec       *prometheus.GaugeVec
		DnsLatencyVec          *prometheus.HistogramVec
		TlsLatencyVec          *prometheus.HistogramVec
		HistVec                *prometheus.HistogramVec
	}
)

func NewHttpClientCollector(ctx context.Context, reg prometheus.Registerer) *HttpClientCollector {
	apiRequestDurationHist := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "client_api_request_duration_seconds",
			Help:    "Duration of external API requests in seconds.",
			Buckets: []float64{0, 0.0001, 0.001, 0.010, 0.100, 0.200, 0.500, 1, 2, 5, 10, 100, 1000},
		},
		[]string{"service", "method", "endpoint", "response_code"},
	)

	countRequest := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "client_api_requests_total",
			Help: "A counter for requests from the wrapped client.",
		},
		[]string{"service", "method", "endpoint", "response_code"},
	)

	inFlightGaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "client_in_flight_requests",
		Help: "A gauge of in-flight requests for the wrapped client.",
	},
		[]string{"service", "endpoint"},
	)

	// dnsLatencyVec uses custom buckets based on expected dns durations.
	// It has an instance label "event", which is set in the
	// DNSStart and DNSDonehook functions defined in the
	// InstrumentTrace struct below.
	dnsLatencyVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "dns_duration_seconds",
			Help:    "Trace dns latency histogram.",
			Buckets: []float64{.005, .01, .025, .05},
		},
		[]string{"event"},
	)

	// tlsLatencyVec uses custom buckets based on expected tls durations.
	// It has an instance label "event", which is set in the
	// TLSHandshakeStart and TLSHandshakeDone hook functions defined in the
	// InstrumentTrace struct below.
	tlsLatencyVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tls_duration_seconds",
			Help:    "Trace tls latency histogram.",
			Buckets: []float64{.05, .1, .25, .5},
		},
		[]string{"event"},
	)

	// histVec has no labels, making it a zero-dimensional ObserverVec.
	histVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "A histogram of request latencies.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{},
	)

	for _, c := range []prometheus.Collector{countRequest, apiRequestDurationHist, tlsLatencyVec, dnsLatencyVec, histVec, inFlightGaugeVec} {
		reg.MustRegister(c)
	}

	return &HttpClientCollector{
		reg:                    reg,
		ApiRequestDurationHist: apiRequestDurationHist,
		CountRequest:           countRequest,
		InFlightGaugeVec:       inFlightGaugeVec,
		DnsLatencyVec:          dnsLatencyVec,
		TlsLatencyVec:          tlsLatencyVec,
		HistVec:                histVec,
	}
}

func (m *HttpClientCollector) Record(duration time.Duration, service, method, endpoint string, statusCode int) {
	m.ApiRequestDurationHist.WithLabelValues(service, method, endpoint, fmt.Sprint(statusCode)).
		Observe(duration.Seconds())
	m.CountRequest.WithLabelValues(service, method, endpoint, fmt.Sprint(statusCode)).Inc()
	m.InFlightGaugeVec.WithLabelValues(service, endpoint).Dec()
}

func (m *HttpClientCollector) RecordGauge(service, endpoint string) {
	m.InFlightGaugeVec.WithLabelValues(service, endpoint).Inc()
}
