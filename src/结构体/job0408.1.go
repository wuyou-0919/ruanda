package main

import "fmt"

func main() {
	i:=Information{111,"无忧","行政部",3500,0}
	i1:=Information{111,"王二狗","人力资源部",13140,0}
	i2:=Information{112,"狗蛋","生产技术部",5200,0}
	i3:=Information{113,"小花","财务部",6000,0}
	fmt.Println(i,i1,i2,i3)
}

type Information struct {
	serialnumber int
	name string
	department string
	salary int
	pay int

}

