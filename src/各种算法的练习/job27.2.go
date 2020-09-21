package main

import "fmt"

func main() {
	wu := 31
	you := 0
	you = wu%8 +wu/8%7*10
	fmt.Printf("%d",you)

}