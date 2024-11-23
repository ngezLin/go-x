package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/super-saga/go-x/graceful"
)

type (
	Redis interface {
		Set(ctx context.Context, key string, value interface{}, expirationTime time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		SetIfNotExist(ctx context.Context, key string, value interface{}, expirationTime time.Duration) (bool, error)
		DeleteIfExist(ctx context.Context, lockKey string) error
	}
	cache struct {
		Client *redis.Client
	}
)

func New(c *redis.Client) (Redis, graceful.ProcessStopper) {
	stopper := func(context.Context) error { return nil }
	rc := &cache{
		Client: c,
	}
	stopper = func(ctx context.Context) error {
		return rc.Client.Close()
	}
	return rc, stopper
}

func (rc *cache) Get(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return rc.Client.Set(ctx, key, value, expiration).Err()
}

func (rc *cache) SetIfNotExist(ctx context.Context, key string, value interface{}, expirationTime time.Duration) (bool, error) {
	return rc.Client.SetNX(ctx, key, value, expirationTime).Result()
}

func (rc *cache) DeleteIfExist(ctx context.Context, lockKey string) error {
	//idemp key is set - remove distributed key
	val := rc.Client.Get(ctx, lockKey).Val()
	if len(val) > 0 {
		// del locker key
		return rc.Client.Del(ctx, lockKey).Err()
	}
	return nil
}
