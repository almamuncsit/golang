package main

import "fmt"

func main() {
	var x, y int = 10, 12
	result := max(x, y)
	fmt.Println("Max is:", result)

	// Swap X U
	first, second := swap(x, y)
	fmt.Println(first, second)
}

func max(num1 int, num2 int) int {
	var output int
	if num1 > num2 {
		output = num1
	} else {
		output = num2
	}

	return output
}

// Function return multiple values
func swap(a, b int) (int, int) {
	return b, a
}
