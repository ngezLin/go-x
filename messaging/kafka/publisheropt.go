package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
)

type PublisherOption func(*publisherOption)

type publisherOption struct {
	saramaCfg *sarama.Config
	origin    string
	metrics   *publisherPrometheusMetrics
}

var defaultPublisherOptions = []PublisherOption{
	WithPublisherPartitionStrategy("hash"),
}

func newPublisherOption() *publisherOption {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	return &publisherOption{
		saramaCfg: config,
	}
}

func (p *publisherOption) applyOptions(opts ...PublisherOption) {
	for _, o := range opts {
		o(p)
	}
}

// WithOrigin is used to add your service name so we can track the publisher service.
func WithOrigin(origin string) PublisherOption {
	return func(po *publisherOption) {
		po.origin = origin
		po.saramaCfg.ClientID = origin
	}
}

// WithPublisherPartitionStrategy will override the default partition strategy.
// The default is hash.
func WithPublisherPartitionStrategy(strategy string) PublisherOption {
	return func(po *publisherOption) {
		switch strategy {
		case "manual":
			po.saramaCfg.Producer.Partitioner = sarama.NewManualPartitioner
		case "hash":
			po.saramaCfg.Producer.Partitioner = sarama.NewHashPartitioner
		case "random":
			po.saramaCfg.Producer.Partitioner = sarama.NewRandomPartitioner
		case "roundrobin":
			po.saramaCfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
		default:
			po.saramaCfg.Producer.Partitioner = sarama.NewHashPartitioner
		}
	}
}

// WithPublisherIdempotent will set the idempotent while publishing message.
func WithPublisherIdempotent(enable bool, transactionID string) PublisherOption {
	return func(po *publisherOption) {
		po.saramaCfg.Producer.Idempotent = enable
		po.saramaCfg.Producer.Transaction.ID = transactionID
		po.saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
		po.saramaCfg.Net.MaxOpenRequests = 1
	}
}

func WithPublisherKafkaVersion(version string) PublisherOption {
	return func(po *publisherOption) {
		kafkaVersion, err := sarama.ParseKafkaVersion(version)
		if err != nil {
			kafkaVersion = sarama.DefaultVersion
		}
		po.saramaCfg.Version = kafkaVersion
	}
}

func WithPublisherGenericPromMetrics(reg prometheus.Registerer, namespace string, flushInterval time.Duration) PublisherOption {
	return func(po *publisherOption) {
		appMetrics := metrics.NewPrefixedChildRegistry(metrics.NewRegistry(), "sarama.")
		po.saramaCfg.MetricRegistry = appMetrics
		po.metrics = newPublisherPrometheusMetrics(
			namespace, flushInterval,
			reg, appMetrics,
		)
	}
}
