package main

import (
	"fmt"

	"example.com/user/hello/test"
)

func main() {
	fmt.Println("Hello, world.")
	test.Hello()

	m := test.Sarkar{Name: "Mamun Sarkar"}
	m.Display()
	m.Hello()
}
