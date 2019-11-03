package main

import (
	"fmt"
	"sync"
)

func CreateAccount(accountChannel chan string) {
	fmt.Println("Account Waiting to get Created")
	accountType := <-accountChannel
	fmt.Println(accountType)
}

func main() {
	var syncData sync.WaitGroup
	syncData.Add(2)
	accountChannel := make(chan string)
	go func() {
		defer syncData.Done()
		CreateAccount(accountChannel)
	}()

	go func() {
		defer syncData.Done()
		CreateAccount(accountChannel)
	}()
	accountChannel <- "Saving"
	accountChannel <- "Saving"
	fmt.Println("Account Complete...")
	syncData.Wait()
}
