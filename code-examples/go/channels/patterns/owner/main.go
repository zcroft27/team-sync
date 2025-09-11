package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	// Because the owner owns the returned channel:
	// Won't write to a nil channel and deadlock
	// Won't close a nil channel and panic
	// Won't write to a closed channel and panic
	// Won't close a channel more than once and panic
	// Won't have improper-typed writes and error at compile time

	// Consumer only needs to:
	// Know when a channel is closed
	// Handle blocking for any reason

	// If these values were passed onto a shared buffer,
	// all future developers of readers would need to follow
	// an agreed-upon mutex dance, or else, face a data race.
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
