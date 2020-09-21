package main

import "fmt"

func main() {
	a :=make([]int,5,7)
	fmt.Printf("长度：%d 容量：%d\n",len(a),cap(a))
	a[0]=1
	a[1]=2
	a[2]=3
	a[3]=4
	a[4]=5
	fmt.Println("赋值后：",a)

}