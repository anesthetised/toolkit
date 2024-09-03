package ds

import "sync"

// Cache implements a simple generic thread-safe cache.
type Cache[K comparable, V any] struct {
	mu    sync.RWMutex
	state map[K]V
}

// NewCache returns a new Cache with specified type parameters.
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		state: make(map[K]V),
	}
}

// Get retrieves a value from the cache.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	v, ok := c.state[key]
	c.mu.RUnlock()
	return v, ok
}

// Set places a value into the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	c.state[key] = value
	c.mu.Unlock()
}

// Delete removes a value from the cache.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	delete(c.state, key)
	c.mu.Unlock()
}
