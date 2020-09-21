package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "root:123456@tcp(127.0.0.1:3306)/mydate?charset=utf8"
	sql, err :=sql.Open("mysql",connStr)
	fmt.Printf("%T",sql)
	if err != nil{
		fmt.Println("数据库连接失败：",err.Error())
		return
	}
	defer sql.Close()
}