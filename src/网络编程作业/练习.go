package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}

func (mux *MyMux) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	path := request.URL.Path
	fmt.Println(path)

	if path =="/login"{
		login1(writer,request)
		return
	}else if path == "/register" {
		register(writer,request)
		return
	}
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe("127.0.0.1:8080",mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func register(writer http.ResponseWriter,request *http.Request){
	//注册逻辑
	writer.Write([]byte("注册成功"))
}


func login1(writer http.ResponseWriter,request *http.Request){
	//用户名和密码：服务器和客户端约定一个固定名字，比如说username, password
	username := request.Form.Get("username")//字段名为username的请求数据
	password := request.Form.Get("password")//字段名为password的请求数据
	if username == "hello" && password == "123456" {
		writer.Write([]byte("恭喜登录成功"))
	}else {
		writer.Write([]byte("对不起，您的账号或者密码不正确，请重新尝试"))
	}
}