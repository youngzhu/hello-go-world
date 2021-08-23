package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
列出指定目录下文件名中包含指定关键字（不区分大小写）的文件名
*/

var keyword *string = flag.String("k", "keyword", "The keyword to match")

func main() {
	flag.Parse()

	log.Println("keyword:", *keyword)

	// 都转为大写进行比较
	var upper1, upper2 string

	upper1 = strings.ToUpper(*keyword)

	books, err := ioutil.ReadDir("D:/GitHub/ituring_books/图灵程序设计丛书")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, book := range books {
		bookName := book.Name()
		upper2 = strings.ToUpper(bookName)

		if strings.Contains(upper2, upper1) {
			log.Println(bookName)
		}

	}

}
