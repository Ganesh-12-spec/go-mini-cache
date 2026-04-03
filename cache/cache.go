package cache

import (
	"time"
)

type CacheItem struct {
	Value     string
	ExpiresAt time.Time
}

type Cache struct {
	items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.items[key] = CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	item, exists := c.items[key]

	if !exists {
		return "", false
	}

	if time.Now().After(item.ExpiresAt) {
		delete(c.items, key)
		return "", false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

func (c *Cache) Clear() {
	c.items = make(map[string]CacheItem)
}

func (c *Cache) Size() int {
	return len(c.items)
}

func (c *Cache) Cleanup() int {
	removed := 0
	for key, item := range c.items {
		if time.Now().After(item.ExpiresAt) {
			delete(c.items, key)
			removed++
		}
	}
	return removed
}