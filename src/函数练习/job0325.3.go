package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("圆的半径是：%0.0f \n圆面积为：%0.2f",5.0,fun(5.0))
}
func fun(a float64) float64{
	num := a *a *math.Pi
	return num
}