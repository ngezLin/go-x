package metrics

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	HttpClientCollector struct {
		reg                    prometheus.Registerer
		ApiRequestDurationHist *prometheus.HistogramVec
		CountRequest           *prometheus.CounterVec
		InFlightGaugeVec       *prometheus.GaugeVec
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

	// histVec has no labels, making it a zero-dimensional ObserverVec.
	histVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "A histogram of request latencies.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{},
	)

	for _, c := range []prometheus.Collector{countRequest, apiRequestDurationHist, histVec, inFlightGaugeVec} {
		reg.MustRegister(c)
	}

	return &HttpClientCollector{
		reg:                    reg,
		ApiRequestDurationHist: apiRequestDurationHist,
		CountRequest:           countRequest,
		InFlightGaugeVec:       inFlightGaugeVec,
		HistVec:                histVec,
	}
}

func (m *HttpClientCollector) Record(duration time.Duration, service, method, endpoint string, statusCode int) {
	URL, _ := url.Parse(endpoint)
	if URL != nil {
		endpoint = fmt.Sprintf("%s://%s%s", URL.Scheme, URL.Host, URL.Path)
	}
	if m.ApiRequestDurationHist != nil {
		m.ApiRequestDurationHist.WithLabelValues(service, method, endpoint, fmt.Sprint(statusCode)).
			Observe(duration.Seconds())
	}
	if m.CountRequest != nil {
		m.CountRequest.WithLabelValues(service, method, endpoint, fmt.Sprint(statusCode)).Inc()
	}
	if m.InFlightGaugeVec != nil {
		m.InFlightGaugeVec.WithLabelValues(service, endpoint).Dec()
	}
}

func (m *HttpClientCollector) RecordGauge(service, endpoint string) {
	if m.InFlightGaugeVec != nil {
		URL, _ := url.Parse(endpoint)
		if URL != nil {
			endpoint = fmt.Sprintf("%s://%s%s", URL.Scheme, URL.Host, URL.Path)
		}
		m.InFlightGaugeVec.WithLabelValues(service, endpoint).Inc()
	}
}
