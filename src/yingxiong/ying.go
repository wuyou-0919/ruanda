package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"yingxiong/yx"
)

func main() {
	fmt.Println("welcome to the league of legends")
	url:="https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	client:=http.Client{}
	request,err1:=http.NewRequest("GET",url,nil)
	if err1 != nil {
		fmt.Println(request)
		return
	}
	response,err2:=client.Do(request)
	if err2 != nil {
		fmt.Println(request)
		return
	}
	heroData,err3:=ioutil.ReadAll(response.Body)
	if err3 != nil {
		fmt.Println(request)
		return
	}
	fmt.Println(string(heroData))
	var heroLi yx.HeroList
	err4:=json.Unmarshal(heroData,&heroLi)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}
	for i :=0;i < len(heroLi.Hero) ;i++  {
		fmt.Printf("第%d个英雄：%s %s %s %s %s",i+1,heroLi.Hero[i].Name,heroLi.Hero[i].Alias,heroLi.Hero[i].Title,heroLi.Hero[i].Roles,heroLi.Hero[i].IsWeekFree)
		fmt.Println()
	}
	fmt.Println(heroLi.Version)
	fmt.Println(heroLi.FileName)
	fmt.Println(heroLi.FileTime)
}
