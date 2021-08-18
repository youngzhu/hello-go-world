package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		// 主线程不等待go线程
		// 所以，可能看不到输出，或少量的输出
		go Add(i, i)
	}
}
