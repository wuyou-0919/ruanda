package main

import "fmt"

func main() {
	for i:=0;i<6;i++{
		for k:=0;k<6-i;k++ {
			fmt.Print(" ")
		}
		for j:=0;j<9;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}