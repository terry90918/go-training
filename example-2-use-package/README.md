# 練習-2-如何使用Package

```bash
$ go run ./main.go ./hello-wrold.go
> Hello World!
```

不可以在不同的目錄下引用，子目錄除外

例如：

```go
package test

func helloworld() string {
	return "Hello Wrold!"
}
```

`package test` 會提示 `can't load package: package <repo>: found packages test (hello-wrold.go) and main (main.go) in <repo>`

引用 **子目錄** 的使用方式