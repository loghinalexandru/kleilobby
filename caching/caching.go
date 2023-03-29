package caching

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	data map[string]T
	time map[string]time.Time
	mtx  sync.RWMutex
}

func New[T any]() *Cache[T] {
	return &Cache[T]{
		data: make(map[string]T),
		time: make(map[string]time.Time),
		mtx:  sync.RWMutex{},
	}
}

func (c *Cache[T]) Get(key string) T {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.data[key]
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
	c.time[key] = time.Now().UTC()

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

	_, ok := c.data[key]

	return ok
}
