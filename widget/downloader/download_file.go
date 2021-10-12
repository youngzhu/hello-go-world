package main

import (
	"net/http"
	"io/ioutil"
)

var url = "https://algs4.cs.princeton.edu/21elementary/words3.txt"

// 从指定url下载文件
// 如果文件大于1000行，则压缩
func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/words3.txt", body, 0777)
	if err != nil {
		panic(err)
	}
}