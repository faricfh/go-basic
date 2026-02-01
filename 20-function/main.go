package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "John", "Doe"
}

// func getCompleteName() (firstName string, middleName string, lastName string) {
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Doe"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

// alias function
type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Budi" {
		return "..."
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blasklist Blacklist) {
	if blasklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("You are registered", name)
	}
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	sayHello()
	sayHelloTo("John", "Doe")

	result := getHello("John Doe")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// firstName, _ := getFullName()
	// fmt.Println(firstName)

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(sumAll(slice...))

	totalSlice := sumAllSlice(slice)
	fmt.Println(totalSlice)

	// function value (function as variable)
	goodbye := getGoodBye
	fmt.Println(goodbye("John"))

	// function parameter (function as parameter)
	sayHelloWithFilter("John", spamFilter)
	sayHelloWithFilter("Budi", spamFilter)

	// anonymous function
	blacklist := func(name string) bool {
		return name == "Budi"
	}

	registerUser("John", blacklist)

	registerUser("Budi", func(name string) bool {
		return name == "Budi"
	})

	// recursive function
	// fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
