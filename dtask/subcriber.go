package dtask

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/prometheus/client_golang/prometheus"
)

type (
	SubscriberOption struct {
		queueName string
		Handler   SubscriptionHandler
	}
	SubscriptionHandler func(ctx context.Context, task Task) Response

	Subscriber interface {
		Subscribe(ctx context.Context, queues map[string]int, mux *asynq.ServeMux) (err error)
	}
)

type (
	SubcriberConfig struct {
		Redis           Redis  `json:"redis"`
		Concurrency     int    `json:"concurrency"`
		PrometheusPort  int    `json:"prometheus_port"`
		RetryDelay      string `json:"retry_delay"`
		MetricsRegistry prometheus.Registerer
	}
	Redis struct {
		Address string `json:"address"`
		DB      int    `json:"db"`
	}
)

func (s SubscriberOption) Queue() string {
	return s.queueName
}

func WithSubscriber(queueName string, handler SubscriptionHandler) SubscriberOption {
	return SubscriberOption{
		queueName: queueName,
		Handler:   handler,
	}
}
