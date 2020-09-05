package main

import "fmt"

func main() {
	var a int = 17
	var b int = 5
	var c float32 = 234.345
	var avg float32

	fmt.Println(float32(a))
	fmt.Println(int(c))
	fmt.Println("Division: ", a/b, "Reminder: ", a%b)
	avg = float32(a) / float32(b)
	fmt.Printf("Value of mean : %.2f\n", avg)
}
