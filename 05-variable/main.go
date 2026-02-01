package main

import "fmt"

func main() {

	var name string
	fmt.Println("Name (empty):", name)

	var name2 = "John Doe"
	fmt.Println("Name2:", name2)

	name3 := "Jane Smith"
	fmt.Println("Name3:", name3)

	name = "Michael Jordan"
	fmt.Println("Name (updated):", name)

	var (
		firstName = "John"
		lastName  = "Doe"
		age       = 30
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}
