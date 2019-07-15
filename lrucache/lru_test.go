package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRU_Set(t *testing.T) {

	t.Run("Set stores item in an internal storage", func(t *testing.T) {
		c := New(10)

		evicted := c.Set("foo", "bar")

		item := c.storage["foo"]
		data, ok := item.value.(string)
		assert.False(t, evicted)
		assert.True(t, ok)
		assert.Equal(t, "bar", data)
	})

}

func TestLRU_Get(t *testing.T) {
	t.Run("Get from empty cache", func(t *testing.T) {
		c := New(1)

		val, ok := c.Get("foo")

		assert.False(t, ok)
		assert.Nil(t, val)
	})
	t.Run("Get item is in cache", func(t *testing.T) {
		c := New(1)
		c.Set("foo", "bar")

		item, ok := c.Get("foo")

		val, _ := item.(string)
		assert.True(t, ok)
		assert.Equal(t, "bar", val)

	})
}

func Test_DoubleLinkedList(t *testing.T) {
	assert := assert.New(t)
	t.Run("Add to a list for empty", func(t *testing.T) {
		l := &doubleLinkedList{}
		item := l.addItem("foo", 42)
		assert.Equal("foo", item.key)
		assert.Equal(l.head, item)
		assert.Equal(l.tail, item)
	})
	t.Run("Add to list with items", func(t *testing.T) {
		existing := &cacheItem{
			value: "foo",
		}
		l := &doubleLinkedList{
			head: existing,
			tail: existing,
		}
		new := l.addItem("bar", "baz")
		assert.Equal(new, l.head)
		assert.Equal(l.tail, l.head.next)
		assert.Equal(existing, l.tail)
	})

	t.Run("moveToHead. for head does nothing", func(t *testing.T) {
		item := &cacheItem{
			value: "foo",
		}
		l := &doubleLinkedList{
			head: item,
			tail: item,
		}
		l.moveToHead(item)
		assert.Equal(item, l.head)
	})

	t.Run("moveToHead. moves a middle element to head", func(t *testing.T) {
		first := &cacheItem{
			value: 1,
		}
		second := &cacheItem{
			value: 2,
		}
		third := &cacheItem{
			value: 3,
		}
		first.next = second
		second.prev = first
		second.next = third
		third.prev = second

		l := &doubleLinkedList{
			head: first,
			tail: third,
		}
		l.moveToHead(second)

		assert.Equal(second, l.head)
		assert.Nil(l.head.prev)
		assert.Equal(first, second.next)
		assert.Equal(third, first.next)
		assert.Equal(first, third.prev)
		assert.Nil(third.next)
		assert.Equal(third, l.tail)
	})

	t.Run("moveToHead. moves a the last element to head", func(t *testing.T) {
		first := &cacheItem{
			value: 1,
		}
		second := &cacheItem{
			value: 2,
		}
		third := &cacheItem{
			value: 3,
		}
		first.next = second
		second.prev = first
		second.next = third
		third.prev = second
		l := &doubleLinkedList{
			head: first,
			tail: third,
		}

		l.moveToHead(third)

		assert.Equal(third, l.head)
		assert.Nil(third.prev)
		assert.Equal(second, l.tail)
		assert.Nil(second.next)
		assert.Equal(first, third.next)
		assert.Equal(third, first.prev)
	})
}

func Test_UseCases(t *testing.T) {

	t.Run("Cache evicts value if capacity reached", func(t *testing.T) {
		c := New(1)

		evicted1 := c.Set("foo", 42)
		evicted2 := c.Set("bar", 43)

		assert.False(t, evicted1)
		assert.True(t, evicted2)
		assert.False(t, c.Has("foo"))
	})
}
