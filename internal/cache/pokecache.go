package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	caches map[string]cacheEntry
	Mu     sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		caches: make(map[string]cacheEntry),
		Mu:     sync.Mutex{},
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.caches[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	ce, ok := c.caches[key]

	if !ok {
		return nil, false
	}

	return ce.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	cacheClearTicker := time.NewTicker(interval)

	for range cacheClearTicker.C {
		currentTime := time.Now()
		for k, v := range c.caches {
			if currentTime.Sub(v.createdAt) > interval {
				delete(c.caches, k)
			}
		}
	}
}
