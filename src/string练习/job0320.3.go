package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "WUHANJIAYOUZHONGGUOJIAYOU"
	str1 := strings.Contains(str,"ZHONGGUO")
	if str1 {
		fmt.Println("存在,则位于:", strings.Index(str, "ZHONGGUO"))
	}else{
		fmt.Println("不存在ZHONGGUO",str)
	}
}