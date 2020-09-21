package main

import (
	"fmt"
	"math"
	"net/http"
)
func personal(q []string) int{
	var nu = len(q[0])
	var nua int = 0
	for i :=0;i < nu;i++{
		nua = (int(q[0][i])-48)*int(math.Pow10(int(nu-i-1))) + nua
	}
	return nua
}
func per(w http.ResponseWriter,r *http.Request){
	er := r.URL
	r.ParseForm()
	date :=r.Form
	if len(date)==0{
		return
	}
	fmt.Printf("path:%s\n",er.Path)
	fmt.Printf("scheme:%s\n",er.Scheme)
	fmt.Printf("from:%s\n",r.Form)
	year := personal(date["year"])
	month := personal(date["month"])
	day := personal(date["day"])
	if len(date["username"][0])>12{
		w.Write([]byte("姓名不符合规范"))
	}else if year>2020||(year==2020 && month > 5)||(year==2020 && month==5 && day > 9){
		w.Write([]byte("太着急了，还没到日子呢。"))
	}else if len(date["ID"][0])!=18{
		w.Write([]byte("身份证长度不符合规范"))
	}else{
		w.Write([]byte("完成注册!"))
	}
	fmt.Printf("姓名: %s\n",date["username"])
	fmt.Printf("出生日期: %d年%d月%d日",year,month,day)
}
func main(){
	fmt.Printf("开始运行")
	http.HandleFunc("/",per)
	err :=http.ListenAndServe("127.0.0.110:8080",nil)
	if err!=nil{
		fmt.Printf("ListenAndServe:%s",err)
		return
	}
	return
}
