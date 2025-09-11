package main

import "sync"

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10) // slice of size 0, capacity 10.
}
