package main

import (
	"fmt"
	"sync"
)

// Using WaitGroup to block until wg.Done()
var wg sync.WaitGroup

// counter is an int that will be used in
// incrementor
var counter int

// Mutex locks to prevent race conditions
var mutex sync.Mutex

func main() {
	// Add 2 items to the wait group this will
	// block until the wait groups call done.
	wg.Add(2)

	// Run both functions as go routine, they will
	// have to call Done.
	go incrementor("Foo: ")
	go incrementor("Bar: ")

	// Wait will block until all items in the wait group
	// will call done.
	wg.Wait()
}

// incrementor will increment the counter and use mutex
// to lock access to the counter.
func incrementor(s string) {

	for i := 0; i < 20; i++ {
		// use mutex to lock access to increment counter
		// this should ensure counter increments from 0 - 40
		// sequentially.
		mutex.Lock()

		counter++
		fmt.Println(s, i, "Counter: ", counter)

		// unlock mutex, allows access to increment counter
		mutex.Unlock()
	}

	wg.Done()
}
