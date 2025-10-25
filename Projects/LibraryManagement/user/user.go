// Package user handles user-related functionality
package user

import "fmt"

// Logs stores all book borrowing logs
var Logs []string

// User represents a library user with name, age, and borrowed books
type User struct {
	Name  string   // User's name
	Age   int      // User's age
	Books []string // List of borrowed books
}

// Borrow allows a user to borrow a book and logs the transaction
func (u *User) Borrow(bookName string) {
	// Create log entry
	str := "Student " + u.Name + " Borrow Book" + bookName
	
	// Update logs and user's book list
	Logs = append(Logs, str)
	u.Books = append(u.Books, bookName)

	// Display all logs
	fmt.Println("Total Logs")
	for i, v := range Logs {
		fmt.Println(i+1, " ", v)
	}
}
