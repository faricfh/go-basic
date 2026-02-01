package main

import "fmt"

func main() {
	name1 := "John"
	name2 := "John"

	result := name1 == name2
	fmt.Println(result)

	fmt.Println(name1 != name2)
	fmt.Println(name1 > name2)
	fmt.Println(name1 < name2)
	fmt.Println(name1 >= name2)
	fmt.Println(name1 <= name2)

	a := 10
	b := 20

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}
