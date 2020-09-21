package main

import "fmt"

func main() {
	get := []int{33,45,76,98,11}
	fmt.Println("排序前：",get)
	get1:= []*int{&get[0],&get[1],&get[2],&get[3],&get[4]}
	pa1 :=getnum(get1)
	fmt.Println("排序后：",*pa1)
}
func getnum(pa []*int) *[]int{
	a := len(pa)
	b := make([]int,a,a)
	for i := 0;i < a; i++{
		for j := i; j < a; j++ {
			if *pa[j] < *pa[i] {
				*pa[i],*pa[j] =*pa[j],*pa[i]
			}
		}
		b[i] = *pa[i]
	}
	return &b
}