package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Cache interface {
	Get() int
}

type cache struct {
	mu          sync.RWMutex
	cachedValue atomic.Int64
}

func NewCache() Cache {
	c := &cache{}
	c.cachedValue.Store(int64(aiPredict()))
	go c.update()
	return c
}

func (c *cache) Get() int {
	value := c.cachedValue.Load()

	return int(value)
}

func (c *cache) update() {
	for _ = range time.Tick(1 * time.Second) {
		pred := aiPredict()
		c.mu.Lock()
		defer c.mu.Unlock()

		c.cachedValue.Store(int64(pred))
	}
}

func aiPredict() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main() {
	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"result\":%d}\n", aiPredict())
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
