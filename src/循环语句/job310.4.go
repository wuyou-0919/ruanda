package main

import "fmt"

func main() {
	var a [10]int
	var i int
	fmt.Println("请输入长度为10的int类型的数组")
	for i=0;i<10;i++{

		fmt.Scanln( &a[i])
	}
	fmt.Println("数组")
	for i=0;i<10;i++ {
		fmt.Println(a[i])
	}


	}