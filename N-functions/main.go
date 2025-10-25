package main

import "fmt"

func main() {
	fmt.Println("Functions in Go Lang")

	nums := []int{1, 2, 3, 4, 5}
	proAdder(nums...) // (1,2,3,4,5) instead of nums [1,2,3,4,5]
}

// take normal (1,2,3,4,5) and then get as type slice
func proAdder(nums ...int) {
	fmt.Printf("%T\n", nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}
