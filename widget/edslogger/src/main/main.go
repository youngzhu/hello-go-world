package main

import (
	"log"
	"os"
	"secret"
)

func main() {
	log.Println(os.Getwd()) // 可以查看当前的工作目录
	secret, err := secret.RetrieveSecret()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(secret)
}
