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
类似于Java中的 `finally`。
需要注意的是，defer的调用遵照“先进后出”的原则，即最后一个defer（同一个函数中）将最先被执行。

只不过，当你需要为到底哪个先执行这种细节而烦恼的时候，说明代码的架构可能需要调整一下了。

```Go
// 即使函数执行过程中出现异常，dstFile和srcFile都会被关闭
func CopyFile(dst, src string) (w int64, err error) {
    srcFile, err := os.Open(src)
    if err != nil {
        return
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dstFile.Close()

    return io.Copy(dstFile, srcFile)
}

// 如果一句话做不完清理工作，也可以在 defer 后使用一个匿名函数
defer func() {
    // 复杂的清理工作
} ()
```

#### 2.6.3 panic()和recover()
没太明白。后面再看吧。
Go的两个内置函数，处理异常用的。

## ch03 面向对象编程
### 3.1 类型
Go语言和C语言一样，类型都是基于值传递的。想要修改变量的值，只能传递指针。
详见 value_pointer.go

Go语言中的数组和基本类型没有区别，都是值类型（传递）。
详见 array_value.go

### 3.2 初始化
```Go
//矩形类型
type Rect struct {
    x, y float64
    width, height float64
}
// 初始化
rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width:100, height:200}

// Go语言中没有构造函数的概念
// 可由一个全局的创建函数来创建对象，通常 NewXXX 来命名

// 重要的是格式，或称作模板
func NewRect(x, y, width, height float64) *Rect {
    return &Rect{x, y, width, height}
}

```

### 3.3 匿名组合
关于继承的，看的稀里糊涂。有必要的话，再回头看吧。

### 3.4 可见性
首字母大写表示公开，小写表示私有。
需要注意的是，可访问性是包（package）一级的而不是类型（type）一级的。
也就是说，尽管小写的是私有的，但在同一个包中的其他类型也可以访问。

### 3.5 接口
只要实现了接口中的方法就算实现了接口，没有`implements`之类的关键字

#### 3.5.3 接口赋值
将对象实例赋值给接口
要求：对象实现了接口所有的方法
```Go
type LessAdder interface {
    Less(b Integer) bool
    Add(b Integer)
}

type Integer int

func (a Integer) Less(b Integer) bool {
    return a < b
}

func (a *Integer) Add(b Integer) {
    *a += b
}

var a Integer = 1
var b LessAdder = &a // 正确
var c LessAdder = a // 错误，因为 Add 方法需要引用传递

type Lesser interface {
    Less(b Integer) bool
}

var a Integer = 1
var b Lesser = &a // 两种赋值都可以
var c Lesser = a

```

将一个接口赋值给另一个接口
只要两个接口拥有相同的方法列表（顺序不重要），那么它们就是等价的，可以互相赋值。
```Go
// 第一个接口
package one

type ReadWriter interface {
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
}

// 第二个接口
package two

type IStream interface {
    // 顺序和接口1相反
    Write(buf []byte) (n int, err error)
    Read(buf []byte) (n int, err error)
}

// 赋值
var file1 two.IStream = new(File)
var file2 one.ReadWriter = file1
var file3 two.IStream = file2
```

接口赋值并不要求两个接口完全等价。
如果接口A的方法列表是接口B方法列表的子集，那么接口B可以赋值给A，反之则不行。
```Go
type Writer interface {
    Write(buf []byte) (n int, err error)
}

var file1 two.IStream = new(File)
var file2 Writer = feil1 // 可以

//反之则不行
var file3 Writer = new(File)
var file4 two.IStream = file3 // 编译不通过，因为file3没有Read方法
```

#### 3.5.4 接口查询
```Go
// 检查file指向的对象实例是否实现了two.IStream接口
// 如何实现了，则执行
if stream, ok := file.(two.IStream); ok {
    ...
}

// 判断file1指向的对象实例是不是*File类型
// 如果是，则执行
if file2, ok := file1.(*File); ok {
    ...
}
```

#### 3.5.5 类型查询
```Go
var v1 interface{} = ...
switch v := v1.(type) {
    case int:  // v 的类型是int
    case string: // v 的类型是string
    ...
}
```

