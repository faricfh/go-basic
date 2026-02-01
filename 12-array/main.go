package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Bob"

	fmt.Println(names)
	fmt.Println("First Name:", names[0])
	fmt.Println("Second Name:", names[1])
	fmt.Println("Third Name:", names[2])

	// Error: index out of range
	// names[3] = "Alice"

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)

	var numbers = [...]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	fmt.Println("Array Length:", len(numbers))
}
