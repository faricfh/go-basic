package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := &address1 // pass by value

	// address2.City = "Jakarta"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// var address2 *Address = &address1
	address2 := &address1 // pass by reference

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}
