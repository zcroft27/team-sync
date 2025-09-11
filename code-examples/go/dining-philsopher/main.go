package main

import (
	"fmt"
	"sync"
	"time"
)

func waiter(requests chan int, grants chan int, returns chan int) {
	forks := 5

	for {
		select {
		case id := <-requests:
			if forks >= 2 {
				forks -= 2
				fmt.Printf("Waiter: Granted 2 forks to philosopher %d (%d left)\n", id, forks)
				grants <- id // Signal that philosopher can eat.
			} else {
				fmt.Printf("Waiter: No forks for philosopher %d, trying again...\n", id)
				go func() {
					time.Sleep(100 * time.Millisecond)
					requests <- id // Try again later.
				}()
			}

		case id := <-returns:
			forks += 2
			fmt.Printf("Waiter: Got 2 forks back from philosopher %d (%d total)\n", id, forks)
		}
	}
}

func philosopher(id int, requests chan int, grants chan int, returns chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for meal := 1; meal <= 3; meal++ {
		// Think:
		fmt.Printf("Philosopher %d thinking (meal %d/3)\n", id, meal)
		time.Sleep(500 * time.Millisecond)

		// Request forks:
		fmt.Printf("Philosopher %d hungry, requesting forks\n", id)
		requests <- id

		// Wait for waiter to grant forks:
		<-grants
		fmt.Printf("Philosopher %d got forks, eating\n", id)

		// Eat
		time.Sleep(500 * time.Millisecond)

		// Return forks:
		returns <- id
		fmt.Printf("Philosopher %d finished meal %d\n", id, meal)
	}

	fmt.Printf("Philosopher %d done with all meals.\n", id)
}

func main() {
	requests := make(chan int)
	grants := make(chan int)
	returns := make(chan int)

	go waiter(requests, grants, returns)

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go philosopher(i, requests, grants, returns, &wg)
	}

	wg.Wait()
	fmt.Println("All philosophers finished.")
}
