package main

import (
	"compress/gzip"
	"os"
	"io"
	"fmt"
)

// 读取压缩文件里的内容
func main() {
	f, err := os.Open("out/words3.txt.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	// io.CopyBuffer(os.Stdout, data)
	fmt.Println(string(data))

}