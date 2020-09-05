package main

import "fmt"

type Circle struct {
	a, b int
}

// Define a function for Circle
func (circle Circle) area() int {
	return circle.a * circle.b
}

func main() {
	circle := Circle{a: 10, b: 20}
	fmt.Println("Area is: ", circle.area())
}
