package main

import (
	"fmt"
)

func main() {
	fmt.Println("Producer Consumer")

	// Create a channel for communication
	ch := make(chan int)

	// Start the producer in a goroutine
	go producer(ch)

	// Start the consumer in a goroutine (or run it in main)
	// We'll run it in main to block until done
	consumer(ch)
}

// producer generates numbers and sends them to the channel
func producer(ch chan<- int) {
	fmt.Println("Producer Started")
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	// Close the channel to signal we are done
	close(ch)
}

// consumer reads from the channel until it is closed
func consumer(ch <-chan int) {
	fmt.Println("Consumer Started")
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
