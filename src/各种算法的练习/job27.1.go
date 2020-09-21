package main

import "fmt"

func main() {
   var a int16 = 26
   var b int16 = 0
   b = a%2 + a/2%2*10 + a/4%2*100 + a/8%2*1000 + a/16%2*10000
	fmt.Printf("%d",b)
}
