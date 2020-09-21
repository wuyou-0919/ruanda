package main

import "fmt"

func main() {
	a:=9
	num (a,a-1,1)
}
func num(a int,b int,c int){
	for i:=1;i<=c;i++{
		fmt.Print(i)
	}
	fmt.Print(" ")
	if b>0{
		c++
		num(a,b-1,c)
	}
}