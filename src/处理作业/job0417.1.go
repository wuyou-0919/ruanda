package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var result sync.WaitGroup
func main() {
  	rand.Seed(time.Now().Unix())
	result.Add(2)
  	go one()
  	go two()
	result.Wait()
}
func one()  {
	for i :=0;i < 99;i++{
		time.Sleep(200)
		fmt.Printf("%d\n",rand.Intn(1000))
	}
	result.Done()
}
func two()  {
	for i :=0;i < 199;i++{
		time.Sleep(100)
		fmt.Printf("%c\n",rand.Intn(26)+65)

	}
	result.Done()

}