package user

import "fmt"

var Logs []string

type User struct {
	Name  string
	Age   int
	Books []string
}

func (u *User) Borrow(bookName string) {
	str := "Student " + u.Name + " Borrow Book" + bookName
	Logs = append(Logs, str)
	u.Books = append(u.Books, bookName)

	fmt.Println("Total Logs")
	for i, v := range Logs {
		fmt.Println(i+1, " ", v)
	}
}
