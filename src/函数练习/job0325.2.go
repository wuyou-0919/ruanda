package main

import "fmt"

func main() {
	num := "WUHANJIAYOUZHONGGUOJIAYOU"
	num1(num)
}
func num1(num2 string){
	fmt.Println(num2)
	num3 :=make(map[string]int)
	var a string = ""
	for i:=0;i<len(num2);i++{
		a = num2[i:i+1]
		if num3[a]!=0 {
			continue
		}
		for j:=0;j<len(num2);j++{
			if a == num2[j:j+1]{
				num3[a]++
			}
		}
	}
	fmt.Println(num3)
}
