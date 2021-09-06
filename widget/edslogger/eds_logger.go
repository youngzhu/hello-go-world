package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "http://eds.newtouch.cn/eds3/login.html"

func main() {
	// 检验网站是否正常
	r, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		// 正常返回还不行，需要有错误发送邮件通知
		return
	}

	// 查看一下
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

}
