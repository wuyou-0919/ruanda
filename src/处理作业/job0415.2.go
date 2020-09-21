package main

import "fmt"

func main() {
	a := []int{23,34,45,56,67,78}
	for i:=0;i<10;i++{
		if i>=len(a){
			panic("超出数组长度")
		}else if i<=5 {
			fmt.Println(a[i])
		}
	}
}
