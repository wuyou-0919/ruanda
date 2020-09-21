package main

import "fmt"

func main() {
	a1:=Truck{name:"解放者一号"}
	a2:=electric{name:"特斯拉电动车"}
	a3:=Tricycle{name:"时风三轮车"}
	a1.drive()
	a2.drive()
	a3.drive()
}

type Vehicle interface{
	brand()    string
	drive()    string
}
type Truck struct {
	name string
}

func (a Truck)brand() string  {
	return a.name
}
func (a Truck) drive() {
	fmt.Println("解放者一号")
}
type electric struct {
	name string
}
func (b electric)brand()string  {
	return b.name
}
func (b electric) drive() {
	fmt.Println("特斯拉电动车")
}
type Tricycle struct {
	name string
}
func (c Tricycle)brand()string  {
	return c.name
}
func (c Tricycle) drive() {
	fmt.Println("时风三轮车")
}