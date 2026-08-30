package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache {
		cacheEntries: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, exists := cache.cacheEntries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}
