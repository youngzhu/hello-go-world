# 学习过程中的疑问

_注：答案不是100%准确，是基于仅有的经验的总结。_

## Q1: build/run命令和gopath及项目目录之间的关系
~~在如今IDE盛行的年代好像也没多大必要。~~  
项目不是在IDE中运行的

可以直接运行`go run main.go`  
也可以先编译再执行
```
go build main.go
./main
```

## Q2: *和&的区别和用途
### &
获取地址运算符。`&foo`，获取值类型变量foo的地址
定义一个指针变量
    lisa := &user{"Lisa", "lisa@mail.com"}
### *
1. 声明方法时，表示指针接收者
    func (u *user) changeEmail
2. 表示指针变量指向的值
    (*lisa) // lisa是指针变量

## Q3: package和目录名必须一致吗？
必须一致。可参见ch03 SMP，根据原文应该是path：`mlib`，package：`libaray`，但怎么都编译不通过。将path改成`libaray`就好了。  
main则不必。

~~其实想多了，main包就不在main目录下。~~  以偏概全了

## Q4：项目里要不要src目录？
Go In Action 一书中写道
> Go语言新手常犯的一个错误是，在公共代码库里创建一个名为src或code的目录。...只需要把包的源文件放在共用代码库的根目录就好。

这里是不是有一个特定的场景：创建想要分享的代码库的时候？
如果没有src目录，在引入非标准库的包时，无法调试啊。
怎么破？

update 2021-09-03
用`src`吧。官方都这么用了，`golang.org/src/os/file.go`

## Q5: Go语言里有继承的概念吗？
没有。也不需要。同一个包内，方法均可见。
嵌入类型也能达到继承的效果。

## Q6: Go语言中有范型吗？

## Q7: Go语言中有断言吗？
没有。但可以用下列函数代替
```Go
if condition {
  panic(err)
}
```

## Q8: `panic`是做什么用的？

## Q9: `init()`是所有文件都可以有吗？还是仅main文件？
实测证明，所有文件里都能放，都有效

## Q10: 如何添加外部依赖
`go mod tidy` 或者 `go get ./...`
