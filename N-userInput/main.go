package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome User"
	fmt.Println(welcome)

	
	reader := bufio.NewReader(os.Stdin)

	// comma ok Syntax (it wait user to read string till \n)
	input , err := reader.ReadString('\n')
	
	if err != nil{
		fmt.Printf("There is Error: %v",err.Error())
	}else{
		fmt.Printf("Input: %v",input)
	}
}