# Troubleshooting

## Q0: VSCode（因为墙）无法下载必要的Module？
设置代理，命令窗口下执行：
```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB="sum.golang.google.cn" #不确定这个是否必要，我都改了
```

可通过`go env`查看当前的值。

## Q1: VSCode中如何运行Go程序？
Extensions，搜索并安装 **Code Runner** 插件。
有时候会失败，如`os.Open()`无法读取其他文件。这时还要用原始的命令行。

## Q2: 错误提示：... cannot find main module ...
在命令窗口执行如下命令
```
go env -w GO111MODULE=off
```