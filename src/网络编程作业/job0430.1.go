package main

import (
	"fmt"
	"io"
	"net/http"
)

type num struct {

}

func (a *num) ServeHTTP(writer http.ResponseWriter,Request *http.Request)  {
	path:= Request.URL.Path
	fmt.Println(path)
	if path=="/userinfo"{
		io.WriteString(writer,"查询用户信息")
	}else if path == "/student"{
		io.WriteString(writer,"查询学生信息")
	}
}
func main() {
	a := &num{}
	err:= http.ListenAndServe("127.0.0.110:8090",a)
	if err!=nil {
		fmt.Println(err)
	}
}