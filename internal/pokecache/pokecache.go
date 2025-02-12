package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex    // Is initialized here
	quit  chan struct{} // Channel to signal shutdown
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache { // Return a pointer to Cache
	c := &Cache{ // Create and return a point to Cache
		cache: make(map[string]cacheEntry),
		mu:    sync.Mutex{}, // Not required, just for visual clarity
		quit:  make(chan struct{}),
	}
	go c.reapLoop(interval) // On another goroutine to avoid blocking main thread
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Ticker that ticks at specified interval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: // Wait for the next tick
			c.mu.Lock()       // Locking for safe access
			now := time.Now() // Capture time at start of each tick
			for key, val := range c.cache {
				if now.Sub(val.createdAt) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock() //Need to explicity unlock, otherwise won't ever unlock since this is an endless loop
		case <-c.quit: // Exit the loop upon receiving a quit signal
			return
		}
	}
}

func (c *Cache) Stop() {
	close(c.quit) // Signal the quit channel to stop the goroutine
}
