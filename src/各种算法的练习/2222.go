package main

import "fmt"

func main() {
	var a int
	var b float64
	fmt.Println("a的值,b的值")
	fmt.Scanln(&a,&b)
	fmt.Scanf("%d &f\n",&a,&b)
	fmt.Printf("a:%d,b:%f\n",a,b)

}
