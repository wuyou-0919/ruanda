package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	client := http.Client{}

	for a := 0; a <= 10; a++ {//页数
		fmt.Printf("采集第 %d页数据\n",a)
		offsetNum := strconv.Itoa(a * 25)
		request,err := http.NewRequest("GET","https://movie.douban.com/top250?start="+offsetNum+"&filter=",nil)
		if err != nil {
			log.Fatal(err.Error())
			return
		}


		request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

		request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		request.Header.Add("Cache-Control","no-cache")
		request.Header.Add("Connection","keep-alive")
		request.Header.Add("Host","movie.douban.com")
		request.Header.Add("Pragma","no-cache")
		request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 Edg/83.0.478.45")


		response,err := client.Do(request)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		htmlBytes,err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err.Error())
			return
		}


		html := string(htmlBytes)

		movieIdReg := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/">`)
		idSlice := movieIdReg.FindAllStringSubmatch(html,-1)


		nameReg := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
		nameSlice := nameReg.FindAllStringSubmatch(html,-1)



		ratReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
		ratSlice := ratReg.FindAllStringSubmatch(html,-1)


		voteReg := regexp.MustCompile(`<span>(.*?)人评价</span>`)
		voteSlice := voteReg.FindAllStringSubmatch(html,-1)


		descReg := regexp.MustCompile(`<span class="inq">(.*?)</span>`)
		descSlice := descReg.FindAllStringSubmatch(html,-1)


		coverImgReg := regexp.MustCompile(`src="(.*?) class="">`)
		imgSlice := coverImgReg.FindAllStringSubmatch(html,-1)

		fmt.Println("电影编号   电影名称  评分  评价人数  一句话总结  封面图")
		for i :=0; i < len(nameSlice); i++{
			fmt.Printf("%s  %s  %s  %s %s %s \n",
				idSlice[i][1],
				nameSlice[i][1],
				ratSlice[i][1],
				voteSlice[i][1],
				descSlice[i][1],
				imgSlice[i][1])
		}


		database,err :=sql.Open("mysql","root:20000919@tcp(127.0.0.1:3306)/orve_movie?charset=utf8")
		defer database.Close()
		if err != nil{
			log.Fatal(err.Error())
			return
		}


		for j :=0; j < len(nameSlice); j++{
			_, err :=database.Exec("insert into " +
				"movie_info(movie_id, movie_name, movie_rate, movie_vote, movie_desc, movie_url)" +
				"values(?,?,?,?,?,?)",
				idSlice[j][1],
				nameSlice[j][1],
				ratSlice[j][1],
				voteSlice[j][1],
				descSlice[j][1],
				imgSlice[j][1])
			if err != nil {
				fmt.Println("ERR")
				log.Fatal(err.Error())
				break
			}
		}

		fmt.Println()
	}
	fmt.Println("结束")
}