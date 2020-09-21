package main

import (
	"fmt"
	"结构体/homework"
)

func main(){
	fmt.Println("作业一")
	ro:=rot{
		r: 5,
		s:  0,
		Long: 0,
	}
	ro.s = ro.S()
	ro.Long = ro.L()
	fmt.Println(ro)
	fmt.Println("作业二")
	sp := homework.Sporter{
		Name:  "王二狗",
		Sex:   "男",
		Sport: "男子蛙泳100米",
		Grade: 100,
	}
	sp.Run()
	fmt.Print("\n==========")
	fmt.Println("作业三")
	car := homework.Car(homework.ElectroCar{
		Bran:"特斯拉",
		Kin: "电动汽车",
	})
	car.Brand()
	car.Kind()
	car = homework.Tricycle{
		Bran: "时风",
		Kin:  "三轮车",
	}
	car.Brand()
	car.Kind()
	car = homework.Truck{
		Bran: "解放牌",
		Kin:  "卡车",
	}
	car.Brand()
	car.Kind()
}

type rot struct {
	r float64
	s float64
	Long float64
}

func (a rot)S() float64 {
	r := a.r
	s:= r*r*3.14
	return s
}

func (a rot)L() float64 {
	r:=a.r
	long:= 2*r*3.14
	return long
}
