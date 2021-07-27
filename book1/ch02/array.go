package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(array)                  // 试图修改这个数组
	fmt.Println("In main(), array values:", array)
}
