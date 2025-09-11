package main

import (
	"math/rand"
	"sync"
	"time"
)

type Fork struct {
	id    int
	mutex sync.Mutex
	taken bool
	owner int
}

var forks []Fork

func naivePhilosopher(id int, wg *sync.WaitGroup) {
	// Imagine visually the philosophers each holding a fork in their left hand,
	// waiting to grab the fork to their right.
	// DEADLOCK!
	defer wg.Done()

	for range 2 {
		// Think
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		leftFork := id
		rightFork := (id + 1) % len(forks)

		// DEADLOCK TRAP: Always acquire left fork first, then right fork.
		forks[leftFork].mutex.Lock()
		forks[leftFork].taken = true
		forks[leftFork].owner = id

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		forks[rightFork].mutex.Lock() // DEADLOCK!
		forks[rightFork].taken = true
		forks[rightFork].owner = id

		// Eat
		time.Sleep(1 * time.Second)

		forks[leftFork].taken = false
		forks[leftFork].owner = -1
		forks[leftFork].mutex.Unlock()

		forks[rightFork].taken = false
		forks[rightFork].owner = -1
		forks[rightFork].mutex.Unlock()
	}
}
