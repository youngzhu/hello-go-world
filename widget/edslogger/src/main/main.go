package main

import (
	"log"
	"time"

	"logger"
)

func main() {
	err := logger.Login()
	if err != nil {
		// 正常返回还不行，需要有错误发送邮件通知
		// return
		log.Fatalln("网站服务错误", err)
		// os.Exit(1)
	}
	log.Println("登陆成功")

	logFrom, _ := time.Parse("2006-01-02", "2021-09-20")
	log.Println(logFrom)

	// logFromSpecificDay(logFrom)

}