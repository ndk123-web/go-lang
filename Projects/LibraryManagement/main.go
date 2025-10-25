package main

import (
	"LibraryManagement/book"
	"LibraryManagement/user"
	"fmt"
)

func main() {
	fmt.Println("LIBRARY MANAGEMENT SYSTEM")

	// it means it sends address of object everytime when we call method 
	user := &user.User{"ndk", 20, []string{}}
	user.Borrow("khallas")

	book1 := &book.Book{"ndk", []string{}}
	book1.AddBooks("khallas", "opbevda")
}
