package main

import "fmt"

func main() {
	var num string = "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	fmt.Println(num)
	num2 :=make(map[string]int)
	var a string = ""
	for i:=0;i<len(num);i++{
		a = num[i:i+1]
		if num2[a]!=0 {
			continue
		}
		for j:=0;j<len(num);j++{
			if a == num[j:j+1]{
				num2[a]++
			}
		}
	}
	fmt.Println(num2)
}
