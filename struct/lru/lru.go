package lru

import (
	"container/list"
)

// LRUCache LRU 缓存 不安全 使用外部锁
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	lruList  *list.List
}

// _Entry 表示缓存中的条目
type _Entry struct {
	key   string
	value any
}

// NewLRUCache 创建一个新的 LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

// Get 从缓存中获取值
func (c *LRUCache) Get(key string) any {
	if elem, exists := c.cache[key]; exists {
		c.lruList.MoveToFront(elem)
		return elem.Value.(*_Entry).value
	}
	return nil
}

// Add 添加值到缓存
func (c *LRUCache) Add(key string, value any) {
	if elem, exists := c.cache[key]; exists {
		c.lruList.MoveToFront(elem)
		elem.Value.(*_Entry).value = value
	} else {
		for c.lruList.Len() >= c.capacity {
			c.removeOldest()
		}
		entry := &_Entry{key, value}
		newElem := c.lruList.PushFront(entry)
		c.cache[key] = newElem
	}
}

// ScalingChange 扩/缩容
func (c *LRUCache) ScalingChange(capacity int) {
	c.capacity = capacity
	for c.lruList.Len() >= c.capacity {
		c.removeOldest()
	}
}

// removeOldest 移除最久未使用的项
func (c *LRUCache) removeOldest() {
	if c.lruList.Len() == 0 {
		return
	}
	oldestElem := c.lruList.Back()
	if oldestElem != nil {
		delete(c.cache, oldestElem.Value.(*_Entry).key)
		c.lruList.Remove(oldestElem)
	}
}
