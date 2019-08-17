# 練習-3-如何使用子 Package

引用 **子組建** 或引用非官方提供的組建 Package

```go
import (
	"fmt"

	"github.com/<your-repo>/go-training/example-3-use-sub-package/helloworld"
)

func main() {
	fmt.Println(helloworld.HelloWorld())
}
```

子組建的路徑引用方式 `github.com/<your-repo>/go-training/example-3-use-sub-package/helloworld`

子組建的函式命名會是會變成 **PascalCase** `helloworld.HelloWorld())` 

```bash
$ go run ./main.go
> Hello World!
```

若只需要呼叫 **子組建** 的初始函式，未使用內部其他函式時，要用 `_` 引入

例如：

```go
import (
	_ "github.com/<your-repo>/go-training/example-3-use-sub-package/helloworld"
)
```

這樣將會呼叫 **子組建** 的初始函

```go
package helloworld

func init() string {
	fmt.Println("func init!")
}
```
