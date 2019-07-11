package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cache_Methods(t *testing.T) {
	assert := assert.New(t)
	t.Run("Empty cache's Has returns false", func(t *testing.T) {
		c := NewCache(10)

		assert.False(c.Has("unknown_key"))
	})

}
