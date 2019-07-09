# General

[LRU visually explained](https://www.youtube.com/watch?v=DUbEgNw-F9c&t=1539s)

## Implementation

User can create a cache instance with predefined capacity
For simplicity cache can store string to string mapping

```go
c := Cahce(1000)

# evicted is a bool value identify if capacity has reached
evicted := c.Set("somekey", "somevalue")

val, ok := c.Get("somekey")

# returns true if a value is in cache
ok := c.Has("somekey")
```
