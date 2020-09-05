package main

import "fmt"

var g_v int = 10

func main() {
	/* local variable declaration */
	var l int
	l = 20
	fmt.Printf("Value of local variable l: %d\n", l)

	fmt.Printf("Value of globar variable g_v: %d\n", g_v)
	chageGV()
	fmt.Printf("Value of globar variable after change g_v: %d\n", g_v)
}

func chageGV() {
	g_v = 100
}
