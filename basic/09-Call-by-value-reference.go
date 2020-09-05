package main

import "fmt"

func main() {
	var x, y int = 10, 12
	fmt.Println("Call By Value:")
	fmt.Println(byValue(x, y))
	fmt.Println("x and y is: ", x, y)

	fmt.Println("\nCall By Reference:")
	fmt.Println(byReference(&x, &y))
	fmt.Println("x and y is: ", x, y)
}

func byValue(x, y int) int {
	x = 20
	y = 30
	return x + y
}

func byReference(x *int, y *int) int {
	*x = 20
	*y = 30
	return *x + *y
}
