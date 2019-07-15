package lrucache

// cacheItem represents a value stored in a cache
type cacheItem struct {
	value interface{}
	key   string
	prev  *cacheItem
	next  *cacheItem
}

type doubleLinkedList struct {
	head *cacheItem
	tail *cacheItem
}

func (dl *doubleLinkedList) addItem(key string, val interface{}) *cacheItem {
	item := &cacheItem{
		value: val,
		key:   key,
		next:  dl.head,
	}
	if dl.head != nil {
		dl.head.prev = item
	}
	if dl.tail == nil {
		dl.tail = item
	}
	dl.head = item

	return item
}

func (dl *doubleLinkedList) moveToHead(item *cacheItem) {
	if dl.head == item {
		return
	}
	prev := item.prev
	next := item.next

	if dl.tail == item {
		dl.tail = prev
	}

	item.prev = nil
	item.next = dl.head
	dl.head.prev = item
	dl.head = item

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}

func (dl *doubleLinkedList) popTail() *cacheItem {
	item := dl.tail
	item.prev = nil
	dl.tail = item.prev
	if dl.tail == nil {
		dl.head = nil
	}
	if dl.tail != nil {
		dl.tail.next = nil
	}
	return item
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
	list     doubleLinkedList
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
	item := lru.list.addItem(key, val)

	lru.size++
	lru.storage[key] = item

	if lru.size <= lru.capacity {
		return false
	}
	last := lru.list.popTail()
	delete(lru.storage, last.key)
	lru.size = lru.capacity

	return true
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
	lru.list.moveToHead(item)

	return item.value, true
}
