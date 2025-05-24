package main

import (
	"container/list"
	"sync"
)

type DB interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
	Delete(key string) bool
}

type entry struct {
	key   string
	value interface{}
}

type Cache struct {
	mu       sync.RWMutex
	data     map[string]*list.Element
	evic     *list.List
	capacity int
}

func New(capacity int) DB {
	return &Cache{
		capacity: capacity,
		data:     make(map[string]*list.Element),
		evic:     list.New(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, found := c.data[key]; found {
		c.evic.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

func (c *Cache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.data[key]; found {
		c.evic.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.evic.Len() >= c.capacity {
		elem := c.evic.Back()
		if elem != nil {
			c.evic.Remove(elem)
			ent := elem.Value.(*entry)
			delete(c.data, ent.key)
		}
	}

	ent := &entry{key: key, value: value}
	elem := c.evic.PushFront(ent)
	c.data[key] = elem
	return
}

func (c *Cache) Delete(key string) bool {
	return false
}
