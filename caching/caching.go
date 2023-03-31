package caching

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	data    map[string]T
	time    map[string]time.Time
	expTime time.Duration
	mtx     sync.RWMutex
}

func New[T any](expTime time.Duration) *Cache[T] {
	return &Cache[T]{
		data:    make(map[string]T),
		time:    make(map[string]time.Time),
		expTime: expTime,
		mtx:     sync.RWMutex{},
	}
}

func (c *Cache[T]) Get(key string) T {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	var result T
	timestamp, ok := c.time[key]

	if ok && time.Now().UTC().Before(timestamp) {
		return c.data[key]
	}

	return result
}

func (c *Cache[T]) GetTimestamp(key string) time.Time {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.time[key]
}

func (c *Cache[T]) Add(key string, object T) T {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.data[key] = object
	c.time[key] = time.Now().UTC().Add(c.expTime)

	return object
}

func (c *Cache[T]) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.data = make(map[string]T)
}

func (c *Cache[T]) Contains(key string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	timestamp, ok := c.time[key]

	if ok && time.Now().UTC().Before(timestamp) {
		return true
	}

	return false
}
