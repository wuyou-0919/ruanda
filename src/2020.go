package main

import (
	"encoding/hex"
	"fmt"
)

func main()  {
	arr := []byte{'1','0','0','0','a','b','c','d','e'}
	fmt.Println(string(arr))
	str := BytesToHexString(arr)
	fmt.Println(str)
	str = ReverseHexString(str)
	arr,_ = HexStringToBytes(str)
	fmt.Printf("%x \n",arr)
	ReverseBytes(arr)
	fmt.Println(string(arr))
}
//一。。。。。。。。。。。。。。。。。。。。。。。。。
//字节数组转十六进制字符:

func BytesToHexString(arr []byte) string {
	return hex.EncodeToString(arr)
}

//十六进制转字节数组

func HexStringToBytes(a string) ([]byte, error) {
	arr,err := hex.DecodeString(a)
	return arr, err
}


//二。。。。。。。。。。。。。。。。。。。。。。。。。
//十六进制大端小端颠倒
func ReverseHexString(hexStr string) string {
	arr,_ := hex.DecodeString(hexStr)
	ReverseBytes(arr)
	return hex.EncodeToString(arr)
}

//字节数组大端小端颠倒
func ReverseBytes(data []byte)  {
	for i,n := 0,len(data)-1 ; i<n; i,n=i+1,n-1 {
		data[i],data[n] = data[n],data[i]
	}
}

