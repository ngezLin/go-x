package ratelimiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

type (
	RateLimit struct {
		Allowed    bool
		RetryAfter time.Duration
	}
	RateLimiter interface {
		GetRateLimit(ctx context.Context, key string, limit, value int, lu time.Duration) (rl RateLimit, err error)
	}
	rateLimiter struct {
		limiter *redis_rate.Limiter
	}
)

func New(rc *redis.Client) RateLimiter {
	limiter := &rateLimiter{
		limiter: redis_rate.NewLimiter(rc),
	}
	return limiter
}

func (r rateLimiter) GetRateLimit(ctx context.Context, key string, limit, period int, lu time.Duration) (rl RateLimit, err error) {
	var req redis_rate.Limit
	switch lu {
	case time.Second:
		req = redis_rate.PerHour(limit)
	case time.Minute:
		req = redis_rate.PerMinute(limit)
	default:
		req = redis_rate.PerSecond(limit)
	}
	req.Period = req.Period * time.Duration(period)
	res, err := r.limiter.Allow(ctx, key, req)
	if err != nil {
		return
	}
	rl.Allowed = res.Remaining > 0 || (res.Remaining == 0 && res.Allowed > 0)
	rl.RetryAfter = res.RetryAfter
	return
}
