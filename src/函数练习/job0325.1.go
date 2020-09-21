package main

import "fmt"

func main() {
	//a := 5
	//b := 3
	var a int
	var b int
	fmt.Println("请输入矩形的长")
	fmt.Scanln(&a)
	fmt.Println("请输入矩形的宽")
	fmt.Scanln(&b)
	c := num(a,b)
	fmt.Println("矩形的周长为：",c)
}
func num(num1 int,num2 int) int{
	c := (num1 + num2)*2
	return c
}
