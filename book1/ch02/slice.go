package main

import "fmt"

func main() {
	// 定义一个数组
	var myArr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个数组切片
	var mySlice []int = myArr[:5]

	fmt.Println("Elements of Arr: ")
	for _, v := range myArr {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of Slice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
}