#### 3.5.6 接口组合
可以把接口组合当成是类型匿名组合的一个特定场景。
只不过接口只包含方法，而不包含任何成员变量。
```Go
type ReadWriter interface {
    Reader
    Writer
}

// 完全等同于如下写法
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```
## ch04 并发编程
### 4.5.1 channel的基本语法
声明
`var chanName chan ElementType`
举例
```Go
// 传递类型为int的channel
var ch  chan int

// 声明一个map，元素是bool类型的channel
var m map[string] chan bool

// 声明并初始化一个int类型的名为ch的channel
ch := make(chan int)

// 写入和读取都会导致阻塞
// 写入数据
ch <- value

// 读取数据
value := <- ch
```
### 4.5.2 select
```Go
select {
    case <- ch1:
    // 如果ch1成功读到数据，则执行
    case ch2 <- 2:
    // 如果成功向ch2写入数据，则执行
    default: // 可以没有
    // 如果上面的都没成功，则执行
}
```
### 4.5.3 缓冲机制
```Go
// 即使没有读取方，也可以一直写入
// 在缓冲区被填满之前都不会阻塞
ch := make(chan int, 1024)

//也可以使用range
for i := range ch {
    fmt.Println("Received:", i)
}
```
### 4.5.4 超时机制
用select实现超时机制。select的一个特点就是：只要其中一个case满足，程序就会继续执行，而不会考虑其他的case。
```Go
//首先，实现并执行一个匿名的超时等待函数
timeout := make(chan bool, 1)
go func() {
    time.Sleep(1e9) // 等待1秒
    timeout <- true
}()

select {
    case <- ch:
        // 从ch读取数据
    case <- timeout:
        // 一直没能从ch中读取到数据，但从timeout中读取到了
}
```

### 4.5.6 单向channel
```Go
// 声明
var ch1 chan int // 正常的channel
var ch2 chan<- float64 // 单向的，只能写float64类型的数据
var ch3 <-chan int // 单向的，只用于读取int类型的数据

// 初始化
ch4 := make(chan int)
ch5 := <-chan int(ch4) // 单向的读取
ch6 := chan<- int(ch4) // 单向的写入

//用法
// 只读，类似于java中的final
func Parse(ch <-chan int) {
    for value := range ch {
        fmt.Println("Parsing value", value)
    }
}
```

### 4.5.7 关闭channel
`close(ch)` 内置函数

如何判断一个channel是否已关闭？
`_, ok := <- ch`，ok为false则表示已关闭。

## ch05 网络编程
### 5.1 Socket编程
#### Dial() 函数
```Go
// 定义
// net 是网络协议的名字，tcp 等
// addr 是ip地址或域名
func Dial(net, addr string) (Conn, error)

// 例子
// TCP连接
conn, err := net.Dial("tcp", "192.168.0.10:2100")

// UDP连接
conn, err := net.Dial("udp", "192.168.0.12:975")

// ICMP连接（使用协议名称）
conn, err := net.Dial("ip4:icmp", "www.google.cn")
// 使用协议编号
conn, err := net.Dial("ip4:1", "10.0.0.3")

// 发送数据
conn.Write()
// 接收数据
conn.Read()

// Dial() 是对 DialTCP(),DialUDP(),DialIP(),DialUnix()的封装，也可以直接调用这些函数
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)
...

// 工具函数
net.ParseIP() // 验证IP的有效性
IPv4Mask(a, b, c, d byte) IPMask // 创建子网掩码
(ip IP) DefaultMask() IPMask // 获取默认子网掩码

// 根据域名查找IP
ResolveIPAddr(net, addr string) (*IPAddr, error)
LookupHost(name string) (cnname string, addrs []string, err error)
```
### 5.3 RPC编程
只有满足以下条件的方法才能作为RPC服务端被远程访问
- 公开（首字母大写）
- 必须有两个参数，且参数的类型都是包外可访问的类型或内建的类型
- 第二个参数必须是指针
- 必须返回一个`error`类型的值

如 `func (t *T) MethodName(arg1 T1, reply *T2) error`
第一个参数是由RPC客户端传入的参数
第二个参数表示要返回给客户端的结果