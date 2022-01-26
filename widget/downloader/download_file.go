package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const MAX_BYTES int = 10000 // 文件是否压缩的阈值

// var url = "https://algs4.cs.princeton.edu/13stacks/tobe.txt"
// var url = "https://algs4.cs.princeton.edu/14analysis/1Kints.txt"
// var url = "https://algs4.cs.princeton.edu/14analysis/2Kints.txt"
// var url = "https://algs4.cs.princeton.edu/14analysis/4Kints.txt"
// var url = "https://algs4.cs.princeton.edu/15uf/mediumUF.txt"

// var url = "https://algs4.cs.princeton.edu/21elementary/words3.txt"
// var url = "https://algs4.cs.princeton.edu/24pq/tiny.txt"
// var url = "https://algs4.cs.princeton.edu/24pq/m1.txt"
// var url = "https://algs4.cs.princeton.edu/24pq/m2.txt"
// var url = "https://algs4.cs.princeton.edu/24pq/m3.txt"
// var url = "https://algs4.cs.princeton.edu/25applications/domains.txt"
//var url = "https://algs4.cs.princeton.edu/25applications/california-gov.txt"

// var url = "https://algs4.cs.princeton.edu/31elementary/tinyTale.txt"
// var url = "https://algs4.cs.princeton.edu/31elementary/tale.txt"
// var url = "https://algs4.cs.princeton.edu/31elementary/leipzig100K.txt"

// var url = "https://algs4.cs.princeton.edu/41graph/tinyG.txt"
// var url = "https://algs4.cs.princeton.edu/41graph/tinyCG.txt"
// var url = "https://algs4.cs.princeton.edu/41graph/routes.txt"
// var url = "https://algs4.cs.princeton.edu/41graph/movies.txt"
// var url = "https://algs4.cs.princeton.edu/42digraph/tinyDG.txt"
// var url = "https://algs4.cs.princeton.edu/42digraph/tinyDAG.txt"
// var url = "https://algs4.cs.princeton.edu/42digraph/mediumDG.txt"
// var url = "https://algs4.cs.princeton.edu/43mst/tinyEWG.txt"
// var url = "https://algs4.cs.princeton.edu/43mst/mediumEWG.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/tinyEWD.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/tinyEWDAG.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/jobsPC.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/tinyEWDn.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/tinyEWDnc.txt"
// var url = "https://algs4.cs.princeton.edu/44sp/rates.txt"

//var url = "https://algs4.cs.princeton.edu/63suffix/abra.txt"
//var url = "https://algs4.cs.princeton.edu/63suffix/tinyTale.txt"
var url = "https://algs4.cs.princeton.edu/63suffix/mobydick.txt"

// 从指定url下载文件
// 如果文件大于设定值，则压缩
func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fileName := getFileName()
	fmt.Println(fileName)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if len(body) > MAX_BYTES {
		// 压缩
		gzipCompress(fileName, body)

	} else {
		err = ioutil.WriteFile("out/"+fileName, body, 0777)
		if err != nil {
			panic(err)
		}
	}

}

// 从url中获取文件名
func getFileName() string {
	s := strings.Split(url, "/")
	len := len(s)
	return s[len-1]
}

// gzip 压缩
func gzipCompress(fileName string, data []byte) {
	fw, err := os.Create("out/" + fileName + ".gz")
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
