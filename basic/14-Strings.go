package main

import "fmt"

func main() {
	var greeting = "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s\n", greeting)
	fmt.Println("Length: ", len(greeting))
	fmt.Printf("First char: %c\n", greeting[0])
	fmt.Printf("Second Char: %c\n", greeting[1])

	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x %c, ", greeting[i], greeting[i])
	}
	fmt.Printf("\n")

}
