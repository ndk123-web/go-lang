// Package book handles book-related functionality
package book

import (
	"LibraryManagement/user"
	"fmt"
)

// AllBooks stores a global list of all books in the library
var AllBooks []string

// Book represents a book with its author and list of book titles
type Book struct {
	Author string   // Name of the book's author
	Books  []string // List of books by this author
}

// AddBooks adds multiple books to the library and displays logs
// It accepts a variadic parameter of book titles to add
func (b *Book) AddBooks(books ...string) {
	// Add each book to the global AllBooks list
	for _, v := range books {
		AllBooks = append(AllBooks, v)
	}

	// Display all user logs
	totalLogs := user.Logs
	for _, v := range totalLogs {
		fmt.Println(v)
	}

	// Display all books being added
	fmt.Println("Total Books")
	for _, v := range books {
		fmt.Println(v)
	}
}
