package lrucache

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cache_Empty_Cache_Return_False(t *testing.T) {
	assert := assert.New(t)
	t.Run("", func(t *testing.T) {
		c, _ := NewCache(10)

		assert.False(c.Has("unknown_key"))
	})
	t.Run("Test_Cache_First_El_Doesnt_Evicted", func(t *testing.T) {
		c, _ := NewCache(10)

		assert.False(c.Set("key", 1))
	})
	t.Run("Test_Cache_Cache_With_Zero_Cap_Not_Allow", func(t *testing.T) {
		_, err := NewCache(0)
		assert.Error(err)
	})
	t.Run("Same key doesn't increase sixe", func(t *testing.T) {
		c, err := NewCache(10)
		assert.NoError(err)
		evicted := c.Set("key", 1)
		assert.False(evicted)
		evicted = c.Set("key", 1)
		assert.False(evicted)
	})
}

func TestCache_Has(t *testing.T) {
	tests := []struct {
		name     string
		capacity uint
		key      string
		want     bool
	}{
		{
			name:     "",
			capacity: 1,
			want:     false,
		},
		{
			name:     "",
			capacity: 1,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewCache(tt.capacity)
			if got := c.Has(tt.key); got != tt.want {
				t.Errorf("Cache.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Set(t *testing.T) {
	tests := []struct {
		name     string
		capacity uint
		key      string
		value    interface{}
		want     bool
	}{
		{
			name:     "Empty cache's Set returns false",
			key:      "a",
			value:    1,
			want:     false,
			capacity: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewCache(tt.capacity)
			if got := c.Set(tt.key, tt.value); got != tt.want {
				t.Errorf("Cache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Set_ot_Seregi(t *testing.T) {
	key := "razdva"
	c, err := NewCache(1)
	assert.NoError(t, err)
	evicted := c.Set(key, 1)
	assert.False(t, evicted)
	assert.True(t, c.Has(key))
}

func TestCache_Get(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		capacity uint
		want     interface{}
		want1    bool
	}{
		{
			name: "Simple Get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{}
			got, got1 := c.Get(tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Cache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCache_Get_ot_Pashi(t *testing.T) {
	key := "razdva"
	c, err := NewCache(1)
	assert.NoError(t, err)
	evicted := c.Set(key, 1)
	assert.False(t, evicted)
	val, ok := c.Get(key)
	assert.True(t, ok)
	assert.Equal(t, int(1), val)
}

func TestCache_Set_Eviction(t *testing.T) {
	key := "razdva"
	c, err := NewCache(1)
	assert.NoError(t, err)
	evicted := c.Set("keyevicted", 1)
	assert.False(t, evicted)
	evicted = c.Set(key, 1)
	assert.True(t, evicted)
}
