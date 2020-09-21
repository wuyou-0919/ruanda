package main

import "fmt"

func main() {
	var a int
	var b int
	var c string
	fmt.Println("输入一个整数")
	fmt.Scanln(&a)
	fmt.Println("再输入一个整数")
	fmt.Scanln(&b)
	fmt.Println("输入执行的操作符号(++,--,%):")

	fmt.Scanln(&c)

	switch c {
	case "++":
		fmt.Println("两数自增的结果是:",a+1,b+1)
	case "--":
		fmt.Println("两数自减的结果是:",a-1,b-1)
	case "%":
		fmt.Println("两数求余的结果是:",a%b)

	}
}
