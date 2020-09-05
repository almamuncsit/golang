package main

import "fmt"

func findFactorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * findFactorial(i-1)
}
func main() {
	var i int = 10
	fmt.Printf("Factorial of %d is %d \n", i, findFactorial(i))
}
