package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)
	fmt.Println("numbers[1:4]", numbers[1:4])
	fmt.Println("numbers[5:]", numbers[5:])

	numbers1 := numbers[1:4]
	numbers1[1] = 111
	fmt.Println(numbers1)
	fmt.Println(numbers)

	numbers = append(numbers, 9)
	numbers = append(numbers, 10, 11)
	fmt.Println(numbers)

	numbers2 := make([]int, len(numbers), cap(numbers))
	fmt.Println(numbers2)
	copy(numbers2, numbers)
	fmt.Println(numbers2)
	numbers2[2] = 2222

	fmt.Println(numbers)
	fmt.Println(numbers2)
}
