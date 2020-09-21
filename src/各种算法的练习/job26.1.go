package main

import "fmt"

func main() {
	a := 20
	b := 3
	sum1 := a+b
    sum2 := a-b
    sum3 := a*b
    sum4 := a/b
    sum5 := a%b
    fmt.Printf("\n %d + %d=%d",a,b,sum1)
	fmt.Printf("\n %d - %d=%d",a,b,sum2)
	fmt.Printf("\n %d * %d=%d",a,b,sum3)
	fmt.Printf("\n %d / %d=%d",a,b,sum4)
    fmt.Printf("\n %d %% %d=%d",a,b,sum5)
}