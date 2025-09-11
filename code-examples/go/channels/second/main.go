package main

import "fmt"

func main() {
	// It's useful to indicate to another goroutine that no more values will be sent.
	// Close the channel!

	stringStream := make(chan string)
	go func() {
		defer close(stringStream) // convention to close a channel with defer.
		for i := 1; i <= 5; i++ {
			stringStream <- fmt.Sprintf("%d", i)
		}
	}()
	for str := range stringStream {
		fmt.Printf("%v ", str)
	}

	intStream := make(chan string)
	close(intStream)

	integer, ok := <-intStream
	fmt.Printf("(%v): %v", ok, integer)
	fmt.Printf("(%v): %v", ok, integer)
	fmt.Printf("(%v): %v", ok, integer)

	// Since you can infinitely read from a closed channel, you could signal
	// to ALL listening goroutines by closing the channel.
}
