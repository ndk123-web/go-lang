package normalcontext

import (
	"context"
	"fmt"
	"time"
)

func ContextUse() {

	// create parent context which will never cancels and never timeout
	ctx := context.Background()

	// send the ctx to DoSomething()
	DoSomething(ctx)
}

func DoSomething(ctx context.Context) {
	fmt.Println("Starting")
	time.Sleep(3 * time.Second)
	fmt.Println("Ending")
}
