package main

import (
	"fmt"
	"math"
)

func main() {
	ending:=count{
		radius:5,
	}
	ending.b()
}

type count struct {
	radius float64
	Perimeter float64
	area  float64
}

func (rad count) b() {
	circumference:=math.Pi*2*rad.radius
	area:=math.Pi*rad.radius*rad.radius
	fmt.Printf("周长:%0.2f,面积：%0.2f",circumference,area)
}