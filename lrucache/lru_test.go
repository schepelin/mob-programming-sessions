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
