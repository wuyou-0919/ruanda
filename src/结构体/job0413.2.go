package main

import "结构体/home"

func main()  {
	results :=home.Athletes{
		Name:"孙杨",
		Gender: "男",
		SportEvent:"游泳",
		Achi: 1500,
	}
	results.Run()
}