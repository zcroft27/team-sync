package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	myChannel := make(chan int)
	otherChannel := make(chan int)

	wg := sync.WaitGroup{}
	readerFunc := func() {
		defer wg.Done()
		select {
		case <-myChannel:
			fmt.Println("received on my channel")
		case <-otherChannel:
			fmt.Println("received on other channel")
		}
	}

	wg.Add(1)
	go readerFunc()

	time.Sleep(2 * time.Second)

	fmt.Println("sending!")
	myChannel <- 2

	wg.Wait()

	// this will block forever!
	// select {}

	// c1 := make(chan interface{})
	// close(c1)
	// c2 := make(chan interface{})
	// close(c2)
	// var c1Count, c2Count int
	// for i := 1000; i >= 0; i-- {
	// 	select {
	// 	case <-c1:
	// 		c1Count++
	// 	case <-c2:
	// 		c2Count++
	// 	}
	// }
	// fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

	// Closed channels always send their default values.
	// Go runtime will pretty fairly pick between channels if they are all receiving.

	// time package offers nifty channel to implement timeouts!
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}

	// You can have a default case to immediately go to
	// if all other cases are blocking.
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
