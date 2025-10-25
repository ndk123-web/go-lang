package main

import "fmt"

func main(){
	var intArr [3]int32 
	// var intArr2 [3]int32 = [3]int32{1,2,3}
	// var intArr3 = [...]int32{4,5,6}

	// access array
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// access memory address
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	// we can initiliase specific indexes
	arr1 := [5]int{1:10,2:40}
	fmt.Println(arr1)

	// go automatically counts the values
	arr2 := [...]int64{1,2,3,4}

	// len() built in function to get the length of array
	fmt.Printf("Legth %v",len(arr2))

	// Append slices
	var intSlice1 = []int {1,3,2} 	// this is slice
	var intSlice2 = []int {4,56,6} 	// this is slice
	var final []int = append(intSlice1,intSlice2...)
	fmt.Println(final)
}