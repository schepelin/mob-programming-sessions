# General

[LRU visually explained](https://www.youtube.com/watch?v=DUbEgNw-F9c&t=1539s)

## Principles

- Test-Driven
- Consumer-First
- Refine existing code

## Implementation

Your assignment is to implement the next interface

```go
// Cache interface
type Cache interface {
    Set(key string, value interface{}) (evicted bool)
    Get(key string) (value interface{}, ok bool)
    Has(key string) (ok bool)
}
```

Examples of `Has` method:

```go
c := NewCache(10)
c.Has("a key") // false

c.Set("answer", 42)
c.Has("anwser") // true
```

Example of cache `Get` and `Set` methods:

```go
c := NewCache(10)

_ = c.Set("foo", 42)

item, ok := c.Get("foo")

fmt.Println(ok) // true

val, ok := item.(int)
fmt.Println(val, ok) // 42, true
```

Examples of `Set` method with eviction:

```go
c := NewCache(1)

if ok := c.Set("foo", 42); ok {
    fmt.Println("It will not be printed")
}
if ok := c.Set("bar", 100500); ok {
    fmt.Println("Previous value was evicted")
}

c.Has("foo") // false
c.Has("bar") // true
```
