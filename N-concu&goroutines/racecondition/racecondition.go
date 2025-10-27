package racecondition

import (
	"fmt"
	"sync"
)

var mut sync.Mutex
var wg sync.WaitGroup
var scores []int = []int{0}

func RaceConditionExample() {

	fmt.Println("Race Condition")

	// Added 3 Goroutines in WaitList 
	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("First Routine")

		// Lock the scores shared Memory
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()

		// Send to the Wait Group
		wg.Done()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Second Routine")

		// Lock the scores shared Memory
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()

		// Send to the Wait Group
		wg.Done()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Third Routine")

		// Lock the scores shared Memory
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()

		// Send to the Wait Group
		wg.Done()
	}(&wg, &mut)

	// Wait Till All the Routines to be complete
	wg.Wait()

	fmt.Println("Scores: ", scores)
}
