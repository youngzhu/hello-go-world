package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "www.youngzy.com"

	get(url)

}

// 请求一个资源
func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
