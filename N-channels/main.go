package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(3)

	// sender goroutine
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 10  // send
		ch <- 20  // send
		ch <- 30  // send
		close(ch) // close the channel after sending (only sender should close)
	}(myCh, &wg)

	// receiver goroutine 1
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, ok := <-ch // listen
		if ok {
			fmt.Println("Received 1:", val)
			fmt.Println("Received 3:", <-ch) // listen
		} else {
			fmt.Println("Channel closed, no value received")
		}
	}(myCh, &wg)

	// receiver goroutine 2
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, ok := <-ch // listen
		if ok {
			fmt.Println("Received 2:", val)
		} else {
			fmt.Println("Channel closed, no value received")
		}
	}(myCh, &wg)

	wg.Wait()
}
