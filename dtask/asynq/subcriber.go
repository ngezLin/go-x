package asynq

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynq/x/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/super-saga/go-x/dtask"
	"github.com/super-saga/go-x/graceful"
)

type (
	subcriber struct {
		redisOption asynq.RedisClientOpt
		Server      *asynq.Server
		stopper     graceful.ProcessStopper

		cfg asynq.Config

		subscriberConfig dtask.SubcriberConfig
	}
)

func NewSubscriber(ctx context.Context, subscriberConfig dtask.SubcriberConfig) (*subcriber, graceful.ProcessStopper, error) {
	p := &subcriber{
		redisOption: asynq.RedisClientOpt{
			Addr: subscriberConfig.Redis.Address,
			DB:   subscriberConfig.Redis.DB,
		},
		cfg: asynq.Config{
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
				if subscriberConfig.RetryDelay != "" {
					dur, err := time.ParseDuration(subscriberConfig.RetryDelay)
					if err != nil {
						return asynq.DefaultRetryDelayFunc(n, e, t)
					}
					return dur
				}
				return asynq.DefaultRetryDelayFunc(n, e, t)
			},
			Concurrency: subscriberConfig.Concurrency,
			Queues:      map[string]int{},
			BaseContext: func() context.Context {
				return ctx
			},
		},
		subscriberConfig: subscriberConfig,
	}

	stopper := func(ctx context.Context) error {
		if p.Server != nil {
			p.Server.Shutdown()
		}
		return nil
	}

	p.stopper = stopper

	return p, p.stopper, nil
}

func (s *subcriber) Subscribe(ctx context.Context, queues map[string]int, mux *asynq.ServeMux) (err error) {
	s.cfg.Queues = queues
	s.Server = asynq.NewServer(
		s.redisOption,
		s.cfg,
	)
	if s.subscriberConfig.PrometheusPort != 0 {
		// monitoring
		mx := http.NewServeMux()
		inspector := asynq.NewInspector(s.redisOption)

		if s.subscriberConfig.MetricsRegistry != nil {
			s.subscriberConfig.MetricsRegistry.MustRegister(
				metrics.NewQueueMetricsCollector(inspector),
			)
			mx.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
		}

		mx.Handle("/health", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Default().Println(r.Context(), "ok")
			},
		))
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", s.subscriberConfig.PrometheusPort),
			Handler: mx,
		}

		//running monitor
		go func() {
			log.Default().Println("server mux prometheus x asynq run")
			err := srv.ListenAndServe()
			if err != nil {
				log.Default().Println("server mux prometheus x asynq closed: %s", err.Error())
			}
		}()
	}

	err = s.Server.Run(mux)
	if err != nil {
		return err
	}

	return
}
