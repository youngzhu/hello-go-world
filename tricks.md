# 小技巧

## 001 不要提前写 `import`，甚至都不用自己写
有些工具（如VSCode）可以自动引入。
提前写还可能导致副作用，如VSCode，在保存时一直提示 `Get code actions from "Go"`，导致修改无法保存。

## 002 `os.Open()`老是报错？
使用`log.Println(os.Getwd())`可以查看当前的工作目录，就可以找到你要打开的文件的相对路径