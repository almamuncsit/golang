package main

import "fmt"

func main() {

	var number int = 10

	// If condition
	if number < 20 {
		fmt.Printf("number is less than 20\n")
	}

	// If else if conditiona
	if number == 5 {
		fmt.Printf("number is 5\n")
	} else if number == 10 {
		fmt.Printf("number is 10\n")
	} else {
		fmt.Printf("number greter than 10\n")
	}

	// switch statementy
	gpa := "B"
	switch {
	case gpa == "A":
		fmt.Printf("Excellent Result!\n")
	case gpa == "B", gpa == "C":
		fmt.Printf("Well done\n")
		fmt.Printf("Good Boy\n")
	case gpa == "D":
		fmt.Printf("You passed\n")
	case gpa == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid gpa\n")
	}
}
