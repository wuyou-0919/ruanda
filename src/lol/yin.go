package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lol/wen"
	"net/http"
)

func  main(){
	fmt.Println("欢迎来到英雄联盟")
	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	client := http.Client{}
	request, err := http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(request)
		return
	}
	resp, err :=client.Do(request)
	if err != nil {
		fmt.Println(request)
		return
	}
	heroData,err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(request)
		return
	}
	fmt.Println(string(heroData))
	var herolist  wen.HeroList
	err = json.Unmarshal(heroData, &herolist)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(herolist)
	fmt.Println(len(herolist.Hero))
	fmt.Println(herolist.Version)
	fmt.Println(herolist.FileName)
	fmt.Println(herolist.FileTime)
}
