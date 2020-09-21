package main

import "fmt"

func main() {
	a :=[]byte{97,98,99,100}
	//a1 :=string(a)
	//fmt.Println(a1)
	for j :=0;j <len(a);j++{
		fmt.Printf("%c\t",a[j])
	}
}
