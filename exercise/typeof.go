package main

import (
	"fmt"
	"reflect"
)

type MyType []int

var v1 MyType = MyType{}
var v2 = 1

func main() {
	// .(type) 不好用，需要 iterface 类型
	// switch t := v1.(type) {
	// default:
	// 	fmt.Println(t)
	// }
	
	fmt.Println(reflect.TypeOf(v1)) // main.MyType
	fmt.Println(reflect.TypeOf(v2)) // int

}