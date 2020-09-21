package main

import "fmt"

func main() {
	var date int
	fmt.Println("今天属于该月份的:")
	fmt.Scanln(&date)
	switch date {
	case 1,2,3,4,5,6,7,8,9,10:
		fmt.Println("上旬")
	case 11,12,13,14,15,16,17,18,19,20:
		fmt.Println("中旬")
	case 21,22,23,24,25,26,27,28,29,30:
		fmt.Println("下旬")

	}

}
