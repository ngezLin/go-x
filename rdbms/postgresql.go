package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
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

		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxIdleTime time.Duration
		ConnMaxLifeTime time.Duration

		ApplicationName            string
		MetricRegistererPrometheus prometheus.Registerer
	}
)

func NewPostgreSQL(ctx context.Context, option *PgOption) (conn *sql.DB, stopper func(ctx context.Context) error, err error) {
	stopper = func(ctx context.Context) error { return nil }
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s&timezone=%s&application_name=%s",
		option.User, url.QueryEscape(option.Pass), option.Host, option.Port, option.Name, option.SslMode, option.Schema, option.Tz, option.ApplicationName,
	)

	prefix := "pgx"
	conn, err = sql.Open(prefix, dsn)
	if err != nil {
		return
	}

	err = conn.Ping()
	if err != nil {
		return
	}

	conn.SetConnMaxIdleTime(option.ConnMaxIdleTime)
	conn.SetConnMaxLifetime(option.ConnMaxIdleTime)
	conn.SetMaxIdleConns(option.MaxIdleConns)
	conn.SetMaxOpenConns(option.MaxOpenConns)

	if option.MetricRegistererPrometheus != nil {
		if err = option.MetricRegistererPrometheus.Register(collectors.NewDBStatsCollector(conn, option.Name)); err != nil {
			return
		}
	}

	stopper = func(ctx context.Context) error { return conn.Close() }

	return
}
