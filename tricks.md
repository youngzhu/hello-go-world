# 小技巧

## 001 不要提前写 `import`，用到了再写
有些工具（如VSCode）可以自动引入。
提前写还可能导致副作用，如VSCode，在保存时一直提示 `Get code actions from "Go"`，导致修改无法保存。

用Go插件，可以有自动导入/格式化等功能，副作用就是会提示 `Get code actions from "Go"`，保存很慢。
也可以不借助插件，完全手动输入。
孰优孰劣，自己权衡了。

## 002 `os.Open()`老是报错？
使用`log.Println(os.Getwd())`可以查看当前的工作目录，就可以找到你要打开的文件的相对路径

## 003 魔法字符串 "2006-01-02 15:04:05"
使用`time.Now().Format("2006-01-02 15:04:05")`可以获得当前时间的_yyyy-MM-dd HH:mm:ss_格式
这是Golang诞生的时间

这个方法也让我意识到，一个类型（type）和它的方法不一定要在一个go文件里。例如这里的Time类在time.go中，而Format方法在format.go中

## 004 使用 flag 解析命令行参数时，bool类型的放最后
```Go
// 这样的输入，会导致n和t都取不到值
go run sort_compare.go -a1 insertion -a2 selection -s true -n 1000 -t 100
```
通过在StackOverflow上提问知道了，bool只看是否存在，如 `ls -l`中的`-l`
所以，正确的写法是：
```Go
// 表示s为true
go run sort_compare.go -a1 insertion -a2 selection -s -n 1000 -t 100
// 表示s为false
go run sort_compare.go -a1 insertion -a2 selection -n 1000 -t 100
```

## 005 3种方法将整数转为字符串
**涉及到string的转换，首先`strconv`**

1. `fmt.Sprintf("%d", number)`
2. `strconv.FormatInt(number, 10)`
3. `strconv.Itoa(number)`
通过基准测试，效率排名为：2-3-1

## 006 example_test时，空格不要写在末尾
```go
fmt.Printf(" %v", v) // pass

fmt.Printf("%v ", v) // fail
```

## 007 切片的空判断
要检查切片是否为空，请始终使用 `len(s) == 0`。而非 `nil`
