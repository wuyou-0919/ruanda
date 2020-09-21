package main

import (
	"errors"
	"fmt"
)
func main() {
	information := Person{
		name:   "吴忧",
		age:   0,
		address: "江西省九江市",
	}

	a, b := when(information)
	if b != nil {
		fmt.Println(b)
	} else {
		if a == true {
			fmt.Println("已成年")
		} else {
			fmt.Println("未成年")
		}
	}
}

type Person struct {
	name string
	age int
	address string
}

func when(p Person) (bool,error){
	if p.age>=18{
		return true,nil
	}else if 0<p.age{
		return false,nil
	}else{
		var err error =errors.New("年龄不合法")
		return false,err
	}
}

