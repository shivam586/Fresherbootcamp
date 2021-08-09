package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var student[5] int
var sum = 0
func Rating(wg *sync.WaitGroup, ch chan bool, id int) {

	ch <- true
	fmt.Println("please enter rating")
	fmt.Scanln(&student[id])
	time.Sleep(time.Duration(rand.Float32()))
	sum = sum + student[id]
	    <- ch
	wg.Done()
}

func main () {
     var wtng sync.WaitGroup
     channel := make(chan bool, 1)

     for i:=0 ; i<5 ; i++ {
     	wtng.Add(1)
		 go Rating(&wtng,channel,i)
	 }
	 wtng.Wait()
     fmt.Println("Average rating given to teacher",sum/5)
}
