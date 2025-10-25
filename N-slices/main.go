package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go Lang")

	slice1 := []int {1,2,3};
	arr := [4]int {1,2,3,4};

	fmt.Printf("Type %T \n",slice1) // type slice
	fmt.Printf("Type %T \n",arr) 	// type array 

	slice1 = append(slice1, 3,4,5)
	fmt.Println(slice1)

	// Maximum we are going to use slices 
	slice1 = append(slice1[1:5])
	fmt.Println(slice1)

	// create a slice of fefault size 4 
	highScores := make([]int , 4)
	highScores[0] = 11
	highScores[1] = 121
	highScores[2] = 14
	highScores[3] = 11

	// Dynamically Increases Slices
	highScores = append(highScores, 1,2,3)

	fmt.Println("Before",highScores)
	sort.Ints(highScores)	// sort
	fmt.Println("After",highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // gives true if sorted else false

	// Remove value from Slices
	// assume removing 2nd as 1-index
	highScores = append(highScores[:2],highScores[3:]...)
	fmt.Println("Removed 2nd 1-index",highScores)
}
