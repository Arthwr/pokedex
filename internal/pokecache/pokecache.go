package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux   sync.Mutex
	store map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mux:   sync.Mutex{},
		store: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.store[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for k, entry := range c.store {
			if time.Since(entry.createdAt) > interval {
				delete(c.store, k)
			}
		}
		c.mux.Unlock()
	}
}
