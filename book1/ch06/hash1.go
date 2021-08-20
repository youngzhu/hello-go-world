package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestStr := "Hi, golang"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestStr))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("md5: %x\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestStr))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("sha1: %x\n", Result)
}
