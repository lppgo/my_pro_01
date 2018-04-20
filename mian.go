package main

import (
	"log"

	"git.yeeuu.com/yeeuu/mypro_requestURL/controller"
	"git.yeeuu.com/yeeuu/mypro_requestURL/db"
)

func main() {
	// 初始化数据库连接
	if e := db.InitDB(); e != nil {
		log.Fatalf("InitDB: %v", e)
		panic(e.Error())
	}

	//业务处理
	controller.Getdata()
}
