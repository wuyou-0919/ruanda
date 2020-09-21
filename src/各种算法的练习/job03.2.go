package main

import (
	"fmt"
)

func main() {
	var a float64
	var b int
	fmt.Printf("你的身高是/m：")
	fmt.Scanf("%f", &a)
	fmt.Printf("你的体重是/kg：")
	fmt.Scanf("%d", &b)
	fmt.Printf(" 身高%.3fm 体重%dkg",a,b)

}
