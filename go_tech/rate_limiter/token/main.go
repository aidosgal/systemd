package main

import (
	"sync"
	"time"
)

type rateLimiter struct {
	mu          sync.Mutex
	tokens      float64
	capacity    float64
	refillRate  float64
	lastRefill  time.Time
	userBuckets map[string]*UserBucket
}

type UserBucket struct {
	tokens     float64
	lastRefill time.Time
}

type RateLimiter interface {
}

func New(capacity, refillRate float64) RateLimiter {
	return &rateLimiter{
		tokens:      capacity,
		capacity:    capacity,
		refillRate:  refillRate,
		lastRefill:  time.Now(),
		userBuckets: make(map[string]*UserBucket),
	}
}
