# generago 介绍

generago 是一个代码生成器的库， 目的是为了将开发者平时重复的代码批量生成。

目前支持的功能如下：

-   批量生成 go 和数据库之间的映射

下一步将要增加的功能：

-   实现 dao 层的增删改查

# 使用

该库保持简约的风格， 用户不用填写十分繁琐的配置，即可立即使用， 做到了真正的开箱即用。

步骤：

1. 添加如下配置， 使用者需要根据自己的配置填写

```go
func main() {
	generago.SetDataSource(generago.DataSource{
		DbName:   "数据库",
		User:     "用户",
		Password: "密码",
	})
	generago.SetOutDir("./model")
	generago.Execute()
}
```

2. 执行 `go run xx.go`, xx 表示 main 包所在的文件，然后就会在指定目录下生成对应的文件。

3. 不要忘记在生产环境中删除该库，减少项目的体积。

Tip：

1. `./` 表示的是运行时目录，使用者不要忘记哦!
