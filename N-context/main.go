package main

import (
	"contextingo/normalcontext"
	"contextingo/timeout"
	"contextingo/timeoutcancel"
	"fmt"
)

func main() {

	// This is Example for the Normal Context
	// Using TODO / Background
	// which means it never cancels / never be timout
	fmt.Println("THIS IS NORMAL CONTEXT")
	fmt.Println("-----------------------------------")
	normalcontext.ContextUse()
	fmt.Println("-----------------------------------")
	fmt.Println("NORMAL CONTEXT ENDS\n\n")

	// Its example of timeout context that we can cancels whenever we want
	fmt.Println("THIS IS CANCEL CONTEXT")
	fmt.Println("-----------------------------------")
	timeoutcancel.ContextUse()
	fmt.Println("-----------------------------------")
	fmt.Println("TIMOUT CANCEL ENDS\n\n")

	// Its example of timeout context that we can cancels whenever we want
	fmt.Println("THIS IS TIMOUT CONTEXT")
	fmt.Println("-----------------------------------")
	timeout.ContextUse()
	fmt.Println("-----------------------------------")
	fmt.Println("TIMOUT CONTEXT ENDS\n\n")
}
