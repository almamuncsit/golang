package main

import "fmt"

func main() {
	var number int = 20
	var pointer *int
	pointer = &number

	fmt.Printf("Address of number: %x\n", &number)

	fmt.Printf("Address of pointer: %x\n", pointer)

	fmt.Printf("Value of *pointer: %d\n", *pointer)
}
