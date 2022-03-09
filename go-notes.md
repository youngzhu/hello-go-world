# 知识点

## 001 关于 `main` 包
只有 `package main`，`build` 时才有 exe 文件，`install` 时才有 bin 文件夹。

## 002 方法和函数居然不是一个东西。。
方法能给用户定义的类型添加新的行为。
方法实际上也是函数只是在声明时，在关键字`func`和函数名之间多了一个参数，这个参数就接收者。
接收者，将函数与接收者的类型绑定在一起。
如果一个函数有接收者，这个函数就被称作方法。
via Go In Action

## 003 switch还能这么用
```Go
func isShellSpecialVar(c uint8) bool {
    switch c {
        case '*', '#', '$', '@', '!', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return true
    }
    return false
}
```

## 004 数组和切片作为参数在方法之间传递：数组是值（副本）传递，切片是指针（引用）传递
可参见：/book1/ch02/array_vs_slice.go

## 005 初始化二维数组/切片
```Go
m, n := 3, 4
a := make([][]int, m) // 二维切片，3行
for i := range a {
    a[i] = make([]int, n) // 每一行4列
}
```

## 006 `new`和`make`的区别
简单来说： `new` 请求分配内存，返回指针，而 `make` 只用于初始化切片，映射，channel。