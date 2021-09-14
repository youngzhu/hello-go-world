package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// log.Println(os.Getwd()) // 可以查看当前的工作目录
	// secret, err := secret.RetrieveSecret()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(secret)

	// 获取当前时间
	t := time.Now()
	log.Println(t)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		t = t.Add(time.Hour * 24)
		// log.Println(fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day()))
		// "2006-01-02 15:04:05"
		// 这算小彩蛋吗？还必须这个时间才行。。。
		log.Println(t.Format("2006-01-02"))
	}

}
