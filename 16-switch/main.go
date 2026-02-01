package main

import "fmt"

func main() {
	name := "John Doe"

	switch name {
	case "John Doe":
		fmt.Println("Hello, John!")
	case "Jane Doe":
		fmt.Println("Hello, Jane!")
	default:
		fmt.Println("Hello, Stranger!")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Your name is long!")
	case false:
		fmt.Println("Your name is short!")
	}

	name = "Alexander Hamilton"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Your name is very long!")
	case length > 5:
		fmt.Println("Your name is long!")
	default:
		fmt.Println("Your name is short!")
	}
}
