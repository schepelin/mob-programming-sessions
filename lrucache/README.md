# General

[LRU visually explained](https://www.youtube.com/watch?v=DUbEgNw-F9c&t=1539s)

## Principles

- Test-Driven
- Consumer-First
- Refine existing code

## Implementation

Your assignment is to implement the next interface

```go
// Cache describes the common interface
type Cache interface {
    Set(key string, value interface{}) (evicted bool)
    Get(key string) (value interface{}, ok bool)
    Has(key string) (ok bool)
}
```

Example of cache Get/Set:

```go
c := NewCache(10)

_ = c.Set("foo", 42)

item, ok := c.Get("foo")

fmt.Println(ok) // true

val, ok := item.(int)
fmt.Println(val, ok) // 42, true
```

Examples of the eviction mechanics:

```go
c := NewCache(1)

if ok := c.Set("foo", 42); ok {
    fmt.Println("It will newer called")
}
if ok := c.Set("bar", 100500); ok {
    fmt.Println("Previous value was evicted")
}

c.Has("foo") // false
c.Has("bar") // true
```

```go
c := NewCache(1000)

// evicted is a bool value identify if capacity has reached
evicted := c.Set("somekey", "somevalue") // evicted == false

val, ok := c.Get("somekey")

// returns true if a value is in cache
ok := c.Has("somekey") // ok == true
```
