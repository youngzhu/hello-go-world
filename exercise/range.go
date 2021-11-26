package main

import "fmt"

func main() {
	s := "aAbc d"
	for _, v := range s {
		fmt.Println(v)
	}
}