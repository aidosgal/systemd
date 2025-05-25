package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	value     interface{}
	expiresAt time.Time
	lastUsed  time.Time
}
type cache struct {
	mu       sync.RWMutex
	data     map[string]*CacheItem
	capacity int
	ttl      time.Duration
}

type Cache interface{}

func New(ttl time.Duration, capacity int) Cache {
	cache := &cache{
		data:     make(map[string]*CacheItem),
		capacity: capacity,
		ttl:      ttl,
	}

	go cache.cleanup()
	return cache
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.data[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(value.expiresAt) {
		return nil, false
	}

	c.mu.RUnlock()
	c.mu.Lock()
	value.lastUsed = time.Now()
	c.mu.Unlock()
	c.mu.RLock()

	return value, true
}

func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.data) >= c.capacity {
		if _, exists := c.data[key]; !exists {
			c.evictLRU()
		}
	}

	c.data[key] = &CacheItem{
		value:     value,
		lastUsed:  time.Now(),
		expiresAt: time.Now().Add(c.ttl),
	}
}

func (c *cache) evictLRU() {
	var oldestKey string
	var oldestTime time.Time
	first := true

	for key, item := range c.data {
		if first || item.lastUsed.Before(oldestTime) {
			oldestKey = key
			oldestTime = item.lastUsed
			first = false
		}
	}

	if oldestKey != "" {
		delete(c.data, oldestKey)
	}
}

func (c *cache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, item := range c.data {
				if time.Now().After(item.expiresAt) {
					delete(c.data, key)
				}
			}
		}
	}
}
