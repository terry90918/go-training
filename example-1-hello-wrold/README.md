# 練習-1-Hello-World

VSCode 偏好設定

```json
{
  "go.formatTool": "goimports",
  "go.goroot": "/usr/local/Cellar/go/1.12.7/libexec",
  "go.gopath": "/Users/<user-name>/Documents/code-workspace"
}
```

`"go.formatTool": "goimports"` 可以自動引入使用到的 package

```
$ go run ./main.go
> Hello World!
```