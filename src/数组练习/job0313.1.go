package main

import "fmt"

func main() {
	i := [6]int{5,9,20,120,450,340}
	fmt.Println("排序前：",i)
	for j := 1; j < len(i); j++{
		for k := 0; k < len(i) - j ;k++ {
			if i[k] < i[k+1]{
				i[k], i[k+1] = i[k+1], i[k]
			}
		}
	}

	fmt.Println("排序后：",i)
}
