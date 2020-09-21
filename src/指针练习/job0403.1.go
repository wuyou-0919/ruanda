package main

import "fmt"

func main() {
	nuam :=[5]int{33,44,55,66,77}
	p := &nuam
	fmt.Println("数组的地址:",&nuam)
	fmt.Printf("P变量的类型:%T",p)
}
