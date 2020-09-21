package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	when := http.DefaultClient
	wh, en := when.Get("http://127.0.0.110:9001/userlogin")
	if en != nil {
		fmt.Printf("请求失败 %s\n", en)
		return
	}
	w, h := ioutil.ReadAll(wh.Body)
	if h != nil {
		fmt.Printf("读取失败 %s\n", h)
		return
	}
	wh.Body.Close()
	fmt.Printf("程序运行效果:%s", w)
}