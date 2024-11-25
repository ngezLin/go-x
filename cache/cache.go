package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expirationTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	SetIfNotExist(ctx context.Context, key string, value interface{}, expirationTime time.Duration) (bool, error)
	DeleteIfExist(ctx context.Context, lockKey string) error
}
