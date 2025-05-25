package main

import (
	"sync"
	"time"
)

type RateLimiter interface {
	Allow(userID string) bool
}

type rateLimiter struct {
	mu   sync.RWMutex
	ttl  time.Duration
	data map[string][]time.Time
}

func New(ttl time.Duration) RateLimiter {
	rateLimiter := &rateLimiter{
		ttl:  ttl,
		data: make(map[string][]time.Time),
	}
	go rateLimiter.cleanUp()

	return rateLimiter
}

func (r *rateLimiter) cleanUp() {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticker := time.NewTicker(r.ttl)
	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			r.data = make(map[string][]time.Time)
			r.mu.Unlock()
		}
	}
}

func (r *rateLimiter) Allow(userID string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	times, exists := r.data[userID]
	if !exists {
		return true
	}
	for _, t := range times {
		if time.Now().Add(r.ttl).After(t) {
			return false
		}
	}
	return true
}
