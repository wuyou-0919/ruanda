package main

import (
	"fmt"
	"strings"
)

func main() {
	job := "http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
	fmt.Println("文件路劲:",job)
	job1 := strings.LastIndex(job,"/")
	str := job[job1+1:]
	job1 = strings.LastIndex(str,".")
	str1 := str[job1+1:]
	str = str[:job1]
	fmt.Println("文件的名称：",str)
	fmt.Println("文件的扩展名类型：",str1)
}
