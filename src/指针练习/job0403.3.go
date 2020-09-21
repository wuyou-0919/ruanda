package main

import "fmt"

func main() {
	nu1,nu2 :=19,20
	deng :=jia(nu1,nu2)
	fmt.Println("两数相加等于:",*deng)

}
func jia(nu1 int,nu2 int)*int{
	nu3 :=nu1+nu2
	return &nu3

}

