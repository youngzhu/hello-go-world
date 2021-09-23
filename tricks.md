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

## 004 