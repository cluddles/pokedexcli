package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

// Creates new cache.
// Contents will be scanned at the given reapInterval, and any items older than ttls will be removed.
func NewCache(ttl time.Duration, reapInterval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(ttl, reapInterval)
	return &cache
}

// Add value to the cache
func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
}

// Get value from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, exists := c.entries[key]
	if exists {
		// fmt.Printf("(got cached: %s)\n", key)
		return data.data, true
	}
	return []byte{}, false
}

// Checks cache at given interval for items to discard according to ttl
func (c *Cache) reapLoop(ttl time.Duration, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		// fmt.Println("(cache tick)")
		c.reapOnce(ttl)
	}
}

// Discards any cached items that are older than ttl
func (c *Cache) reapOnce(ttl time.Duration) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if now.After(v.createdAt.Add(ttl)) {
			// fmt.Printf("(cache entry expired: %s)\n", k)
			delete(c.entries, k)
		}
	}
}
