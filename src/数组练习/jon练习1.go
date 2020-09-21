package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("randn = ", rand.Intn(15))
	}
}