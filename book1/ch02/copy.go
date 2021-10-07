package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	fmt.Print("slice1: ")
	for _, v := range slice1 {
		fmt.Print(v, ", ")
	}
	fmt.Print("\nslice2: ")
	for _, v := range slice2 {
		fmt.Print(v, ", ")
	}

	// copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	// 2:3 都是下标（不是长度），从0开始
	copy(slice1[2:3], slice2) // 部分复制

	fmt.Print("\nafter copy \nslice1: ")
	for _, v := range slice1 {
		fmt.Print(v, ", ")
	}
	fmt.Print("\nslice2: ")
	for _, v := range slice2 {
		fmt.Print(v, ", ")
	}
}
