package main

import (
	"test.com.zhj/myblog/utils"

	"github.com/beego/beego"
	_ "test.com.zhj/myblog/routers"
)

func main() {

	utils.InitMysql()
	beego.Run()
}
