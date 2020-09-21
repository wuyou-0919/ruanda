package main

import "fmt"

func main() {
	nuam := []int{1,2,3,4,5}
	fmt.Println("反转前的数组：",nuam)
	nuam = fun1(nuam)
	fmt.Println("反转后的数组：",nuam)
}
func fun1(num []int) []int{
	a := len(num)
	for i:=0;i<a/2;i++{
		temp := num[a-1-i]
		num[a-1-i] = num[i]
		num[i] = temp
	}
	return num
}