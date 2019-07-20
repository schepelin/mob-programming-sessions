package lrucache

import (
	"errors"
)

type CacheI interface {
	Set(key string, value interface{}) (evicted bool)
	Get(key string) (value interface{}, ok bool)
	Has(key string) (ok bool)
}

type Cache struct {
	values   map[string]*Node
	capacity uint
	head     *Node
	tail     *Node
}
type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func NewCache(capacity uint) (*Cache, error) {
	if capacity == 0 {
		return nil, errors.New("Erorr! Cap 0")
	}
	return &Cache{
		values:   make(map[string]*Node),
		capacity: capacity,
	}, nil
}

func (c *Cache) Has(key string) bool {
	if _, ok := c.values[key]; ok {
		return true
	}
	return false
}

func (c *Cache) Set(key string, value interface{}) bool {
	n := &Node{value: value}
	if c.head == nil {
		c.head = n
		c.tail = n
	} else{
		c.head.prev = n
		n.next = c.head
		c.head = n
	}

	_, ok := c.values[key]
	if !ok && len(c.values) == int(c.capacity) {
		// TODO: eviction
	}
	c.values[key] = n

	return len(c.values) > int(c.capacity)
}

func (c *Cache) popTail() *Node{
	if c.tail == nil{
		return nil
	}
	n := c.tail
	c.tail = c.tail.prev
	if c.tail == nil{
		c.head = nil
	}
	n.next = nil
	n.prev = nil
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if v, ok := c.values[key]; ok {
		return v.value, true
	}
	return nil, false
}

// func (c *Cache) Len() int {
// 	return len(c.values)
// }
