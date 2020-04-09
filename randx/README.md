# x/randx
> goroutine-safe rand

## Getting Started
```
go get -u github.com/goroom/x
```

## Example
```go
package main

import (
    "fmt"
    "github.com/goroom/x/randx"
)

func main() {
    r := randx.GetRand()
    fmt.Println(r.Intn(10))
    r.Release()
}
```