package main

import "fmt"

func main() {
	for i:=1;i<=5;i++{
		for j:=0;j<5 - i;j++{
			fmt.Print(" ")
		}
		for k:=0;k<2*i-1;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=1;i<5;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for k:=0;k< 2*5-1 - 2*i;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
