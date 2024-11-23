package kafka

import (
	"time"

	prometheusmetrics "github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
)

type publisherPrometheusMetrics struct {
	namespace       string
	flushInterval   time.Duration
	registerer      prometheus.Registerer
	metrics         metrics.Registry
	publishTimeHist *prometheus.HistogramVec
}

func newPublisherPrometheusMetrics(
	namespace string, flushInterval time.Duration,
	reg prometheus.Registerer, metrics metrics.Registry) *publisherPrometheusMetrics {

	publishTimeHist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "kafka_publisher_publish_time",
		Help:    "publish time of kafka publisher handler",
		Buckets: []float64{0, 0.0001, 0.001, 0.010, 0.100, 0.200, 0.500, 1, 2, 5, 10, 100, 1000},
	}, []string{"topic"})

	reg.MustRegister(publishTimeHist)

	return &publisherPrometheusMetrics{
		namespace:       namespace,
		flushInterval:   flushInterval,
		registerer:      reg,
		metrics:         metrics,
		publishTimeHist: publishTimeHist,
	}
}

func (p *publisherPrometheusMetrics) Run() {
	prometheusClient := prometheusmetrics.NewPrometheusProvider(
		p.metrics, p.namespace, "", p.registerer, p.flushInterval,
	)
	go prometheusClient.UpdatePrometheusMetrics()
}

func (p *publisherPrometheusMetrics) Record(topic string, diff time.Duration) {
	p.publishTimeHist.WithLabelValues(topic).Observe(diff.Seconds())
}
