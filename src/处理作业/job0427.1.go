package main

import (
	"fmt"
	"math/rand"
	"time"
)
var ch1 chan int
var ch2 chan int
var num bool=false
var num1 bool=false

func main() {
	ch1 =make(chan int,5)
	ch2 =make(chan int,5)
	nu := 0
	go call()
	go call1()
	for {
		nu =<- ch2
		fmt.Printf("该随机数的三次方为：%d\n",nu)
		if num && num1{
			break
		}
	}
}
func call()  {
	rand.Seed(time.Now().Unix())
	var one int
	for i:= 0;i< 100;i++{
		one=rand.Intn(1000)
		time.Sleep(20*time.Microsecond)
		ch1 <- one

	}
	num1 = true
}
func call1()  {
	one:= 0
	for {
		one =<-ch1
		fmt.Printf("随机数的数值为：%d",one)
		ch2 <-one*one*one
		if len(ch1)== 0 && num1{
			break
		}
	}
	num =true
}