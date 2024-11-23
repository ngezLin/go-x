package kafka

import (
	"strconv"
	"time"

	prometheusmetrics "github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
	"github.com/super-saga/go-x/messaging"
)

type subscriberPrometheusMetrics struct {
	namespace          string
	subsystem          string
	flushInterval      time.Duration
	registerer         prometheus.Registerer
	metrics            metrics.Registry
	consumeTimeHist    *prometheus.HistogramVec
	processingTimeHist *prometheus.HistogramVec
	getMessageTimeHist *prometheus.HistogramVec
}

func newSubscriberPrometheusMetrics(
	namespace, subsystem string, flushInterval time.Duration,
	reg prometheus.Registerer, metrics metrics.Registry) *subscriberPrometheusMetrics {

	consumeTimeHist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "kafka_consumer_consume_time",
		Help:    "consume time of kafka consumer handler",
		Buckets: []float64{0, 0.0001, 0.001, 0.010, 0.100, 0.200, 0.500, 1, 2, 5, 10, 100, 1000},
	}, []string{"topic", "consumer_group"})

	processingTimeHist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "kafka_consumer_processing_time",
		Help:    "processing time of kafka consumer handler",
		Buckets: []float64{0, 0.0001, 0.001, 0.010, 0.100, 0.200, 0.500, 1, 2, 5, 10, 100, 1000},
	}, []string{"topic", "success", "consumer_group"})

	getMessageTimeHist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "kafka_consumer_get_message_time",
		Help:    "get message time of kafka consumer handler",
		Buckets: []float64{0, 0.0001, 0.001, 0.010, 0.100, 0.200, 0.500, 1, 2, 5, 10, 100, 1000},
	}, []string{"topic", "consumer_group"})

	reg.MustRegister(consumeTimeHist, processingTimeHist, getMessageTimeHist)

	return &subscriberPrometheusMetrics{
		namespace:          namespace,
		subsystem:          subsystem,
		flushInterval:      flushInterval,
		registerer:         reg,
		metrics:            metrics,
		consumeTimeHist:    consumeTimeHist,
		processingTimeHist: processingTimeHist,
		getMessageTimeHist: getMessageTimeHist,
	}
}

func (p *subscriberPrometheusMetrics) Run() {
	prometheusClient := prometheusmetrics.NewPrometheusProvider(
		p.metrics, p.namespace, p.subsystem, p.registerer, p.flushInterval,
	)
	go prometheusClient.UpdatePrometheusMetrics()
}

func (p *subscriberPrometheusMetrics) Middleware(consumerGroup string) messaging.MiddlewareFunc {
	return func(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
		return func(message messaging.Message) messaging.Response {
			startTime := time.Now() // time when a process consumes a message started
			r := next(message)
			endTime := time.Now() // time when a process consumes a message finished

			v := message.GetMessageClaim().(MessageClaim)
			if v != nil {
				p.consumeTimeHist.WithLabelValues(v.Topic, consumerGroup).
					Observe(endTime.Sub(v.Timestamp).Seconds())

				p.processingTimeHist.WithLabelValues(v.Topic, strconv.FormatBool(!r.IsError()), consumerGroup).
					Observe(endTime.Sub(startTime).Seconds())

				p.getMessageTimeHist.WithLabelValues(v.Topic, consumerGroup).
					Observe(startTime.Sub(v.Timestamp).Seconds())
			}

			return r
		}
	}
}
