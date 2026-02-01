package main

import "fmt"

func main() {
	name := "John Doe"

	if name == "John Doe" {
		fmt.Println("Hello, John!")
	} else if name == "Jane Doe" {
		fmt.Println("Hello, Jane!")
	} else {
		fmt.Println("Hello, Stranger!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Your name is long!")
	} else {
		fmt.Println("Your name is short!")
	}
}
