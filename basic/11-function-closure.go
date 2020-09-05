package main

import "fmt"

func getSequence() func() int {
	x := 0
	output := func() int {
		x += 1
		return x
	}

	return output
}

func main() {
	number := getSequence()

	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
}
