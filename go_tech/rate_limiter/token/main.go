package main

import (
	"context"
	"time"
)

type TokenRateLimiter interface {
}

type tokenRateLimiter struct {
	tokenBucketCh chan struct{}
}

func New(ctx context.Context, limit int, period time.Duration) TokenRateLimiter {
	limiter := &tokenRateLimiter{
		tokenBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	replenishmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))
	return limiter
}

func (t *tokenRateLimiter) startPeriodicReplenishment(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			select {
			case t.tokenBucketCh <- struct{}{}:
			default:
			}
		}
	}
}

func (t *tokenRateLimiter) Allow() bool {
	select {
	case <-t.tokenBucketCh:
		return true
	default:
		return false
	}
}
