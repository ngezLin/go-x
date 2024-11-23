package kafka

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
	"github.com/super-saga/go-x/messaging"
)

type SubscriberOption func(*subscriberOption)

type Offset int64

const (
	OffsetNewest Offset = iota
	OffsetOldest
)

type BalanceStrategy int64

const (
	BalanceStrategySticky BalanceStrategy = iota
	BalanceStrategyRoundRobin
	BalanceStrategyRange
)

type subscriberOption struct {
	saramaCfg   *sarama.Config
	saramaLog   sarama.StdLogger
	middlewares []messaging.MiddlewareFunc
	errCb       func(error)
	preStartCb  func(context.Context, map[string][]int32) error
	endCb       func(context.Context, map[string][]int32) error
	metrics     *subscriberPrometheusMetrics
}

var defaultSubscriberOptions = []SubscriberOption{
	WithSubscriberInitialOffset(OffsetOldest),
	WithSubscriberBalanceStrategy(BalanceStrategySticky),
	WithSubsciberPreStartCallback(func(context.Context, map[string][]int32) error { return nil }),
	WithSubsciberEndCallback(func(context.Context, map[string][]int32) error { return nil }),
}

func newSubscriberOption() *subscriberOption {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Consumer.Fetch.Min = 1
	saramaCfg.Consumer.Fetch.Default = 1
	saramaCfg.ClientID, _ = os.Hostname()

	sarama.Logger = log.New(os.Stdout, "[KAFKA-CONSUMER] ", log.LstdFlags)
	return &subscriberOption{
		saramaCfg: saramaCfg,
		saramaLog: sarama.Logger,
	}
}

func (p *subscriberOption) applyOptions(opts ...SubscriberOption) {
	for _, o := range opts {
		o(p)
	}
	if p.errCb != nil {
		p.saramaCfg.Consumer.Return.Errors = true
	}
}

func WithSubscriberKafkaVersion(version string) SubscriberOption {
	return func(o *subscriberOption) {
		kafkaVersion, err := sarama.ParseKafkaVersion(version)
		if err != nil {
			kafkaVersion = sarama.DefaultVersion
		}
		o.saramaCfg.Version = kafkaVersion
	}
}

func WithSubscriberInitialOffset(offset Offset) SubscriberOption {
	return func(o *subscriberOption) {
		switch offset {
		case OffsetOldest:
			o.saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
		case OffsetNewest:
			o.saramaCfg.Consumer.Offsets.Initial = sarama.OffsetNewest
		default:
			o.saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
		}
	}
}

func WithSubscriberBalanceStrategy(strategy BalanceStrategy) SubscriberOption {
	return func(o *subscriberOption) {
		switch strategy {
		case BalanceStrategySticky:
			o.saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
		case BalanceStrategyRange:
			o.saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
		case BalanceStrategyRoundRobin:
			o.saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
		default:
			o.saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
		}
	}
}

func WithMiddleware(middlewares ...messaging.MiddlewareFunc) SubscriberOption {
	return func(o *subscriberOption) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func WithSubscriberErrorCallback(errCb func(error)) SubscriberOption {
	return func(o *subscriberOption) {
		o.errCb = errCb
	}
}

func WithSubsciberPreStartCallback(cb func(context.Context, map[string][]int32) error) SubscriberOption {
	return func(o *subscriberOption) {
		o.preStartCb = cb
	}
}

func WithSubsciberEndCallback(cb func(context.Context, map[string][]int32) error) SubscriberOption {
	return func(o *subscriberOption) {
		o.endCb = cb
	}
}

func WithSubscriberGenericPromMetrics(reg prometheus.Registerer, namespace, subsystem string, flushInterval time.Duration) SubscriberOption {
	return func(so *subscriberOption) {
		appMetrics := metrics.NewPrefixedChildRegistry(metrics.NewRegistry(), "sarama.")
		so.saramaCfg.MetricRegistry = appMetrics
		so.metrics = newSubscriberPrometheusMetrics(
			namespace, subsystem, flushInterval,
			reg, appMetrics,
		)
	}
}
