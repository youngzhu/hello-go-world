package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	// 定义一个channel数组
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		// 将10个channel分配给10个goroutine
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	fmt.Println("Done.")
}
