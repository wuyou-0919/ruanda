package main

import "fmt"

func main() {
	var date int
	fmt.Println("今天星期几:")
	fmt.Scanln(&date)
	switch date {
	case 1 :
		fmt.Println("星期一")
	case 2 :
		fmt.Println("星期二")
	case 3 :
		fmt.Println("星期三")
	case 4 :
		fmt.Println("星期四")
	case 5 :
		fmt.Println("星期五")
	case 6 :
		fmt.Println("星期六")
	case 7 :
		fmt.Println("星期日")
	default:
		fmt.Printf("错误信息")
	}
}
