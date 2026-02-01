package main

import "fmt"

func main() {
	type IDNumber string

	var userID IDNumber = "1111111111"
	fmt.Println(userID)

	fmt.Println(IDNumber("2222222222"))
}
