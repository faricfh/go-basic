package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

	fmt.Println("-----")

	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
	}

	fmt.Println("-----")

	names := []string{"John", "Jane", "Doe"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("Name:", names[i])
	// }

	for index, name := range names {
		// fmt.Printf("Index: %d, Name: %s\n", index, name)
		fmt.Println("index", index, "name", name)
	}

	for _, name := range names {
		fmt.Println("Name:", name)
	}

	// for number := range 5 {
	// 	fmt.Println("Number", number)
	// }
}
