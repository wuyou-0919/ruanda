package home

import "fmt"

type Athletes struct {
	Name string
	Gender string
	SportEvent string
	Achi float64
}
func(a Athletes)Run(){
	fmt.Printf("%s的运动项目是%s",a.Name,a.SportEvent)
}