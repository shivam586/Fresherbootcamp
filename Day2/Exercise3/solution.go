package main

import (
	"fmt"
	"sync"
)

var accountbalance = 500
var mutex sync.Mutex

func withdrawal(wg *sync.WaitGroup) {
	var withdraw int
	fmt.Println("Please enter the amount to be withdrawn")
	fmt.Scanln(&withdraw)
	if withdraw < accountbalance {
		mutex.Lock()
		accountbalance = accountbalance - withdraw
		mutex.Unlock()

	}
    wg.Done()
}

func deposit(wg *sync.WaitGroup) {
	var depositmoney int
	fmt.Println("please enter the amount to be deposited")
	fmt.Scanln(&depositmoney)
	mutex.Lock()
    accountbalance = accountbalance + depositmoney
	mutex.Unlock()
    wg.Done()
}

func main() {
 var waiting sync.WaitGroup

  	waiting.Add(2)
  	go deposit(&w)
  	go withdrawal(&w)
  waiting.Wait()

  fmt.Println("The balance in account :-", accountbalance)

}
