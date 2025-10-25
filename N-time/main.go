package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Study In Go Lang")

	presentTime:= time.Now()	
	fmt.Println(presentTime)

	// if u want only name of Day
	fmt.Println(presentTime.Format("Monday"))

	// if u want only current time
	fmt.Println(presentTime.Format("10:14:12"))

	// if u want only date
	fmt.Println(presentTime.Format("10-04-2023"))

	// if u want current all time date and all 
	fmt.Println(presentTime.Format("10-04-2023 10:43:34 Monday"))

	// Created Date
	createdTime := time.Date(2020,time.April,10,23,23,0,0,time.UTC)
	fmt.Println(createdTime)	
	fmt.Println(createdTime.Format("Monday"))
	fmt.Println(createdTime.Format("10:42:21"))
	fmt.Println(createdTime.Format("10-01-1202"))
}
