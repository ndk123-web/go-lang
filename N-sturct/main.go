package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func main() {
	fmt.Println("Struct in Go Lang")

	// no inheritance
	// no parent , child

	ndk := User{"Ndk", 20, true, "ndk@gmail.com"}
	fmt.Println(ndk)

	// +v for structs
	fmt.Printf("Details: %+v \n", ndk)
	fmt.Printf("Name: %v and Age: %v \n", ndk.Name, ndk.Age)
}

