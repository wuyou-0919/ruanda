package main

import "fmt"

func main() {
	nu :=[6]string{"wuyou","qzl","quan","zhi","long","vip"}
	var nua [6]*string
	for i :=0;i <len(nu);i++{
		nua[i]=&nu[i]
	}
	fmt.Println("变量地址:",nua)
	*nua[5]="VIP"
	fmt.Println("修改原字符串数组中最后一个元素的元素值为:",nu)
}
