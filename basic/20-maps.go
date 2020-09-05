package main

import "fmt"

// Vertex Struct
type Vertex struct {
	lat, long float64
}

var m = map[string]Vertex{
	"one": Vertex{48.68433, -234.39967},
	"two": Vertex{98.42202, -567.08408},
}

func main() {
	// var customer map[string]string
	customer := make(map[string]string)

	customer["name"] = "Mamun"
	customer["email"] = "mamun@gmail.com"
	customer["phone"] = "0124523495"

	for item := range customer {
		fmt.Printf("%s is %s\n", item, customer[item])
	}
	fmt.Println("")
	for index, value := range customer {
		fmt.Printf("%s is %s\n", index, value)
	}

	address, ok := customer["address"]
	if ok {
		fmt.Println(address)
	} else {
		fmt.Println("Address not found")
	}

	delete(customer, "phone")
	fmt.Println("\nPrint after delete")
	for index, value := range customer {
		fmt.Printf("%s is %s\n", index, value)
	}

	fmt.Println(m)
	fmt.Println(m["one"])
	fmt.Println(m["one"].lat, m["two"].long)
}
