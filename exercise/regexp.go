package main

import (
	"strings"
	"fmt"
)
// 匹配所有的单词分隔符
func split(s string) {
	// pattern := regexp.MustCompile(`a`)
	// fmt.Println(pattern.Split(s, -1))

	// 本以为很复杂的功能，被一个预设函数解决了。。
	fmt.Println(strings.Fields(s))
}

func main() {
	split("xx yy   xx z\n aaaa ")
}