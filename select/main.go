package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	// Infinite loop (blocking).
	for {
		// Select will wait to receive or write to multiple
		// go routines.
		select {
		// Send x to c
		case c <- x:
			x, y = y, x+y
		// Receive from quit if there is something from the
		// channel.
		case <-quit:
			fmt.Println("Quit was called")
			return
		}
	}
}

func main() {
	// c is the main channel we will be using
	c := make(chan int)

	// quit will be the semaphore channel. When sending
	// to the quit channel, this will tell the subscriber
	// that the observer has stopped emitting.
	quit := make(chan int)

	// Asynchronously read from channel 'c'. Keep reading
	// up to n times. This is to ensure the loop doesn't block.
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Reading from channel: ", <-c)
		}

		// Send 0 to quit channel. This will act as a
		// semaphore and will tell fibonacci() to stop
		// sending to c.
		quit <- 0
	}()

	// Pass channels to fibonacci()
	fibonacci(c, quit)
}
