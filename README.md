# Assert

Simple assert library for [GO](https://go.dev)


## Simple usage withing test

```go
package main

import "testing"
import . "github.com/chapgx/assert"

func Test(t *testing) {
  db , e := Connet() // some type of connection
  AssertT(t, e == nil, "e should be nil")
}
```


## Simple usage in code

```go
package main

import "os"
import . "github.com/chapgx/assert"

func main() {
  port := os.Getenv("PORT")
  Assert(os != "", "environment port not found")
}
```
