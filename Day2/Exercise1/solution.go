package main

import (
	"fmt"
	"sync"
)

var frequencycount = make(map[byte]int)

func Count(wg *sync.WaitGroup, ch chan bool, str string) {
       ch <- true
       for i:=0 ; i<len(str) ; i++ {
       	  frequencycount[str[i]] = frequencycount[str[i]] + 1
	   }
	   <- ch
	   wg.Done()
}

func main () {
	var wtng sync.WaitGroup
	chnl := make(chan bool, 1)
	fmt.Println("please enter the strings")
	var input[5] string
	for i:=0 ; i<5 ;i++ {
		fmt.Scanln(&input[i])
	}
	for i:=0 ; i<5 ;i++ {
		wtng.Add(1)
		go Count(&wtng,chnl,input[i])
	}
	wtng.Wait()
	for key,value := range frequencycount {
		fmt.Printf("\"%c\" : %d \n",key,value)
	}


}