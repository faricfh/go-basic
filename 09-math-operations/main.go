package main

import "fmt"

func main() {
	var a = 10
	var b = 15
	var c = a + b

	fmt.Println("Addition:", c)

	var i = 10
	i += 5

	fmt.Println("Addition Assignment:", i)

	var j = 1
	j++
	j++

	fmt.Println("Increment:", j)

	j--
	fmt.Println("Decrement:", j)
}
