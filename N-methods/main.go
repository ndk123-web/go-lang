package main

import "fmt"

// Main Point
// it sends copy to the methods

type User struct {
	Name   string
	Age    int
	Status bool
}

type Teachers struct {
	Teacher string
	Course  []int
}

func main() {
	fmt.Println("Methods in Go Lang")

	// Its Like Creating Object of class
	user := User{"ndk", 20, true}
	user.GetStatus()
	user.GetUserName()

	teacher := Teachers{"Mahesh", []int{1, 2, 3}}
	fmt.Println("T ", teacher)

	deferConcept()
}

// methods is of class functions
// here methods are functions of structs
// it sends copy to the methods
func (u User) GetStatus() {
	u.Status = false
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) GetUserName() {
	fmt.Println("User Name: ", u.Name)
}

// output -> Hello , Three , Two , One
// LIFO -> Defer Stack -> One , Two, Three , Prints Hello then
// in LIFO Three then Two then One
func deferConcept() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
}
