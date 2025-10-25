package book

import (
	"LibraryManagement/user"
	"fmt"
)

var AllBooks []string

type Book struct {
	Author string
	Books  []string
}

func (b *Book) AddBooks(books ...string) {
	for _, v := range books {
		AllBooks = append(AllBooks, v)
	}

	totalLogs := user.Logs
	for _, v := range totalLogs {
		fmt.Println(v)
	}

	fmt.Println("Total Books")
	for _, v := range books {
		fmt.Println(v)
	}
}
