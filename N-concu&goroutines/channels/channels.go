package channels

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func ChannelsExample() {

	// Deadlock Error
	// should some routine must listen that channel
	// myCh <- 5
	// fmt.Println(<-myCh)

	// create myCh channel
	myCh := make(chan int, 2)
	fmt.Printf("%T \n", myCh)

	wg.Add(2)

	// here <- means sending the value
	// SEND ONLY CHANNEL
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 10 // 1st sender
		ch <- 20 // 2nd sender

		// close the channel
		close(ch)

		wg.Done()
	}(myCh, &wg)

	// here <- means receiving value
	// RECEIVE ONLY CHANNEL
	go func(ch <-chan int, wg *sync.WaitGroup) {

		// To Check is Channel is Open Or Not
		val, isChannlOpen := <-ch
		if isChannlOpen {
			fmt.Println("Channel is Open: ", val)
		}

		fmt.Println(<-ch) // 1st listener
		fmt.Println(<-ch) // 2nd listener

		wg.Done()
	}(myCh, &wg)

	wg.Wait()

}