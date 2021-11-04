package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3}
	slice := []int{11, 22, 33}

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)

	funcArray(array)
	funcSlice(slice)

	fmt.Println("after...")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice)

	// Output:
	// array: [1 2 3]
	// slice: [11 22 33]
	// after...
	// array: [1 2 3]
	// slice: [-11 22 33]
}

func funcArray(a [3]int) {
	a[0] = -1
}

func funcSlice(s []int) {
	s[0] = -11
}