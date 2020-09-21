package main

import "fmt"

func main() {
	 a := [6]int{5,9,20,120,450,340}
	fmt.Println("排序前：", a)
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] > a[min] {
				min = j
			}
		}
		if min != i {
			temp := a[i]
			a[i] = a[min]
			a[min] = temp
		}
	}
	fmt.Println("排序后：",a)
}
