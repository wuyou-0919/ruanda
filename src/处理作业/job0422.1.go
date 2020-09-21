package main

import (
	"fmt"
	"sync"
	"time"
)
var capacity =300
var ch chan int
var new sync.WaitGroup
var product sync.Mutex
func main(){
	new.Add(4)
	go factory()
	go market1()
	go market2()
	go market3()
	new.Wait()
}
func factory(){

	for  capacity >0 {
		product.Lock()
		time.Sleep(8)
		if capacity< 300 {
			if capacity==0 {
				product.Unlock()
				break
			}
			capacity++
		}else{
			product.Unlock()
			break
		}
		fmt.Printf(" market库存：%d个\n",capacity)
		product.Unlock()
	}
	new.Done()
}

func market1(){
	for {
		product.Lock()
		if capacity>0 {
			time.Sleep(20)
			capacity--
		}else if capacity<=0 {
			product.Unlock()
			break
		}
		fmt.Printf("market1库存剩余：%d个\n",capacity)
		product.Unlock()
	}
	new.Done()
}
func market2(){
	for {
		product.Lock()
		if capacity>0 {
			time.Sleep(20)
			capacity--
		}else if capacity<=0{
			product.Unlock()
			break
		}
		fmt.Printf("market2库存剩余：%d个\n",capacity)
		product.Unlock()
	}
	new.Done()
}
func market3(){
	for {
		product.Lock()
		if capacity>0 {
			time.Sleep(20)
			capacity--
		}else if capacity<=0{
			product.Unlock()
			break
		}
		fmt.Printf("market库存剩余：%d个\n",capacity)
		product.Unlock()
	}
	new.Done()
}
