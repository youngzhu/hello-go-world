package main

import "fmt"

func main() {
	s := "aAbc d"
	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Println("=================")
	
	for v := range s {
		fmt.Println(v)
	}
}