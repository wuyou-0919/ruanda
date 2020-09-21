package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("开始监听，请注意规范")
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		id := request.FormValue("id")
		fmt.Println(id)
		password :=request.Form.Get("password")
		fmt.Println(password)
		if id == "hello" && password == "123456" {
			writer.Write([]byte("恭喜登录成功"))
		}else{
			writer.Write([]byte("对不起，您的账号或者密码不正确，请重新尝试"))
		}
	})
	err:=http.ListenAndServe("127.0.0.11:8080",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
}