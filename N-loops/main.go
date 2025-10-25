package main

import "fmt"

func main() {
	fmt.Println("Loops in GO Lang")

	// Loop Through (For Loop)
	days := []string{"Mon", "Tue", "Wed", "Thur", "Fri", "Sat", "Sun"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Loop Through (For Each Like Loop)
	for i := range days {
		fmt.Println(days[i])
	}

	// _ here is index 
	for _ , value := range days {
		fmt.Println(value)
	}

	// while loop
	val := 1
	for val < 10{
		val++;
		if val == 5{
			goto lco 
		}
	}
	fmt.Println(val)

	lco:
		fmt.Println("Hello Its Label (Goto)")
}
