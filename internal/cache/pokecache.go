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

func NewCache(interval time.Duration) Cache {

	return Cache{
		caches: make(map[string]cacheEntry),
		Mu:     sync.Mutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.caches[key] = cacheEntry{
		createdAt: time.Time{},
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

func reapLoop() {

}
