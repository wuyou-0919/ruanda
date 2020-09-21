package main

import "fmt"

func main() {
	const (
		wuyou1 string = "浮点型变量"
		wuyoua float32 = 9.9
		wuyou2 string = "整型变量"
		wuyoub int = 1
		wuyou3 string = "字符型变量"
		wuyouc string = "无忧"
		wuyou4 string = "布尔类型"
		wuyou bool = true
		qzl bool = false


	)
	fmt.Println(wuyou1,wuyoua,wuyou2,wuyoub,wuyou3,wuyouc,wuyou4,wuyou,qzl)
	fmt.Printf("%s: %T %s: %T %s: %T %s: %T",wuyou1,wuyoua,wuyou2,wuyoub,wuyou3,wuyouc,wuyou4,wuyou)

}
