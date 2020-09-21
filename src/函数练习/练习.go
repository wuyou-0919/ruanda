package main

import "fmt"

func main() {
	n:=9
	sun(n,n-1,1)

}
func sun(n int,m int,k int){
	for i:=1;i<=k;i++{
		fmt.Print(i)
	}
	fmt.Printf(" ")
	if m>0{
		k++
		sun(n,m-1,k)
	}
}
