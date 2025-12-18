package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	someSharedVariable := [2]int{1, 2}

	var rw sync.RWMutex
	var wg sync.WaitGroup

	readFunc := func() {
		defer wg.Done()
		rw.RLock()
		fmt.Printf("Goroutine reading: %d\n", someSharedVariable[0])
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine done reading")
		rw.RUnlock()
	}

	wg.Add(4)
	go readFunc()
	go readFunc()
	go readFunc()
	go readFunc()
	wg.Wait()

}
