package lrucache

// cacheItem represents a value stored in a cache
type cacheItem struct {
	value interface{}
}

// Cache describes a general interface of the cache
type Cache interface {
	Set(key string, value interface{}) (evicted bool)
	Get(key string) (value interface{}, ok bool)
	Has(key string) (ok bool)
}

// LRU implements the cache
type LRU struct {
	size     uint
	capacity uint
	storage  map[string]*cacheItem
}

// New is a factory function for LRU struct
func New(capacity uint) *LRU {
	return &LRU{
		capacity: capacity,
		storage:  make(map[string]*cacheItem),
	}
}

// Set puts value into a cache
func (lru *LRU) Set(key string, val interface{}) bool {
	lru.storage[key] = &cacheItem{
		value: val,
	}
	if lru.size > lru.capacity {
		//TODO: run eviction here
		return true
	}
	lru.size++
	return false
}

// Has checks if a key in the cache
func (lru *LRU) Has(key string) bool {
	_, ok := lru.storage[key]
	return ok
}

// Get returns value if it's presented in the cache
func (lru *LRU) Get(key string) (interface{}, bool) {
	item, ok := lru.storage[key]

	if !ok {
		return nil, false
	}
	// TODO: Update least recently used

	return item.value, ok
}
