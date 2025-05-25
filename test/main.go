package main

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter interface {
	Allow(token string) bool
}

type rateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	ttl      time.Duration
}

func New(ttl time.Duration) RateLimiter {
	r := &rateLimiter{
		requests: make(map[string][]time.Time),
		ttl:      ttl,
	}

	go r.cleanup()

	return r
}

func (r *rateLimiter) cleanup() {
	r.mu.Lock() // idk I need it or not
	defer r.mu.Unlock()

	ticker := time.NewTicker(r.ttl)
	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			r.requests = make(map[string][]time.Time)
			r.mu.Unlock()
		}
	}
}

func (r *rateLimiter) Allow(token string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	times, exists := r.requests[token]
	if !exists {
		r.requests[token] = []time.Time{time.Now()}
		return true
	}

	if len(times) > 5 {
		return false
	}

	r.requests[token] = append(r.requests[token], time.Now())
	return true
}

func main() {
	rateLimiter := New(time.Minute)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if rateLimiter.Allow(token) {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusTooManyRequests)
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
