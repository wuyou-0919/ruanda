package main

import "fmt"

func main() {
	a :=0
	for i :=0;i <10;i++{
		a++
		for j :=0;j <a;j++{
			if j==0 || j== a-1 || i==9{
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}