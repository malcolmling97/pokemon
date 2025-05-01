package pokecache

import (
	"sync"
	"time"
)

// Double check what the expireAfter should do just in case
type Cache struct {
	entries     map[string]cacheEntry
	expireAfter time.Duration
	mu          sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	// need to do a reaploop here
	c := &Cache{
		entries:     make(map[string]cacheEntry),
		expireAfter: interval,
	}

	go c.reapLoop()
	return c

}

func (c *Cache) Add(k string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[k] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

}

func (c *Cache) Get(k string) ([]byte, bool) {

	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[k]
	if exists {
		return entry.val, true
	}
	return []byte{}, false

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.expireAfter)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k, entry := range c.entries {
			if time.Since(entry.createdAt) > c.expireAfter {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()

	}

}
