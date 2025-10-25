package main

import (
	"LibraryManagement/book"
	"LibraryManagement/user"
	"fmt"
)

func main() {
	fmt.Println("LIBRARY MANAGEMENT SYSTEM")

	user := &user.User{"ndk", 20, []string{}}
	user.Borrow("khallas")

	book1 := &book.Book{"ndk", []string{}}
	book1.AddBooks("khallas", "opbevda")
}
