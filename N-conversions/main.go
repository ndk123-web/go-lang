package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v", err.Error())
	} else {

		fmt.Printf("Inp as String: %v", input)

		// TrimSpace() trims the "\n"
		num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil {
			fmt.Printf("%v", err.Error())
		} else {
			fmt.Printf("Inp as Num: %v", num)
		}
	}
}
