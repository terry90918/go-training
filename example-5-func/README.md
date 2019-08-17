# 如何使用 go func

```go
package main

import "fmt"

func add(i, j int) int {
  return i + j
}

func main() {
	fmt.Println(add(1, 2))
}
```

```go
package main

import "fmt"

// 因為回傳兩個數值，所以最後回傳參數部分是兩個 int
func swarp(i, j int) (int, int) {
  return j, i
}

func main() {
  a := 1
  b := 2
  a, b = swarp(a, b)
  fmt.Println("a: ", a, "b: ", b)
  
  a, b = b, a
  fmt.Println("a: ", a, "b: ", b)
}
```

```go
package main

import "fmt"

func foo() func() int {
  return func() int {
    return 100
  }
}

func main() {
  bar := foo()
	fmt.Println("%T\n", bar)
	fmt.Println(bar)
}
```

```go
package main

import "fmt"

func main() {
	bar := func(i, j float32) float32 {
		return i + j
	}
	fmt.Printf("%T\n", bar)
	fmt.Println(bar(2.02, 4.4+2.28))
}
```

```go
package main

import "fmt"

func main() {
	foo := func() string {
		return "foo"
	}
	fmt.Println(foo())

	bar := func() {
		fmt.Println("bar")
	}
	bar()

	func() {
		fmt.Println("dog")
	}()
}
```

```go
package main

import (
	"fmt"
	"strings"
)

type searchOpts struct {
	username string
	email    string
	sexy     int
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}
	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}
	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}
	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = '%d'", opts.sexy))
	}
	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "terry",
		email:    "test@mail.com",
		sexy:     1,
	}))
}
```