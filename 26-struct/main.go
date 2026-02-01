package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", customer.Name, "my name is", name)
}

func main() {
	var customer Customer
	fmt.Println(customer)

	customer.Name = "John"
	customer.Address = "123 Main St"
	customer.Age = 30

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	// struct literal
	customer2 := Customer{
		Name:    "Jane",
		Address: "456 Elm St",
		Age:     25,
	}

	fmt.Println(customer2)

	// struct literal without field name
	customer3 := Customer{"Bob", "789 Oak St", 35}
	fmt.Println(customer3)

	customer.sayHello("Budi")
	customer2.sayHello("Ani")
	customer3.sayHello("Joko")
}
