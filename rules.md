# Go语言编程规范

- 每个源码文件都有一个包（package），包名与所在文件夹名称一致。不能把多个包放到同一个目录中，也不能把同一个包拆分到不同的目录中。
- 包名全部小写。
- 文件名小写，以`_`间隔
- `example_test.go`中的 `import` 不要用 `.`，因为在实际使用中也是`package.Func`。`xxx_test.go`中可以使用，因为它是包内的测试类
- 要习惯用工厂方法`package.NewXXX()`。`package.XXX{}`，不合适也不好看
- 常量不用全大写，也不用下划线，`const MaxSize = 100`就可以，小写也行，源码就这么写的`const startBufSize = 4096`
- 