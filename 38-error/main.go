package main

import "fmt"

func divide(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return value / divider, nil
}

func main() {
	result, err := divide(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
