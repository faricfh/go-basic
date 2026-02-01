package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "John Doe"
	// person["address"] = "123 Main St"
	// person["city"] = "Anytown"

	person := map[string]string{
		"name":    "John Doe",
		"address": "123 Main St",
		"city":    "Anytown",
	}

	fmt.Println("Person Map:", person)

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])
	fmt.Println("City:", person["city"])
	fmt.Println("Country:", person["country"])

	book := make(map[string]string)
	book["title"] = "The Go Programming Language"
	book["author"] = "Alan A. A. Donovan"
	book["publisher"] = "Addison-Wesley"
	book["year"] = "2015"

	fmt.Println("Book Map:", book)

	delete(book, "year")

	fmt.Println("Book Map after deleting 'year':", book)
}
