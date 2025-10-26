package main

import (
	"fmt"
	"time"
)

func main() {

	// main has its own goroutine by default (the main goroutine)
	fmt.Println("Hello, N-concurrency-and-goroutines!")

	// it starts a new goroutine (lightweight thread)
	go greet()

	// main routine sleeps for 3 seconds to allow greet goroutine to complete
	time.Sleep(3 * time.Second)
}

func greet() {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Microsecond)
		fmt.Println("Hi", i)
	}
}
