package main

import "fmt"

// func Ups() interface {}
func Ups() any {
	// return "Ups"
	return 1
}

func main() {
	data := Ups()
	fmt.Println(data)
}
