package test

import "fmt"

// Sarkar is A DataType
type Sarkar struct {
	Name string
}

// Display Function Print the Name
func (m Sarkar) Display() {
	fmt.Println("Hello " + m.Name)
}

// Hello Function Print a Message
func (m Sarkar) Hello() {
	fmt.Println("Welcome to Hello Sarkar")
}
