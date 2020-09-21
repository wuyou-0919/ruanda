package main

import (
	"fmt"
	"strings"
)

func main() {
	a :="hello hello hello word"
	b := strings.Count(a,"llo")
	fmt.Printf("\n%s \n该字符串\"llo\"出现的次数为:%d",a,b)
}
