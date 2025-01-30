package rdbms

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	PgOption struct {
		Host    string
		User    string
		Pass    string
		Port    string
		Name    string
		Schema  string
		SslMode string
		Tz      string
		Dialect string

		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxIdleTime time.Duration
		ConnMaxLifeTime time.Duration

		ApplicationName            string
		MetricRegistererPrometheus prometheus.Registerer
	}
)
