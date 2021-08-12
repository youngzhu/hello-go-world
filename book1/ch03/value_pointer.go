package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

// 指针传递，会改变a的值
func (a *Integer) Add1(b Integer) {
	*a += b
}

// 值传递，a值不变
func (a Integer) Add2(b Integer) {
	a += b
}

func main() {
	var a1 Integer = 1
	a1.Add1(2)
	fmt.Println("a1=", a1)

	var a2 Integer = 1
	a2.Add2(2)
	fmt.Println("a2=", a2)
}
