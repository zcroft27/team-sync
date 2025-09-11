package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]any, 0, 10) // slice of type any, size 0, capacity 10.

	// Anonymous function to simulate removing the first element in the queue.
	// Sleeps a delay, locks, "removes" the first element, unlocks, and signals.
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock() // Doesn't really matter if Unlock is before Signal
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{}) // struct{} is an empty type, and struct{}{} is a 0 byte instance.
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
