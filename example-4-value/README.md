# 宣告變數

宣告變數方式

```go
package main

import "fmt"

var foo string
var bar int

func main() {
  foo = "The number is "
  bar = 1
	fmt.Println(foo, bar)
}
```

```go
package main

import "fmt"

var (
  foo string
  bar int
)

func main() {
  foo = "The number is "
  bar = 1
	fmt.Println(foo, bar)
}
```

```go
package main

import "fmt"

var (
  foo string = "The number is "
	bar int    = 1
)

func main() {
	fmt.Println(foo, bar)
}
```

```go
package main

import "fmt"

func main() {
  var foo string
  foo = "The number is "
  bar := 1
  fmt.Println(foo, bar)
}
```

```go
package main

import "fmt"

func main() {
  foo := "The number is "
  bar := 1
  fmt.Println(foo, bar)
}
```

```go
package main

import "fmt"

const (
  one = iota + 1
  two
  three
)

func main() {
  fmt.Println(one, two, three) // 1 2 3
}
```
