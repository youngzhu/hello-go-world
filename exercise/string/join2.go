package main

import (
	"fmt"
	"strings"
)

// 拼数字
func main() {

	s1 := ""
	for i := 0; i < 10; i++ {
		s1 += fmt.Sprintf("%#v", i)
	}
	fmt.Println(s1)

	var b strings.Builder
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&b, "%d", i)
	}
	fmt.Println(b.String())

	var sb strings.Builder
	for i := 0; i < 10; i++ {
		sb.WriteString(fmt.Sprintf("%2d", i))
	}
	fmt.Println(sb.String())
}