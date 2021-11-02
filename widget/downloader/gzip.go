package main

import (
	"os"
	"io"
	"strings"
	"compress/gzip"
)

const path = "out/in.txt"

// 压缩本地文件
func main() {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	gzipCompress(getFileName(), data)
}

// 从url中获取文件名
func getFileName() string {
	s := strings.Split(path, "/")
	len := len(s)
	return s[len-1]
}

// gzip 压缩
func gzipCompress(fileName string, data []byte) {
	fw, err := os.Create("out/"+fileName+".gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	zw := gzip.NewWriter(fw)
	defer zw.Close()

	_, err = zw.Write(data)
	if err != nil {
		panic(err)
	}

}