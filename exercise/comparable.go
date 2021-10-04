package main

import (
	"fmt"
)

type Comparable interface {
	LessThan(a Comparable) bool
}

type ComparableInt int

func (c ComparableInt) LessThan(another ComparableInt) bool {
	return false
}

func main() {
	c1 := ComparableInt(1)
	c2 := ComparableInt(2)
	fmt.Println(c1.LessThan(c2))

	// s1 := []ComparableInt{c1, c2}
	// var s2 []Comparable
	// s2 = s1
	// fmt.Println(s2[0].LessThan(s2[1]))

	Compare(c2)
}

// 不存在这样的继承方式
// 只到对象，不到数组或切片
// func Compare(a []Comparable) {
// 	fmt.Println(a[0].LessThan(a[1]))
// }

func Compare(a Comparable) {
	fmt.Println(a)
}