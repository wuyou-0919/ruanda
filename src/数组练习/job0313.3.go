package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [10]int
	num :=0
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		a[i] = rand.Intn(15)
		num++
		for j:=0;j<i;j++{
			if a[i]==a[j]{
				i--
			}
		}
	}
	fmt.Println("输出：",a)
	fmt.Println("一共调用了",num)

}
