package main

import "fmt"

func main() {
	var a1 = [3]int{1, 2, 3}
	var b1 = a1 // 数组内容的完整复制
	b1[1]++
	fmt.Println(a1, b1) //[1 2 3] [1 3 3]

	var a2 = [3]int{1, 2, 3}
	// & 表示数组内容的引用
	// 变量b的类型不是[3]int，而是*[3]int
	var b2 = &a2
	b2[1]++
	fmt.Println(a2, b2)  //[1 3 3] &[1 3 3]
	fmt.Println(a2, *b2) //[1 3 3] [1 3 3]
}
