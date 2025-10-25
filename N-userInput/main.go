package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome User"
	fmt.Println(welcome)

	// create reader using bufio and os.Stdin
	reader := bufio.NewReader(os.Stdin)

	// comma ok Syntax (it wait user to read string till \n)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There is Error: %v", err.Error())
	} else {
		fmt.Printf("Input: %v", input)
	}

	// switch case
	num := 1

	switch num {
	case 1:
		fmt.Println("Its 1")
		fallthrough				// it means below one also run 
	case 2:
		fmt.Println("Its 2")
	case 3:
		fmt.Println("Its 3")
	}
}
