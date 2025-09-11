package main

import "sync"

var someSharedVariable bool = false

func main() {
	c := sync.NewCond(&sync.Mutex{})
	// c.L is the 'locker' (sync.Locker) inside the cond.
	// This is locking the mutex, to check the conditions (critical section).
	c.L.Lock()
	for !checkCondition() {
		// The condition is false, so release the lock to let another process
		// eventually update the condition to true.

		// --Important side-effect behavior--
		// c.Wait() **unlocks the lock when entered**
		// and      **locks  the  lock  when exited**
		c.Wait()
	}
	c.L.Unlock()
}

func checkCondition() bool {
	// Check some global state that needs memory synchronization to access...
	return someSharedVariable
}
