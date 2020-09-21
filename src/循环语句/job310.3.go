package main

import "fmt"

func main() {
		var a = [5]float64{8,18, 9.19, 10.21,66.66}
		sum := float64(0)
		for _, v := range a {
			sum += v
			}
		fmt.Printf("数组的各元素的和:%v",sum)

}
