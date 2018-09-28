package main

import (
	"fmt"
)

func main() {
	// Main channel.
	c := make(chan int)

	// Asynchronously send ints to c
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}

		// Close the channel when finished.
		fmt.Println("Closing channel c")
		close(c)
	}()

	// for range will block and receive from channel c.
	// This will be unblocked when the channel is closed.
	for i := range c {
		fmt.Println("Receiving from c: ", i)
	}
}
