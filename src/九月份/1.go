package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	arr :=[]byte{'1','2','3','4','5'}
	fmt.Println(string(arr))
	str:=BytesToHexString(arr)
	fmt.Println(str)
	str=ReverseHexString(str)
	arr, _ =HexStringToBytes(str)
	fmt.Printf("%x\n", arr)
	ReverseBytes(arr)
	fmt.Println(string(arr))
}

func BytesToHexString(arr []byte)string {
	return hex.EncodeToString(arr)
}

func HexStringToBytes(s string)([]byte, error){
	arr,err:=hex.DecodeString(s)
	return arr,err
}

func ReverseHexString(hexStr string)string{
	arr,_:=hex.DecodeString(hexStr)
	ReverseBytes(arr)
	return hex.EncodeToString(arr)
}
func ReverseBytes(data []byte){
	for i,j:=0,len(data)-1 ;i<j;i,j=i+1,j-1  {
		data[i],data[j]=data[j],data[i]
	}
}