package main

import "fmt"

func main() {
	i:=Informatio{111,"无忧","行政部",3500,0}
	i1:=Informatio{111,"王二狗","人力资源部",13140,0}
	i2:=Informatio{112,"狗蛋","生产技术部",5200,0}
	i3:=Informatio{113,"小花","财务部",6000,0}
	fmt.Println(i,i1,i2,i3)
	a1:= two(i2)
	fmt.Println(a1)
}
func two (a Informatio)Informatio {
	num:=0
	if a.salary<=5000 {
		num=0
	}else if a.salary>5000 {
		num1 := (a.salary-5000)*1/5
		num=num1
	}
	a.pay=num
	return a
}

type Informatio struct {
	serialnumber int
	name string
	department string
	salary int
	pay int

}

