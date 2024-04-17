package lru

import (
	"sync"
)

// LRUCacheMap LRU 缓存
type LRUCacheMap struct {
	cache map[string]any
	mutex sync.RWMutex
}

// NewLRUCache 创建一个新的 LRUCacheMap
func NewLRUCacheMap() *LRUCacheMap {
	return &LRUCacheMap{
		cache: make(map[string]any),
	}
}

// Get 从缓存中获取值
func (c *LRUCacheMap) Get(key string) any {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if elem, exists := c.cache[key]; exists {
		return elem
	}
	return nil
}

// Add 添加值到缓存
func (c *LRUCacheMap) Add(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.cache) <= 100 {
		c.cache[key] = value
	}
}
