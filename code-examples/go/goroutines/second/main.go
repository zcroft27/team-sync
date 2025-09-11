package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // sempaphore in C/POSIX.
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}
