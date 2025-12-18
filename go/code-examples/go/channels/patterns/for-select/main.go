package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()
	workCounter := 0

	// Syntax: 'loop:' is a 'label' for the 'for loop'.
	// This is so the inner 'break' breaks out of the for loop,
	// not just the select statement.
loop:
	for { // Infinite for loop. Same as 'while true' in other languages.
		select {
		case <-done:
			break loop
		default:
		}

		// Until another goroutine tells me to stop,
		// do work on this current task.
		workCounter++
		time.Sleep(600 * time.Millisecond)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
