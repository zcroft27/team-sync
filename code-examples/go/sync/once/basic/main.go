package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once
	var incrementsWaitGroup sync.WaitGroup
	incrementsWaitGroup.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer incrementsWaitGroup.Done()
			once.Do(increment)
		}()
	}

	incrementsWaitGroup.Wait()

	fmt.Printf("Count is %d\n", count)
}
