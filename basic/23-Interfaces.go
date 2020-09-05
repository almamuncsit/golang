package main

import (
	"fmt"
	"math"
)

// Shape Interface
type Shape interface {
	area() float64
}

// Circle struct
type Circle struct {
	r float64
}

// Implement area() of Circle
func (circle Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

// Rectangle struct
type Rectangle struct {
	height, width float64
}

// Implement area() of Rectangle struct
func (rectangle Rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{r: 12}
	fmt.Println("Area of Circle is: ", circle.area())

	rectangle := Rectangle{height: 12, width: 30}
	fmt.Println("Area of Rectangle is: ", rectangle.area())

	fmt.Println("")
	fmt.Printf("Area of Circle is:  %.2f\n", getArea(circle))
	fmt.Printf("Area of Rectangle is: %.2f\n", getArea(rectangle))

}
