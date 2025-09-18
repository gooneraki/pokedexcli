package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {

	new_cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go new_cache.reapLoop()

	return &new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, found := c.entries[key]
	c.mu.Unlock()

	if found {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		c.mu.Lock()
		for k, v := range c.entries {

			if v.createdAt.Add(c.interval).Before(now) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}

}
