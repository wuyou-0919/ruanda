package main

import "fmt"

func main(){
	var arr [20]int = [20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	fmt.Println(arr)
	arr2 :=make(map[int]int)
	var a int=0
	for i:=0;i<len(arr);i++ {
		a = arr[i]
		if arr2[a] != 0 {
			continue
		}
		for j := 0; j < len(arr); j++ {
			if a == arr[j] {
				arr2[a]++
			}
		}
	}
		fmt.Println(arr2)
}