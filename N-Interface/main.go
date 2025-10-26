package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func printShape(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func main() {
	fmt.Println("Interface in GO Lang")

	// send the address so Circle methods can take Address and can modifiy original data
	c := &Circle{12}
	r := &Rectangle{10, 20}

	printShape(c)
	printShape(r)
}
