package main

import (
	"fmt"
	"math"
)

func main() {
	/* This is a function variable */
	result := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* use the variable as function */
	fmt.Println(result(25))
}
