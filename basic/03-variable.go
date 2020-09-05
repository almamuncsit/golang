package main

import "fmt"

func main() {
	var x float64
	x = 20.0
	var number int64 = 200
	var a, b, c = 3, 4, "foo"
	d := 42

	fmt.Println(a, b, c)
	fmt.Println(x)
	fmt.Println(d)
	fmt.Println(number)
	fmt.Printf("x is type of %T\n", x)
	fmt.Printf("d is type of %T\n", d)
	fmt.Printf("number is type of %T\n", number)
}
