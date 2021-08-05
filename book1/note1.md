# 《Go语言编程》
许式伟 吕桂华 等编著

## ch00 前言
### 并发执行的“执行体”
进程（Process） - 线程（Thread） - 协程/轻量级线程（Coroutine）

Go在语言级别上支持协程，称作goroutine。这也是其名字的由来。

### Go的特点
+ 代码风格强制统一：
    - public的变量必须大写字母开头
    - private的变量必须小写字母开头
    - {}
        ```Go
        // 正确的且必须的
        if exp {

        }

        // 错误的
        if exp
        {

        }
        ```
+ 函数可以返回多个值

### 原则
> 如果一个特性并不对解决任何问题有显著的价值，那么Go就不提供它。

## ch01 初识

### 函数
```Go
func  函数名(参数列表) (返回值列表) {
    // 函数体
}

func Compute(value1 int, value2 float64) (result float64, err error) {
    // body
}
```

### 注释
```Go
/*
块注释
*/

// 行注释
```

## ch02 顺序编程
Go —— 更好的C语言

### 变量申明
```Go
var 变量名 类型

// 也可以将申明的变量放一起
var (
    v1 int
    v2 string
)
```

### 变量初始化
```Go
var v1 int = 10
var v2 = 10
v3 := 10 // := 操作符， 同时进行变量申明和初始化的工作
// 以上三种写法效果是一样的

/*
需要注意的是 := 左侧的变量不能是已经被申明过的，否则会报编译错误
例如
*/
var i int
i := 2
```

### 变量赋值
```Go
var v10 int
v10 = 123

// 多重赋值
// 在不支持多重赋值的语言中，交换两个变量
t = i;
i = j;
j = t;

// 而在Go中
i, j = j, i
```

### 常量
```Go
// 可以限定类型，也可以不要
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"
```

### 数组
```Go
// 数组申明
[32] byte
[3][5] int

// 数组属于值传递，当作为一个参数传递到另一个函数里，传递的是个副本
```

### 数组切片
创建数组切片的两种方式：
1. 基于数组，见 `slice.go`
2. 直接创建

```Go
// arr[first:last]
selice = arr[:] // 所有元素创建切片
selice = arr[:5] // 前5个元素创建切片
selice = arr[5:] // 从第5个元素开始创建切片

// 直接创建，内置函数make()
// 创建一个初始元素个数为5的数组切片，元素初始值为0
slice1 := make([]int, 5)
// 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
slice2 := make([]int, 5, 10)
// 直接创建并初始化包含5个元素的切片
slice3 := []int{1, 2, 3, 4, 5}
```

切片的遍历
```Go
// 传统方法
for i := 0; i < len(slice); i++ {
    fmt.Println("slice[", i, "] = ", slice[i])
}

// range
// range表达式有两个返回值，第一个是索引，第二个是元素值
for i, v := range slice {
    fmt.Println("slice[", i, "] = ", v)
}
```

内置函数 `cap()` 和 `len()`，见 `slice2.go`
cap: 返回切片分配的空间大小
len: 返回当前所有存储的元素个数

给切片新增元素：`append()`函数
```Go
// 直接增加元素
mySlice = append(mySlice, 1, 2, 3)

// 增加另一个切片中的元素
mySlice2 := []int{8, 9, 10}
// 必须有 ... ，否则编译错误
mySlice = append(mySlice, mySlice2...)
// 上两行等同于下一行
mySlice = append(mySlice, 8, 9, 10)
```

基于切片数组创建切片数组
```Go
oldSlice := []int{1, 2, 3, 4, 5}
newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新的切片数组
```

内容复制
内置函数`copy()`，见 `copy.go`

### map
```Go
// 1. 声明
var myMap map[string] PersonInfo
// myMap - 变量名
// string - 键类型
// PersonInfo 值类型

// 2 创建
// 2.1 键类型为string，值类型为PersonInfo
myMap = make(map[string] PersonInfo)

// 2.2 指定cap
myMap = make(map[string] PersonInfo, 100)

// 2.3 创建并初始化
myMap = map[string] PersonInfo {
    "1234": PersonInfo{"1", "Jack", "Room 101, ..."},
}

// 3 赋值
myMap["123"] = PersonInfo{"1", "Jack", "Room 101, ..."}

// 4 删除
// 如果key不存在，没有影响
// 如果key=nil，将抛出异常
delete(myMap, "123")

// 5 查找
value, ok := myMap["123"]
if ok { // 找到了
    // 处理找到的value
}
```

