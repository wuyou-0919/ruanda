package main

import "fmt"

func main() {
	a := 5
	l := 1
	for i := 1; i <= a; i++ {
		for k := a - i; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= l; j++ {
			fmt.Print("*")
		}
			l += 2
			fmt.Println()



	}
}