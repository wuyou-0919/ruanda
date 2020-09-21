package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var result1 sync.WaitGroup
func main() {
	result1.Add(3)
	go one1()
	go one1()
	go one1()
	result1.Wait()
}
func one1()  {
	for i :=0;i < 99;i++{
		time.Sleep(200)
		fmt.Printf("%d\n",rand.Intn(1000))
	}
	result1.Done()
}