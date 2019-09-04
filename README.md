# Mob programming

## Principles

### Test-Driven

Write a test before you write any line of code

### Consumer-first vs Build-up

```go

// build up
type Point struct {
    x int
    y int
}

// consume first
func NewPoint(x,y int) *Point{

}
```

### Refine existing code

Do not throw existing code, refine it!
Each iteration should make a step forward.

## Links

- [Blog: Mob programming](http://mobprogramming.org/)
- [Video: A whole tea, approach](https://youtu.be/SHOVVnRB4h0)
- [Book: The mob programming guidebook](http://www.mobprogrammingguidebook.com/)
- [BooK: Mob programming](https://leanpub.com/mobprogramming)
