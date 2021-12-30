package main

import (
	"fmt"
	"strings"
)

func main() {
	base := "TeST"

	s1 := ""
	for i := 0; i < 10; i++ {
		s1 += base
	}
	fmt.Println(s1)

	var b strings.Builder
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&b, "%s", base)
	}
	fmt.Println(b.String())

	var sb strings.Builder
	for i := 0; i < 10; i++ {
		sb.WriteString(base)
	}
	fmt.Println(sb.String())
}