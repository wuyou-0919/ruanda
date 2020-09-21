package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
		d = iota
		e = iota
		)
	fmt.Println("abcde的值分别为:",a,b,c,d,e)
	fmt.Println("a+c的值为:",a+c)


}
