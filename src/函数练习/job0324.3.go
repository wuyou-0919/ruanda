package main

import "fmt"

func main() {
	var a int
	for i:= 0;i<2;i++{
		fmt.Scanln(&a)
		if year(a){
			fmt.Println("该年是闰年")
		}else{
			fmt.Println("该年不是闰年")
		}
	}
}

func year(a int)(bool,){
	if ((a%100!=0)&&(a%4==0))||a%400==0{
		return true
	}else{
		return false
	}


}
