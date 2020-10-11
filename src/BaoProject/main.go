package main

import (
	"BaoProject/db_mysql"
	_ "BaoProject/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

