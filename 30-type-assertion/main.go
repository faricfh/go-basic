package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()

	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}

	// var resultInt int = result.(int)
	// fmt.Println(resultInt)
}
