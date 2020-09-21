package main

import "fmt"

func main() {
    wuyou := "wuhanjiayou，zhongguojiayou"
    wuhan := wuyou[0:5]
    wuhan1 :=wuyou[14:22]
    wuhan2 :=wuyou[0:11]
    wuhan3 :=wuyou[14:28]
    fmt.Println("字符串wuyou的长度为：",len(wuyou))
    fmt.Println(wuhan)
    fmt.Println(wuhan1)
    fmt.Println(wuhan2)
    fmt.Println(wuhan3)



}
