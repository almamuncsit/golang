package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Mamun sarkar"

	fmt.Println(len(name))

	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToTitle(name))

	fmt.Println(strings.HasPrefix(name, "Mamun"))
	fmt.Println(strings.HasSuffix(name, "Sarkar"))
	fmt.Println(strings.Contains(name, "Sa"))
	fmt.Println(strings.Count(name, "a"))

	messages := []string{"Welcome", "to", "dhaka"}
	fmt.Println(strings.Join(messages, " "))

	str_to_arr := strings.Split(name, " ")
	fmt.Println(str_to_arr)
	fmt.Println(str_to_arr[0], str_to_arr[1])
	fmt.Printf("%q\n", str_to_arr)

	data := "  hello my     name  is Khan"
	fields := strings.Fields(data)
	fmt.Printf("Fields: %q\n", fields)
	fmt.Printf("Split: %q\n", strings.Split(data, " "))

	fmt.Println(strings.ReplaceAll(name, "Mamun", "Al-Mamun"))
	fmt.Println(strings.Replace("Jone Jone Jone Jone", "e", "ey", 2))
	fmt.Println(strings.Replace("Jone Jone Jone", "Jone", "Jam", -1))

	fmt.Println(strings.Trim("##Welcome to Dhaka$$", "#$"))
	fmt.Println(strings.TrimLeft("##Welcome to Dhaka$$", "#$"))
	fmt.Println(strings.TrimRight("##Welcome to Dhaka$$", "#$"))
	fmt.Println(strings.TrimSpace(" \t\n My Name is Khan \n\t\r\n"))

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	fmt.Println("ba" + strings.Repeat("na", 4))
}
