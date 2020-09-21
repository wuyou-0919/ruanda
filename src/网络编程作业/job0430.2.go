package main

import (
	"fmt"
	"net/http"
)
func main() {
	catalog := http.ListenAndServe("127.1.1.100:9000",http.FileServer(http.Dir("C:/Go")))
	if catalog != nil {
		fmt.Println(catalog.Error())
	}
}
