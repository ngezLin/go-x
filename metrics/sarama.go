package metrics

import (
	"context"
	"time"

	prometheusmetrics "github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	saramaMetrics "github.com/rcrowley/go-metrics"
)

func SaramaRegistry(ctx context.Context, reg prometheus.Registerer, name string, flushInterval time.Duration) saramaMetrics.Registry {
	appMetrics := saramaMetrics.NewPrefixedRegistry(name + "_")
	prometheusClient := prometheusmetrics.NewPrometheusProvider(
		appMetrics, "", "", reg, flushInterval,
	)
	go prometheusClient.UpdatePrometheusMetrics()

	return appMetrics
}
