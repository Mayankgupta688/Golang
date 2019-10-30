package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 5000

var mu = &sync.Mutex{}
var waitGrp sync.WaitGroup

func main() {
	waitGrp.Add(5)
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	waitGrp.Wait()
}

func ReadWriteSynchronousData() {
	defer waitGrp.Done()
	mu.Lock()
	balance = balance + 500
	time.Sleep(2 * time.Second)
	fmt.Println(balance)
	mu.Unlock()
}
