package main 

import "fmt"

func main(){
	fmt.Print("Welcome To Pointers \n")

	var one int= 1

	// pointer creation 
	// if no init then default value in pointer is nil
	var cone *int = &one

	fmt.Println(cone)	// address 
	fmt.Println(*cone)	// data inside that address
}