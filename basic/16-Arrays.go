package main

import "fmt"

func main() {
	var numbers [5]int
	var salaries = []float32{100.0, 20.0, 13.4, 17.0, 55}

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 100
	}

	for j := 0; j < len(numbers); j++ {
		fmt.Printf("Element[%d] = %d\n", j, numbers[j])
	}

	for index, salary := range salaries {
		fmt.Println(index, salary)
	}
}
