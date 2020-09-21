package main

import "fmt"

func main() {
	i :="我爱你中国，I LOVE CHINA"
	fmt.Println("该字符串的长度:",len(i))
	j := []rune(i)
	for k :=0;k<len(j);k++{
		fmt.Printf("%c\t",j[k])
	}
}
