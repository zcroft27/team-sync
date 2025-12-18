package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var bidirectionalStream chan string // bidirectional, unbuffered channel.
	bidirectionalStream = make(chan string)

	// readOnlyStream := make(<-chan interface{})
	// writeOnlyStream := make(chan<- interface{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		readFromChannel(bidirectionalStream)
	}()

	// main goroutine immediately sends "foo"
	// but reader goroutine waits 2 seconds.
	fmt.Println("main goroutine sending!")

	bidirectionalStream <- "foo"

	fmt.Println("main goroutine unblocked!")

	wg.Wait() // necessary to stop the main goroutine from exiting
	//           before the reader goroutine finishes, killing the reader goroutine early.
}

// Implicitly converting a bidirectional channel into a read-only channel.
func readFromChannel(readStream <-chan string) {
	val := <-readStream
	fmt.Println("Read from read-only channel: ", val)
}
