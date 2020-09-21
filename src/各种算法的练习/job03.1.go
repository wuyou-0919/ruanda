package main

import "fmt"

func main() {
	var x string
	var y int
	fmt.Printf("请输入用户名:")
	fmt.Scanf("%s", &x)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%d", &y)
	fmt.Printf("用户名%s 密码%d ",x,y)
}
