package main

import (
	"fmt"
	"errors"
)

func main(){
	// void
	printMe()

	// pass value of copy
	send2Num(1,2)

	var mul int64 = multiply(10,10)
	fmt.Println("Mul: ",mul)

	var q , r, err = division(5,0)
	if err != nil{
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("The Quotient %v and Remainder %v",q,r)
	}

}

func printMe(){
	fmt.Println("Hello")
}

// it takes first name and then data type
// func some(variable type)
func send2Num(a int64, b int64 ){
	fmt.Println(a+b)
}

// return multiplication of 2 numbers
// after func we need to specify return type
func multiply(a int64 , b int64) int64 {
	var c int64 = int64(a*b)
	return c  
}

// can return multiple values
func division(a int64, b int64) (int64,int64,error){

	var err error

	if b == 0 {
		err = errors.New("Can't Divide By Zero")
		return 0,0,err
	}
	var q int64 = int64(a/b)  // explicit cast 
	var r int64 = int64(a%b)

	return q,r,err
}