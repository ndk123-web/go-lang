package timeout

import (
	"context"
	"fmt"
	"time"
)

func ContextUse() {

	// create a context with Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// this is for preventing memory leak
	// how ? currently 3 seconds timer , but what if our program done before 3 seconds
	defer cancel()

	DoSomething(ctx)
}

func DoSomething(ctx context.Context) {
	for { // infinite loop
		select {
		case <-ctx.Done(): // signal after 3 seconds
			fmt.Println("Context Cancelled..")
			return // end to infinite loop
		default:
			fmt.Println("Working..")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
