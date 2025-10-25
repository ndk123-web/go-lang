package main

import (
	"fmt"
)

func main() {

	// 1. Slice Append
	// slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{4, 5, 6}

	slice := append(slice1,slice2...)
	slice = append(slice,slice3...)
	fmt.Print(slice)


	fmt.Print("Hi")
}
