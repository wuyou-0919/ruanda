package main

import "fmt"

func main() {
	var wuyou int
	fmt.Println("今天星期几")
	fmt.Println("1表示星期一，2表示星期二，3表示星期三，4表示星期四，5表示星期五，6表示星期六，7表示星期日，其他数字则是错误信息")
	fmt.Scanln(&wuyou)
	if wuyou == 1{
		fmt.Println("星期一")
	}else if wuyou ==2 {
		fmt.Println("星期二")
	}else if wuyou == 3 {
		fmt.Println("星期三")
	}else if wuyou ==4 {
		fmt.Println("星期四")
	}else if wuyou == 5 {
		fmt.Println("星期五")
	}else if wuyou ==5 {
		fmt.Println("星期六")
	}else if wuyou == 7 {
		fmt.Println("星期天")
	}else{
		fmt.Println("错误信息")
	}
}