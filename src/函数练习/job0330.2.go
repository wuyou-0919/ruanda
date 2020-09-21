package main

import "fmt"

func main() {
	func (num int){
		num1:=1
		for i:=1;i<=num;i++{
			num1*=i
		}
		fmt.Printf("%d的阶乘是:%d",num,num1)
	}(6)
}


