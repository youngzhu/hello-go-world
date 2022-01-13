package main

import (
	"regexp"
	// "strings"
	"fmt"
)


func main() {
	str := "http://www.Youngzy.com https://go.org"
	re := regexp.MustCompile(`(http|https)://(\w+\.)+(edu|com|org|me)`)
	links := re.FindAllString(str, -1)

	fmt.Println(links)
}