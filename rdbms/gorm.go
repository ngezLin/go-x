package rdbms

import (
	"context"
	"fmt"
	"net/url"

	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/super-saga/go-x/graceful"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type orm struct {
	opt PgOption
}

func (dep orm) open(ctx context.Context) (client *gorm.DB, stopper graceful.ProcessStopper, err error) {
	stopper = func(ctx context.Context) error { return nil }
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dep.opt.Host, dep.opt.User, url.QueryEscape(dep.opt.Pass), dep.opt.Name, dep.opt.Port, dep.opt.SslMode, dep.opt.Tz)
	client, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		return
	}

	conn, err := client.DB()
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

func NewGorm(ctx context.Context, opt PgOption) (conn *gorm.DB, stopper graceful.ProcessStopper, err error) {
	dep := &orm{opt}
	conn, stopper, err = dep.open(ctx)
	return
}
