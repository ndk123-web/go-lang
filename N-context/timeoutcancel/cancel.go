package timeoutcancel

import (
	"context"
	"fmt"
	"time"
)

func ContextUse() {

	// craetes context with WithCance()
	ctx, cancel := context.WithCancel(context.Background())

	// this will run separately because it will run separate go routine
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("caceling context...")

		// we send the signal that will signals that context has been cancelled
		cancel()
	}()

	// it will run immediately because this is in main gorouting
	DoSomething(ctx)
}

func DoSomething(ctx context.Context) {

	// its likw while true
	for {

		// select use for channels communications
		select {

		// whatever we send the cancel() it trigger case 1
		case <-ctx.Done():
			fmt.Println("context cancelled..")
			return
		default:
			fmt.Println("Context Working..")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
