package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/super-saga/go-x/graceful"
)

type pg struct {
	opt PgOption
}

func (dep pg) open(ctx context.Context) (conn *sql.DB, stopper graceful.ProcessStopper, err error) {
	stopper = func(ctx context.Context) error { return nil }
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s&timezone=%s&application_name=%s",
		dep.opt.User, url.QueryEscape(dep.opt.Pass), dep.opt.Host, dep.opt.Port, dep.opt.Name, dep.opt.SslMode, dep.opt.Schema, dep.opt.Tz, dep.opt.ApplicationName,
	)

	prefix := "postgres"
	conn, err = sql.Open(prefix, dsn)
	if err != nil {
		return
	}

	err = conn.Ping()
	if err != nil {
		return
	}

	conn.SetConnMaxIdleTime(dep.opt.ConnMaxIdleTime)
	conn.SetConnMaxLifetime(dep.opt.ConnMaxIdleTime)
	conn.SetMaxIdleConns(dep.opt.MaxIdleConns)
	conn.SetMaxOpenConns(dep.opt.MaxOpenConns)

	if dep.opt.MetricRegistererPrometheus != nil {
		if err = dep.opt.MetricRegistererPrometheus.Register(collectors.NewDBStatsCollector(conn, dep.opt.Name)); err != nil {
			return
		}
	}

	stopper = func(ctx context.Context) error { return conn.Close() }
	return
}

func NewPostgreSQL(ctx context.Context, opt PgOption) (conn *sql.DB, stopper graceful.ProcessStopper, err error) {
	dep := &pg{opt}
	conn, stopper, err = dep.open(ctx)
	return
}
