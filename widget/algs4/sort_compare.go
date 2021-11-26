package main

import "fmt"
var newAlg string = "Heap"
var otherAlgs []string

func init() {
	otherAlgs = append(otherAlgs, "Selection")
	otherAlgs = append(otherAlgs, "Insertion")
	otherAlgs = append(otherAlgs, "Shell")
	otherAlgs = append(otherAlgs, "Merge")
	otherAlgs = append(otherAlgs, "Quick")
	// append(otherAlgs, "")
	// append(otherAlgs, "")
	// append(otherAlgs, "")
	// append(otherAlgs, "")


	otherAlgs = append(otherAlgs, "Buildin")
}

func main() {
	for _, v := range otherAlgs {
		fmt.Println("//", newAlg, "vs", v)
		fmt.Printf("// cmd: go run sort_compare.go -a1 %s -a2 %s -n 1000 -t 100 -s\n", newAlg, v)
		fmt.Println("// got: ")
		fmt.Printf("// cmd: go run sort_compare.go -a1 %s -a2 %s -n 1000 -t 100\n", newAlg, v)
		fmt.Println("// got: ")
		fmt.Println("")
	}
}