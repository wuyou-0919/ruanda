package main

import "fmt"

func main() {
	var a int16 = 4
	var b int16 = 3
	res1 := a < b && b / 2 == 1 && a % 3 != 0
	res2 := (a+b)*3 < a<<2 || (a-b) >0
	fmt.Printf("res1表达式的值为：%t \n",res1)
	fmt.Printf("res2的表达式的值为：%t \n ",res2)
}
