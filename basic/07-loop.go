package main

import "fmt"

func main() {
	var num1 int = 15
	var num2 int
	numbers := []int{1, 2, 3, 4}

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	// While Loop
	for num2 < num1 {
		fmt.Println("Value of num2", num2)
		if num2 == 5 {
			break
		}
		num2++
	}

	// Foreach loop with index
	for index, value := range numbers {
		fmt.Printf("Index is: %d and Number is: %d\n", index, value)
	}

	// Foreach loop
	for item := range numbers {
		fmt.Println("Item is: ", item)
	}
}
