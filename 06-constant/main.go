package main

import "fmt"

func main() {
	// const firstName string = "John"
	// const lastName = "Doe"
	// const age int = 30

	const (
		firstName string = "John"
		lastName         = "Doe"
		age              = 30
	)

	// error
	// firstName = "Jane"

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	// fmt.Println("Age:", age)
}
