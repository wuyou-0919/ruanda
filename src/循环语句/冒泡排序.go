package main

import "fmt"

func main() {
	var a [10]int
	var j,i,t int
	fmt.Println("请输入一维数组的值")
	for i=0;i<=9;i++{
	fmt.Scanln( &a[i])
	}
	fmt.Println("一组数组:")
	for i=1;i<=9;i++{
	for j=0;j<=9-i;j++{
	if a[j]>a[j+1]{
	t=a[j];
	a[j]=a[j+1];
	a[j+1]=t;
	}
	}
	}
	for i=0;i<=9;i++{
	fmt.Println(a[i])
	}
}
