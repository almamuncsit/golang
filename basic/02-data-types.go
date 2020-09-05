package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int32
	var y float64
	var z string

	x = 10
	y = 23.8
	z = "Hello"

	fmt.Println(x, y, z)
	fmt.Println("Data Type: ", reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z))
}
