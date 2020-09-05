package main

import "fmt"

// Users Type
type Users struct {
	name  string
	email string
	phone int64
}

func (user Users) printUser() {
	fmt.Println(user.name)
	fmt.Println(user.email)
	fmt.Println(user.phone)
}

func main() {
	var user1 Users
	var user2 Users

	user1.name = "Mamun"
	user1.name = "mamun@gmail.com"
	user1.phone = 2435235

	user2.name = "Sarkar"
	user2.email = "test@gmail.com"

	fmt.Println(user1.name)
	fmt.Println(user2.name)

	user1.printUser()
	user2.printUser()
}
