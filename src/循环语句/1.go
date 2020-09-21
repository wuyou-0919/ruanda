package main

import (
	"fmt"
	"math"
)
func main() {
	for i := 100;i<1000;i++{
		a :=i /100
		b :=i / 10 %10
		c :=i % 10
		if math.Pow(float64(a),3)+math.Pow(float64(b),3)+math.Pow(float64(c),3)==float64(i){
			fmt.Println(i)
		}
	}
}
