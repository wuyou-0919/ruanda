package main

import "fmt"

func main() {
   n :=10
   for i :=1;i <=n;i++{
   	for k:=1;k<=n-1;k++{
   		fmt.Print(" ")
	}
	fmt.Print("*")
   	if i!=n{
   		for j:=1;j<= i-2;j++{
   			fmt.Print(" ")
		}
		if i !=1{
			fmt.Print("*")
		}
		fmt.Println()
	}else {
		for m :=1;m<=i-1;m++{
			fmt.Print("*")
		}
	}
   }
}
