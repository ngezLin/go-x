package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	PapaStreamCollector struct {
		reg          Registry
		CountRequest *prometheus.CounterVec
	}
)

func NewPapaStreamCollector(ctx context.Context, reg Registry) *PapaStreamCollector {
	countRequest := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "papa_stream_requests_total",
			Help: "A counter for papa stream from the deliveries.",
		},
		[]string{"payment_method", "status"},
	)
	for _, c := range []prometheus.Collector{countRequest} {
		reg.MustRegister(c)
	}
	return &PapaStreamCollector{
		reg:          reg,
		CountRequest: countRequest,
	}
}
func (m *PapaStreamCollector) Record(paymentMethod, status string) {
	if m.CountRequest != nil {
		m.CountRequest.WithLabelValues(paymentMethod, status).Inc()
	}
}