### 2.4 流程控制
条件语句，注意点：
1. 条件不需要`()`
2. `{}`不能省略，即使只有一条语句
3. `{`必须与`if`或`else`处于同一行
4. `if`之后，条件语句之前，可以添加变量初始化语句，用`;`分隔
5. 在有返回值的函数中，不允许将“最终的” `return`包含在`if-else`中

```Go
if a < 5 {
    return 0
} else {
    return 1
}

// 注意点5：编译失败
func example(x int) int {
    if x == 0 {
        return 5
    } else {
        return x
    }
}
```

选择语句——switch的结构特点：
- `{`必须与`switch`处于同一行
- 条件表达式不限定为常量或整数
- 单个`case`中可以出现多个结果项
- 不需要**break**来明确退出一个`case`
- 只有明确添加关键字 `fallthrough` 才会继续执行紧跟的下一个`case`
- 可以不设定`switch`后面的表达式，相当于`if-else`

```Go
switch i {
    case 0:
        fmt.Printf("0")
    case 1:
        fmt.Printf("1")
    case 2:
        fallthrough
    case 3:
        fmt.Printf("3")
    case 4, 5, 6:
        fmt.Printf("4, 5, 6")
    default:
        fmt.Printf("Default")
}

switch {
    case 0 <= Num && Num <= 3:
        fmt.Printf("0-3")
    case 4 <= Num && Num <= 6:
        fmt.Printf("4-6")
}
```

循环语句—— 只支持 `for`
注意点：
- `{`必须与`for`处于同一行
- `break`支持标签

```Go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

// 简化，无限循环
sum := 0
for {
    sum ++
    if sum > 100 {
        break
    }
}

// 多重赋值
a := []int{1, 2, 3, 4, 5, 6}
for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
    a[i], a[j] = a[j], a[i]
}

for j := 0; j < 5; j++ {
    for i := 0; i < 10; i++ {
        if i > 5 {
            break JLoop
        }
        fmt.Println(i)
    }
}
JLoop:
//...
```

跳转语句 `goto`
```Go
func myfunc() {
    i := 0
    HERE:
    fmt.Println(i)
    i++
    if i < 10 {
        goto HERE
    }
}
```

### 2.5 函数
函数的 定义
```Go
func Add(a int, b int) (ret int, err error) {
    if a < 0 || b < 0 { // 只支持两个非负数的加法
        err = errors.New("Should be non-negative numbers!")
        return
    }
    return a + b, nil // 支持多重返回
}

// 若类型相同可省略前面的类型声明
func Add(a, b int) (ret int, err error) {

}

// 如果只有一个 值
func Add(a, b int) int {

}
```

不定参数
```Go
func myfunc(args ...int) {
    for _, arg := range args {
        fmt.Println(arg)
    }
}

// 调用
myfunc(2, 3, 4)
myfunc(1, 3, 7, 13)

// 传递
func myfunc3(args ...int) {

}

func myfunc(args ...int) {
    // 原样传递
    myfunc3(args...)

    // 传递片段
    myfunc3(args[1:]...)
}

```

任意类型的不定参数，见 `varg1.go`

可以定义多返回值，不需要的时候，可以不接收
```Go
func MyFunc(i int) (o int, err error)

// 不需要时，用 _ 接收返回
result, _ := MyFunc(100)
```

匿名函数
```Go
// 匿名函数由一个不带函数名的函数声明和函数体组成
func(a, b int, z float64) bool {
    return a*b < int(z)
}

// 匿名函数可以直接赋值给一个变量或直接执行
f := func(x, y int) int {
    return x+y
}

func (ch chan int) {
    ch <- ACK
} (reply_chan) // 花括号后直接跟参数列表 表示函数的调用
```

闭包，参见代码，closure.go 。
不知道在讲什么。。

### 2.6 错误处理
#### 2.6.1 error接口
```Go
// 接口定义
type error interface {
    Error() string
}

// 大多数函数，定义遵循以下模式：将error最后一个返回值
func Foo(param int) (n int, err error) {
    // ...
}

// 错误处理
n, err := Foo(0)

if err != nil {
    // 错误处理
} else {
    // 使用返回值 n
}
```

#### 2.6.2 defer